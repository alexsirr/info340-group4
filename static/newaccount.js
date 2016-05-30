$(document).ready(function(){

    document.getElementById("newaccount").addEventListener("click", function() {
        console.log("blah");
        newaccount();
    });

    function newaccount(){
        $.post("/Qnewaccount", {fname: $("#fname").val(), lname: $("#lname").val(), email: $("#email").val(), phone: $("#phone").val(), password: $("#password").val()})
    }
})