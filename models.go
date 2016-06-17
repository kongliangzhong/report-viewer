package main

type ReportItem struct {
    Id         string
    Title      string
    Desc       string
    CreateTime string
    TableName  string
    HandlerKey string
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
