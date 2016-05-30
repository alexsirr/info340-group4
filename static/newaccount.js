$(document).ready(function(){

    $("#newaccount").click(function(){
        console.log("blah");
        newaccount();
    });

    function newaccount(){
        $.post("/Qnewaccount", {fname: $("#fname").val(), lname: $("#lname").val(), email: $("#email").val(), phone: $("#phone").val(), password: $("#password").val()})
        .done(function(data){
          if(data.result == "failed"){
            console.log(data)
            $("#result").text("Failed to create account! " + data.message);
          } else {
            console.log(data)
            $("#result").text("Created New Account!");
          }
        });
    }
})