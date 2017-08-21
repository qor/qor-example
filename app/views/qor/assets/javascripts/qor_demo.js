var OrderChart,UsersChart;
function RenderChart(ordersData, usersData) {
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

// qor dashboard
$(document).ready(function() {
  var yesterday = (new Date()).AddDate(-1);
  var defStartDate = yesterday.AddDate(-6);
  $("#startDate").val(defStartDate.Format("yyyy-MM-dd"));
  $("#endDate").val(yesterday.Format("yyyy-MM-dd"));
  $(".j-update-record").click(function(){
    $.getJSON("/admin/reports.json",{startDate:$("#startDate").val(), endDate:$("#endDate").val()},function(jsonData){
      RenderChart(jsonData.Orders,jsonData.Users);
    });
  });
  $(".j-update-record").click();

  $(".yesterday-reports").click(function() {
    $("#startDate").val(yesterday.Format("yyyy-MM-dd"));
    $("#endDate").val(yesterday.Format("yyyy-MM-dd"));
    $(".j-update-record").click();
    $(this).blur();
  });

  $(".this-week-reports").click(function() {
    var beginningOfThisWeek = yesterday.AddDate(-yesterday.getDay() + 1)
    $("#startDate").val(beginningOfThisWeek.Format("yyyy-MM-dd"));
    $("#endDate").val(beginningOfThisWeek.AddDate(6).Format("yyyy-MM-dd"));
    $(".j-update-record").click();
    $(this).blur();
  });

  $(".last-week-reports").click(function() {
    var endOfLastWeek = yesterday.AddDate(-yesterday.getDay())
    $("#startDate").val(endOfLastWeek.AddDate(-6).Format("yyyy-MM-dd"));
    $("#endDate").val(endOfLastWeek.Format("yyyy-MM-dd"));
    $(".j-update-record").click();
    $(this).blur();
  });
});
