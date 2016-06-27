var myChart = echarts.init(document.getElementById('main'));

var option = {
    "baseOption": {
        "timeline": {
            "axisType": "category",
            "data": [ "1" ]
        },
        "title": {
            "text": "",
            "subtext": "",
            "x": "center"
        },
        "tooltip": {
            "trigger": "item",
            "formatter": "{b} : {c} ({d}%)"
        },
        "series": [
            {
                "name": "",
                "type": "pie",
                "radius": [
                    "0",
                    "75%"
                ],
                "center": [
                    "50%",
                    "50%"
                ],
                "data": null,
                "itemStyle": {
                    "emphasis": {
                        "shadowBlur": 10,
                        "shadowOffsetX": 0,
                        "shadowColor": "rgba(0, 0, 0, 0.5)"
                    }
                }
            }
        ]
    },
    "options": [
        {
            "title": "",
            "series": [
                {
                    "type": "",
                    "name": "",
                    // "center": null,
                    // "radius": null,
                    "data": [
                        {
                            "name": "a",
                            "value": 0.41
                        },
                        {
                            "name": "b",
                            "value": 0.23
                        },
                        {
                            "name": "c",
                            "value": 0.36
                        }
                    ]
                }
            ]
        }
    ]
}

myChart.setOption(option);
