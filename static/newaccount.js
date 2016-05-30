$(document).ready(function(){

    $("#create-account").click(function() {
		console.log("clicked");
		newaccount();
	})

    function newaccount(){
		var vals = {fname: $("#fname").val(), lname: $("#lname").val(), email: $("#email").val(), phone: $("#phone").val(), password: $("#password").val()};
		console.log(vals);
    	for (var key in vals) {
			if (vals[key].length == 0) {
				$("#result").text("Please Complete All Items");
				return;
        	}
      	}
      
      	if (!(/(\d{3}-){2}\d{4}/.test(vals["phone"]))) {
      		$("#result").text("Phone Number Must Be in Format ###-###-#####");
        	return;
      	}
      
		$.post("/Qnewaccount", vals).done(function(data) {
			$("#result").text(data);
		}, "html");
    }
})
