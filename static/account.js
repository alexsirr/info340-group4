$(document).ready(function(){
    $.get("/QuserInfo", function(data){
        $("#user-info").append(data);
    }, "html")

    $.get("/QuserAddr", function(data){
        $("#user-addr").append(data);
    }, "html")

    $.get("/QuserBooking", function(data){
        $("#user-booking").append(data);
    }, "html")
})