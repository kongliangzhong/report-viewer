package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

var routes = []Route{
    Route{
        "Index",
        "Get",
        "/",
        Index,
    },
    Route{
        "Test Page",
        "Get",
        "/test001",
        TestPage,
    },
    Route{
        "New Chart",
        "Get",
        "/new-chart",
        NewChartPage,
    },
    Route{
        "add report item",
        "Post",
        "/add-chart",
        AddChart,
    },
    Route{
        "Mysql DB Connections",
        "Get",
        "/mysql-conns",
        GetAllMysqlConns,
    },
    Route{
        "Report Item Page",
        "Get",
        "/chart-item/{id}",
        ReportChartPage,
    },
    Route{
        "Chart List Page",
        "Get",
        "/chart-list",
        ChartListPage,
    },
    Route{
        "Chart Data Api",
        "Get",
        "/chart-data/{id}",
        ChartDataRouter,
    },
}
