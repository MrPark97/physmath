<script>	
	var n = 0;
	var scientist = scientists[0];
	nextScientist(scientist, "scientist", browser);
	var scientistChange = setInterval(function() {
		scientist = scientists[n+1];
		nextScientist(scientist, "scientist", browser);
		if((n+2) < scientists.length) {
			++n;
		} else {
			n = -1;
		}
	}, 5000);
	function MathOpen() {
		gracefulScroll(document.querySelector("#mathindexblock").offsetTop);
		sectionOpenAnimation("math", browser, "/math");
	}
	
	function PhysicsOpen() {
		gracefulScroll(document.querySelector("#physindexblock").offsetTop);
		sectionOpenAnimation("phys", browser, "/physics");
	}

	function ProblemsOpen() {
		gracefulScroll(document.querySelector("#problemsindexblock").offsetTop);
		sectionOpenAnimation("problems", browser, "/problems");
	}

	function ScientistOpen() {
		gracefulScroll(document.querySelector("#scientistindexblock").offsetTop);
		clearInterval(scientistChange);
		openScientists("scientist", browser);
	}
</script>
<script src="/static/ajax.js"></script>
<!-- START WWW.UZ TOP-RATING --><SCRIPT language="javascript" type="text/javascript">
<!--
top_js="1.0";top_r="id=35268&r="+escape(document.referrer)+"&pg="+escape(window.location.href);document.cookie="smart_top=1; path=/"; top_r+="&c="+(document.cookie?"Y":"N")
//-->
</SCRIPT>
<SCRIPT language="javascript1.1" type="text/javascript">
<!--
top_js="1.1";top_r+="&j="+(navigator.javaEnabled()?"Y":"N")
//-->
</SCRIPT>
<SCRIPT language="javascript1.2" type="text/javascript">
<!--
top_js="1.2";top_r+="&wh="+screen.width+'x'+screen.height+"&px="+
(((navigator.appName.substring(0,3)=="Mic"))?screen.colorDepth:screen.pixelDepth)
//-->
</SCRIPT>
<SCRIPT language="javascript1.3" type="text/javascript">
<!--
top_js="1.3";
//-->
</SCRIPT>
<SCRIPT language="JavaScript" type="text/javascript">
<!--
top_rat="&col=340F6E&t=ffffff&p=BD6F6F";top_r+="&js="+top_js+"";document.write('<img src="http://cnt0.www.uz/counter/collect?'+top_r+top_rat+'" width=0 height=0 border=0 />')//-->
</SCRIPT><NOSCRIPT><IMG height=0 src="http://cnt0.www.uz/counter/collect?id=35268&pg=http%3A//uzinfocom.uz&col=340F6E&t=ffffff&p=BD6F6F" width=0 border=0 /></NOSCRIPT><!-- FINISH WWW.UZ TOP-RATING -->
</body>
</html>