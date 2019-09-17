$(document).ready(function() {

    $("#uninstallams").on('click', function() {
        var server = document.getElementById("server").value
        var version = document.getElementById("version").value
        var revision = document.getElementById("revision").value
        var username = document.getElementById("username").value
        var password = document.getElementById("password").value
        if (server !=""&&version !=""&&revision !=""&&username !=""&&password!=""){
            $.ajax({
                url: "http://localhost:8080/uninstallams",
                type: "POST",
        
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'text/plain'
                },
        
                dataType: "json",
                data: JSON.stringify({
                    "server":server,
                    "version":version,
                    "revision":revision,
                    "username":username,
                    "password":password
                }),
                success: function( data){
                    console.log(data);
                   document.getElementById("response").innerHTML=data.statusams;
                },
                error: function( jqXhr, textStatus, errorThrown ){
                    console.log( errorThrown );
                }
            });
        }else{
           alert("lack of server or version or revision or username or password")
        }
    });
});