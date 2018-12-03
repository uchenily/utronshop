window.onload = function(){
    var msg = document.getElementById("message");
    var frm = document.getElementById("myform");
    frm.onclick = function() {
        if (msg != null) {
            msg.innerHTML = "";
        }
    }
}
