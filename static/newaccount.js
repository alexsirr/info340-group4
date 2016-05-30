$(document).ready(function(){

    $("#create-account").click(function() {
		console.log("clicked");
		newaccount();
	})

    function newaccount(){
		var vals = {fname: $("#fname").val(), lname: $("#lname").val(), email: $("#email").val(), phone: $("#phone").val(), password: $("#password").val()};
		console.log(vals);
        $.post("/Qnewaccount", vals).done(function(data) {
			$("#result").append(data);
		}, "html");
    }
})
