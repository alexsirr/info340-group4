$(document).ready(function(){

    $("#newaccount").click(function() {
		newaccount();
	})

    function newaccount(){
        $.post("/Qnewaccount", {fname: $("#fname").val(), lname: $("#lname").val(), email: $("#email").val(), phone: $("#phone").val(), password: $("#password").val()})
    }
})
