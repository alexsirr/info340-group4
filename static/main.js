$(document).ready(function(){
    $.get("/QavailableRooms", function(data){
        $("#available-rooms").append(data);
    }, "html")
})