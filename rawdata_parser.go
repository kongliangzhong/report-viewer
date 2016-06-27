package main

import (
    "log"
    "errors"
    "strings"
    "strconv"
)

type DataParseResult struct {
    xdata []string
    ydata [][]float64
}

type DataParseResultArray []DataParseResult

const XAxisType_C = "category"
const XAxisType_V = "value"

func parseRawDatas(rawDatas []string, dataFormat string) (parseResultArr DataParseResultArray, err error) {
    for _, data := range rawDatas {
        xdata, ydata, err := parseSingleRawData(data, dataFormat)
        if err != nil {
            log.Println("Error:", err.Error())
            continue
        }
        res := DataParseResult{xdata, ydata}
        if len(xdata)==0 || len(ydata)==0 {
            log.Println("Error: parse result is empty")
            continue
        }
        parseResultArr = append(parseResultArr, res)
    }

    if len(parseResultArr) == 0 {
        err = errors.New("数据为空！")
    }

    return
}

func parseSingleRawData(data string, format string) (xdata []string, ydata [][]float64, err error) {
    formatFlds := strings.Split(format, ":")
    fLen := len(formatFlds)
    if fLen < 2 || formatFlds[0] != "x" || formatFlds[1] != "y" {
        err = errors.New("invalid dataformat:" + format)
        return
    }

    xdata = make([]string, 0, 10)
    ydata = make([][]float64, 0, 2)
    for i := 0; i < fLen-1; i++ {
        ydata = append(ydata, []float64{})
    }

    items := strings.Split(data, " ")
    for _, item := range items {
        flds := strings.Split(item, ":")
        if len(flds) < fLen {
            log.Println("data not match format. data:", item, "; format:", format)
            continue
        }
        for i := 0; i < fLen; i++ {
            if i == 0 {
                xdata = append(xdata, flds[0])
            } else {
                f64, ferr := strconv.ParseFloat(flds[i], 64)
                if ferr != nil {
                    log.Println("convert string to float64 failed:", ferr.Error())
                }
                ydata[i-1] = append(ydata[i-1], f64)
            }
        }
    }
    return
}

func (pr *DataParseResultArray) detectXAxisType() (string, error) {
    var detectFunc = func(xdata []string) string {
        for _, xd := range xdata {
            _, err := strconv.ParseFloat(xd, 64)
            if err != nil {
                return XAxisType_C
            }
        }
        return XAxisType_V
    }

    var checkXAxisData = func() error {
        var firstXData []string
        for i, res := range *pr {
            if i == 0 {
                firstXData = res.xdata
            } else {
                xd := res.xdata
                if len(xd) != len(firstXData) {
                    return errors.New("类型为category时XAxis数据不一致")
                }

                for j, s := range xd {
                    if s != firstXData[j] {
                        return errors.New("类型为category时XAxis数据不一致")
                    }
                }
            }
        }
        return nil
    }

    var xaxisType string
    for _, res := range *pr {
        xaxisType = detectFunc(res.xdata)
        if xaxisType == XAxisType_C {
            break
        }
    }

    if xaxisType == XAxisType_C {
        return xaxisType, checkXAxisData()
    } else {
        return xaxisType, nil
    }
}
