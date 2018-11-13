window.onload = function() {
    getMethod();
}

var getMethod = function() {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if(xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            var text = xmlhttp.responseText;
            alert("recv_msg: " + text);
        }
    }
    xmlhttp.open("GET", "/getInfo", true)
    xmlhttp.send()
}