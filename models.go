package main

type ReportItem struct {
    Id            string
    Title         string
    Legend        string
    Desc          string
    CreateTime    string
    TableName     string
    DataField     string
    DataFormat    string
    ChartType     string
    HandlerKey    string
    ShowTimeLine  bool // timeline support
    TimeLineField string
    //XAxisType     string
}

// type ChartDataRaw struct {
//     ChartType string
//     DataType  string
//     Content   string
// }

type PieDataItem struct {
    Name  string  `json:"name"`
    Value float64 `json:"value"`
}

type PieData []PieDataItem

func (pd PieData) Len() int           { return len(pd) }
func (pd PieData) Swap(i, j int)      { pd[j], pd[i] = pd[i], pd[j] }
func (pd PieData) Less(i, j int) bool { return pd[i].Value > pd[j].Value }

type ChartData struct {
    Title        string
    HandlerKey   string
    ShowTimeLine bool
    OptionsData  []interface{}
}
