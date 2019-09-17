$(document).ready(function() {
    var server = document.getElementById("server").value
    var version = document.getElementById("version").value
    var revision = document.getElementById("revision").value
    var username = document.getElementById("username").value
    var password = document.getElementById("password").value
    $("#browse").on('click', function() {
        $("#file-input").trigger('click');
       
    });
    $("#file-input[type=file]").change(function() {
        var file =document.getElementById("file-input").value
        console.log(this.files[0].mozFullPath);
        const regex = /(?<=ams-)[0-9.]+|(?<=-)\d+/gm;
        var array=file.match(regex)
        document.getElementById("version").value=array[0]
        document.getElementById("revision").value=array[1]
    });
});

