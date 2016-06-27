package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

type ChartListPageModel struct {
    ReportItems []ReportItem
}

func concatPageList(mainPage string) []string {
    return []string{
        mainPage,
        "assets/pages/navbarTop.html",
        "assets/pages/navbarSide.html",
        "assets/pages/header.html",
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    pageList := concatPageList("assets/pages/indexPage.html")
    t, _ := template.ParseFiles(pageList...)
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func NewChartPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(concatPageList("assets/pages/newChart.html")...)
    err := t.Execute(w, struct {
        Flag    int
        Message string
    }{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func AddChart(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("reportTitle")
    tableName := r.FormValue("tableName")
    dataField := r.FormValue("dataField")
    dataFormat := r.FormValue("dataFormat")
    chartType := r.FormValue("chartType")
    legend := r.FormValue("legend")
    timeLineFiled := r.FormValue("timeLineField")

    ri := ReportItem{Title: title, Legend: legend, Desc: "", TableName: tableName,
        DataField: dataField, DataFormat: dataFormat, ChartType: chartType, HandlerKey: "",
        ShowTimeLine: true, TimeLineField: timeLineFiled}

    err := dataService.SaveReportItem(ri)

    log.Println("title:", title)
    t, _ := template.ParseFiles(concatPageList("assets/pages/newChart.html")...)

    paramLabels := []string{"数据表名", "数据字段名", "数据格式", "图表类型", "图例说明", "时间轴字段名"}
    nonNullableParams := []string{tableName, dataField, dataFormat, chartType, legend, timeLineFiled}
    for i, p := range nonNullableParams {
        log.Println("p:", p)
        if p == "" {
            t.Execute(w, struct {
                Flag    int
                Message string
            }{0, "创建失败，字段" + paramLabels[i] + "不能为空"})
            return
        }
    }

    if err != nil {
        err2 := t.Execute(w, struct {
            Flag    int
            Message string
        }{0, "创建报表失败! 错误信息：" + err.Error()})
        if err2 != nil {
            http.Error(w, err2.Error(), http.StatusInternalServerError)
        }
    } else {
        err = t.Execute(w, struct {
            Flag    int
            Message string
        }{1, "创建报表成功！"})
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func ChartListPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(concatPageList("assets/pages/chartList.html")...)
    reportItems := dataService.GetReportItems()
    model := ChartListPageModel{reportItems}
    err := t.Execute(w, model)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func TestPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(concatPageList("assets/pages/testPage.html")...)
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetAllMysqlConns(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("assets/pages/dbConns.html")
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func ReportChartPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(concatPageList("assets/pages/chart.html")...)
    vars := mux.Vars(r)
    id := vars["id"]
    model := struct{ Id string }{id}
    err := t.Execute(w, model)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func ChartDataRouter(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    idi, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "invalid report id:"+id, http.StatusInternalServerError)
        return
    }

    ri := dataService.GetReportItem(idi)
    //log.Println("ri:", ri)
    //data := getData(ri)
    data := getChartData(ri)
    //log.Println("chart data:", data)

    if data == nil {
        http.Error(w, "data is nil", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        panic(err)
    }
}

func getChartData(ri ReportItem) interface{} {
    datas, timelineData, err := dataService.GetDataByTimeline(ri.TableName, ri.DataField, ri.TimeLineField)
    if err != nil {
        log.Println("Error:", err.Error())
        return nil
    }
    //log.Println("datas", datas)

    switch ri.ChartType {
    case "bar":
        return genBarOption(datas, timelineData, ri.DataFormat)
    case "pie":
        return genPieOption(datas, timelineData, ri.DataFormat)
    default:
        log.Println("Unknown chart-type:", ri.ChartType)
        return nil
    }
}

func genBarOption(rawDatas, timelineData []string, dataFormat string) interface{} {
    var genSeriesData = func(xdata []string, ydata []float64, xaxisType string) (interface{}, error) {
        if xaxisType == XAxisType_C {
            return ydata, nil
        } else {
            xdataF64, err := StrArrToF64Arr(xdata)
            if err != nil {
                return nil, err
            }
            return ZipF64(xdataF64, ydata)
        }
    }

    baseBarOption := NewBaseBarOption()
    baseBarOption.Timeline= TimelineType{XAxisType_C, timelineData}

    parseResultArr, err := parseRawDatas(rawDatas, dataFormat)
    if err != nil {
        log.Println("Error:", err.Error())
        return nil
    }

    xaxisType, err := parseResultArr.detectXAxisType()
    if err != nil {
        log.Println("XAxis Type can not detected. ", err.Error())
        return nil
    }

    baseBarOption.XAxis = append(baseBarOption.XAxis, SimpleData{Type: xaxisType})

    if xaxisType == XAxisType_C {
        baseBarOption.XAxis[0].Data = parseResultArr[0].xdata
    }

    timeLineOptions := []TimelineOption{}
    for _, res := range parseResultArr {
        st := SeriesType{Type: "bar"}
        baseBarOption.Series = append(baseBarOption.Series, st)

        tlo := TimelineOption{}
        for _, yd := range res.ydata {
            seriesData, err := genSeriesData(res.xdata, yd, xaxisType)
            if err != nil {
                log.Println("Error:", err.Error())
                continue
            }
            ser := SeriesType{Data: seriesData}
            tlo.Series = append(tlo.Series, ser)
        }
        timeLineOptions = append(timeLineOptions, tlo)
    }

    resultOption := FullOption{baseBarOption, timeLineOptions}
    return resultOption
}

func genPieOption(rawDatas, timelineData []string, dataFormat string) interface{} {
    basePieOption := NewBasePieOption()
    basePieOption.Timeline= TimelineType{XAxisType_C, timelineData}

    parseResultArr, err := parseRawDatas(rawDatas, dataFormat)
    if err != nil {
        log.Println("Error:", err.Error())
        return nil
    }

    timelineOptions := []TimelineOption{}
    for _, res := range parseResultArr {
        tlo := TimelineOption{}
        data := []PieDataItem{}
        log.Println("res:", res)
        for i, name := range res.xdata {
            item := PieDataItem{name, res.ydata[0][i]}
            data = append(data, item)
        }
        tlo.Series = append(tlo.Series, SeriesType{Data: data})
        timelineOptions = append(timelineOptions, tlo)
    }

    resultOption := FullOption{basePieOption, timelineOptions}
    return resultOption
}
