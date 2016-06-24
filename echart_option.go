package main

// import "encoding/json"
// import "os"

// TODO: find out how to  change defaut json encoding behavior.
// use lower-case First letter name as its default bahavior.
type BaseBarOption struct {
    Timeline TimelineType `json:"timeline"`
    Tooltip  TooltipType  `json:"tooltip"`
    Toolbox  struct {
        Feature struct {
            DataView struct {
                Show     bool `json:"show"`
                ReadOnly bool `json:"readOnly"`
            } `json:"dataView"`
            MagicType struct {
                Show bool     `json:"show"`
                Type []string `json:"type"`
            } `json:"magicType"`
            Restore struct {
                Show bool `json: "show"`
            } `json:"restore"`
            SaveAsImage struct {
                Show bool `json: "show"`
            } `json: "saveAsImage"`
        } `json:"feature"`
    } `json:"toolbox"`

    Legend struct {
        Data []string `json:"data"`
        X    string   `json:"x"`
    } `json:"legend"`

    XAxis []SimpleData `json:"xAxis"`

    YAxis []SimpleData `json:"yAxis"`

    Series []SeriesType `json:"series"`
}

type TimelineType struct {
    AxisType string   `json:"axisType"`
    Data     []string `json:"data"`
}

type SimpleData struct {
    Type string      `json:"type"`
    Data interface{} `json:"data"`
}

type SeriesType struct {
    //SimpleData
    Type   string      `json:"type"`
    Name   string      `json:"name"`
    Center []string    `json:"center"`
    Radius string      `json:"radius"`
    Data   interface{} `json:"data"`
}

type TooltipType struct {
    Trigger string `json:"trigger"`
}

type BasePieOption struct {
    Timeline TimelineType `json:"timeline"`
    Title    struct {
        Text    string `json:"text"`
        Subtext string `json:"subtext"`
        X       string `json:"x"`
    } `json: "title"`

    Tooltip struct {
        Trigger   string `json:"trigger"`
        Formatter string `json:"formatter"`
    } `json:"tooltip"`

    Series []BasePieSeriesType `json: "series"`
}

type TimelineOption struct {
    Title  string       `json:"title"`
    Series []SeriesType `json:"series"`
}

type FullOption struct {
    BaseOption interface{} `json:"baseOption"`
    Options    interface{} `json:"options"`
}

type BasePieSeriesType struct {
    Name   string        `json:"name"`
    Type   string        `json:"type"`
    Radius string        `json:"radius"`
    Center []string      `json:"center"`
    Data   []interface{} `json:"data"`

    ItemStyle ItemStyleType `json:"itemStyle"`
}

type ItemStyleType struct {
    Emphasis EmphasisType `json:"emphasis"`
}

type EmphasisType struct {
    ShadowBlur    int    `json:"shadowBlur"`
    ShadowOffsetX int    `json:"shadowOffsetX"`
    ShadowColor   string `json:"shadowColor"`
}

func NewBaseBarOption() BaseBarOption {
    bbo := BaseBarOption{
        //XAxis:  []SimpleData{SimpleData{Type: "value"}},
        YAxis: []SimpleData{SimpleData{Type: "value"}},
        //Series: []SeriesType{SimpleData{Type: "bar"}},
    }

    bbo.Tooltip.Trigger = "axis"
    bbo.Toolbox.Feature.DataView.Show = true
    bbo.Toolbox.Feature.DataView.ReadOnly = true
    bbo.Toolbox.Feature.MagicType.Show = true
    bbo.Toolbox.Feature.MagicType.Type = []string{"line", "bar"}
    bbo.Toolbox.Feature.Restore.Show = true
    bbo.Toolbox.Feature.SaveAsImage.Show = true

    // enc := json.NewEncoder(os.Stdout)
    // enc.Encode(bbo)

    return bbo
}

func NewBasePieOption() BasePieOption {
    basePieOption := BasePieOption{}
    basePieOption.Title.X = "center"
    basePieOption.Tooltip.Trigger = "item"
    basePieOption.Tooltip.Formatter = "{b} : {c} ({d}%)"

    series0 := BasePieSeriesType{Type: "pie", Radius: "55%", Center: []string{"50%", "60%"},
        ItemStyle: ItemStyleType{
            Emphasis: EmphasisType{
                ShadowBlur: 10, ShadowOffsetX: 0, ShadowColor: "rgba(0, 0, 0, 0.5)",
            },
        },
    }
    basePieOption.Series = []BasePieSeriesType{series0}
    return basePieOption
}
