var myChart = echarts.init(document.getElementById('main'));

var barOption = {
    tooltip: {
        trigger: 'axis'
    },
    toolbox: {
        feature: {
            dataView: {show: true, readOnly: false},
            magicType: {show: true, type: ['line', 'bar']},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    legend: {
        // data:['蒸发量','降水量','平均温度'],
        x: 'left',
    },
    xAxis: [
        {
            type: 'value',
        }
    ],
    yAxis: [
        {
            type: 'value',
        },
    ],
    series: [
        {
            type:'bar',
        },
    ]
};

var basePieOption =
    option = {
        title : {
            text: '',
            subtext: '',
            x:'center'
        },
        tooltip : {
            trigger: 'item',
            formatter: "{b} : {c} ({d}%)" //"{a} <br/>{b} : {c} ({d}%)"
        },
        // legend: {
        //     orient: 'vertical',
        //     left: 'left',
        //     data: ['直接访问','邮件营销','联盟广告','视频广告','搜索引擎']
        // },
        series : [
            {
                name: '',
                type: 'pie',
                radius : '55%',
                center: ['50%', '60%'],
                data:[],
                itemStyle: {
                    emphasis: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
        ]
    };

var id = $("#chartdata-id").val();

$.get("/chart-data/" + id, function(data, status){
    console.debug("data:", data);
    //var handlerKey = data.handlerKey;
    //console.debug("handlerKey:", handlerKey);
    //var setter = allOptionSetters[handlerKey];
    //var option = setter(data);
    //console.debug("option:", option);
    console.debug("data json:", JSON.stringify(data, null, 2));
    myChart.setOption(data);
});

// var setFeatureBarData =  function(data) {
//     var option = barOption;
//     option.legend.data = data.legend;
//     option.xAxis[0].type = data.xAxisType;
//     var series0 = {};
//     series0.name = "xx";
//     series0.type = "bar";
//     series0.data = data.data;
//     option.series[0] = series0;

//     return option;
// }

// var setFeaturePieData = function(data) {
//     var option = basePieOption;
//     option.title.text = data.title;
//     option.series[0].data = data.data;
//     return option;
// }

// var allOptionSetters = {
//     "feat2-bar" : setFeatureBarData,
//     "feat10-pie" : setFeaturePieData
// }
