package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "strings"
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
    t, _ := template.ParseFiles(concatPageList("assets/pages/newChart.html")...);
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
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

func Test(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("assets/pages/testPage.html")
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
    data := getData(ri)
    //log.Println("data:", data)

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

func getData(ri ReportItem) interface{} {
    switch ri.HandlerKey {
    case "feat2-bar":
        return feat2Data(ri)
    case "feat10-pie":
        return feat10Data(ri)
    default:
        log.Println("can not found handler for key:", ri.HandlerKey)
        return nil
    }
}

func feat10Data(ri ReportItem) interface{} {
    pieData := struct {
        HandlerKey string        `json:"handlerKey"`
        Title      string        `json:"title"`
        Data       []PieDataItem `json:"data"`
    }{}
    pieData.HandlerKey = ri.HandlerKey
    pieData.Title = "feat10 piechart"
    pieData.Data = []PieDataItem{
        PieDataItem{32, "a"},
        PieDataItem{43, "b"},
        PieDataItem{87, "c"},
    }
    return pieData
}

func feat2Data(ri ReportItem) interface{} {
    barData := struct {
        HandlerKey string   `json:"handlerKey"`
        Legend     []string `json:"legend"`
        XAxisType  string   `json:"xAxisType"`
        Data       [][]int  `json:"data"`
    }{}

    barData.HandlerKey = ri.HandlerKey
    barData.Legend = []string{"feat2"}
    barData.XAxisType = "value"

    var content string
    qSql := "select distribution from feat2 limit 1"
    err := dataService.QueryRow(qSql, &content)
    if err != nil {
        log.Println("ERRORxxx:", err)
        return nil
    }
    //log.Println("content:", content)

    data := [][]int{}
    entries := strings.Split(content, " ")
    for _, e := range entries {
        flds := strings.Split(e, ":")
        if len(flds) < 2 {
            log.Println("ERROR: can not parse kv:", e)
        } else {
            x, err := strconv.Atoi(flds[0])
            if err != nil {
                log.Println("Atoi failed:", err)
                continue
            }
            y, err := strconv.Atoi(flds[1])
            if err != nil {
                log.Println("Atoi failed:", err)
                continue
            }
            data = append(data, []int{x, y})
        }
    }
    barData.Data = data

    return barData
}
