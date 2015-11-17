var OrderChart,UsersChart,ChannelChart;
function RenderChart(ordersData, usersData, channelsData) {
    Chart.defaults.global.responsive = true;

    var orderDateLables = [];
    var orderCounts = [];
    for (var i = 0; i < ordersData.length; i++) {
        orderDateLables.push(ordersData[i].Date.substring(5,10));
        orderCounts.push(ordersData[i].Total)
    }
    if(OrderChart){
        OrderChart.destroy();
    }
    var orders_context = document.getElementById("orders_report").getContext("2d");
    var orders_data = ChartData(orderDateLables,orderCounts);
    OrderChart = new Chart(orders_context).Line(orders_data, "");

    var usersDateLables = [];
    var usersCounts = [];
    for (var i = 0; i < usersData.length; i++) {
        usersDateLables.push(usersData[i].Date.substring(5,10));
        usersCounts.push(usersData[i].Total)
    }
    if(UsersChart){
        UsersChart.destroy();
    }
    var users_context = document.getElementById("users_report").getContext("2d");
    var users_data = ChartData(usersDateLables,usersCounts);
    UsersChart = new Chart(users_context).Bar(users_data, "");

    if(ChannelChart){
        ChannelChart.destroy();
    }
    var channelData = [];
    for (var i = 0; i < channelsData.length; i++) {
        channelData.push(channelsData[i].Total)
    }
    var channels_context = document.getElementById("channels_report").getContext("2d");
    var channels_data=ChannelData(channelData);
    ChannelChart=new Chart(channels_context).Pie(channels_data, {segmentShowStroke : false,segmentStrokeWidth : 4,animationEasing:"linear"});
}

function ChartData(lables, counts) {
    var chartData = {
      labels: lables,
      datasets: [
      {
        label: "Users Report",
        fillColor: "rgba(151,187,205,0.2)",
        strokeColor: "rgba(151,187,205,1)",
        pointColor: "rgba(151,187,205,1)",
        pointStrokeColor: "#fff",
        pointHighlightFill: "#fff",
        pointHighlightStroke: "rgba(151,187,205,1)",
        data: counts
      }
      ]
    };
    return chartData;
}
function ChannelData(data){
    var channels_data = [
        {
            value: data[1],
            color: "#00CC00",
            highlight: "#00CC00",
            label: "Computer"
        },
        {
            value: data[0],
            color:"#0000CC",
            highlight: "#0000CC",
            label: "Mobile"
        },
        {
            value: data[2],
            color: "#CC0000",
            highlight: "#CC0000",
            label: "Others"
        }
    ]
    return channels_data;
}

Date.prototype.Format = function (fmt) {
    var o = {
        "M+": this.getMonth() + 1,
        "d+": this.getDate(),
        "h+": this.getHours(),
        "m+": this.getMinutes(),
        "s+": this.getSeconds(),
        "q+": Math.floor((this.getMonth() + 3) / 3),
        "S": this.getMilliseconds()
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

Date.prototype.AddDate = function (add){
    var date = this.valueOf();
    date = date + add * 24 * 60 * 60 * 1000
    date = new Date(date)
    return date;
}