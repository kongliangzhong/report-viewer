package main

type ReportItem struct {
    Id            string
    Title         string
    Desc          string
    CreateTime    string
    TableName     string
    HandlerKey    string
    ShowTimeLine  bool   // timeline support
    TimeLineField string
}

// type ChartDataRaw struct {
//     ChartType string
//     DataType  string
//     Content   string
// }

type PieDataItem struct {
    Value int    `json:"value"`
    Name  string `json:"name"`
}

type ChartData struct {
    Title       string
    HandlerKey  string
    OptionsData []interface{}
}
