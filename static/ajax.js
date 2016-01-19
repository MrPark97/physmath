request = new ajaxRequest();
request.open("POST", "/posthandler", true);
request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
request.setRequestHeader("Connection", "close");
		
request.onreadystatechange = function() {
	if (this.readyState == 4) {
		if (this.status == 200) {
			if (this.responseText != null) {
				$(".mainblock").html(this.responseText);
			} else alert("Ошибка Ajax: " + this.statusText);
		} else alert("Ошибка Ajax: " + this.statusText);
	}
}

function ajaxRequest() {
	try {
		var request = new XMLHttpRequest();
	} catch(e1) {
		try {
			request = new ActiveXObject("Msxml2.XMLHTTP");
		} catch(e2) {
			try {
				request = new ActiveXObject("Microsoft.XMLHTTP");
			} catch(e3) {
				request = false;
			}
		}
	}
	return request;
}