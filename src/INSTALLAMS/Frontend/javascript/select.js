$(document).ready(function() {
    
    $("#selectams").on('click', function() {
        var server = document.getElementById("server").value
        var version = document.getElementById("version").value
        var revision = document.getElementById("revision").value
        var username = document.getElementById("username").value
        var password = document.getElementById("password").value
        $.ajax({
            url: "http://localhost:8080/selectams",
            type: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'text/plain'
            },
            dataType: "json",
            success: function(data) {
                document.getElementById("version").value=data.version
                document.getElementById("revision").value=data.revision    
            },
        });
    });
});