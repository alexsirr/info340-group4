$(document).ready(function(){
    $.get("/QuserInfo", function(data){
        $("#user-info").append(data);
    }, "html")
})