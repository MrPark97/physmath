<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width">
		<title>PHYSMATH.uz</title>
		<link rel="stylesheet" type="text/css" href="/static/styles/style1.css">
		<link rel="stylesheet" type="text/css" href="/static/styles/small1.css" media="(min-width:320px) and (max-width:769px)">
		<script src="/static/classes.js"></script>
		<script src="/static/jquery.js"></script>
		<script type="text/javascript" id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js">
		</script>
	</head>
<body>
</div>
<div class="indexblock" id="physindexblock">
	<h1>{{.Physh}}</h1>
	<img src="/static/images/physics.svg" id="physfrontimg">
</div>
<div class="mainblock" id="physmainblock">
		{{if gt (len .Article) 2 }}

			{{.Article}}
		{{else}}
		<h2>{{.Physh}}</h2>
		<ul>
			<li style="background-image: url('/static/images/mechanics.png');"><a href="/physics/mechanics">{{.Mechanics}}</a></li>
			<li style="background-image: url('/static/images/thermodynamics.png');"><a href="/physics/thermodynamics">{{.Thermodynamics}}</a></li>
			<li style="background-image: url('/static/images/electrodynamics.png');"><a href="/physics/electrodynamics">{{.Electrodynamics}}</a></li>
			<li style="background-image: url('/static/images/oscillations.png');"><a href="/physics/oscillations">{{.Oscillations}}</a></li>
			<li style="background-image: url('/static/images/optics.png');"><a href="/physics/optics">{{.Optics}}</a></li>
			<li style="background-image: url('/static/images/atoms.png');"><a href="/physics/atoms">{{.Atoms}}</a></li>
			<li style="background-image: url('/static/images/hard.png');"><a href="/physics/hard">{{.Hard}}</a></li>
			<li style="background-image: url('/static/images/nucleus.png');"><a href="/physics/nucleus">{{.Nucleus}}</a></li>
			</ul>
		{{end}}
</div>
<script>
	window.MathJax = {
		loader: {load: ["input/tex", "output/chtml"]}
	};
</script>
<div class="footer" id="physfooter"><a href="#">&copy; MrPark 2015</a></div>
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
top_rat="&col=340F6E&t=ffffff&p=BD6F6F";top_r+="&js="+top_js+"";document.write('<img src="http://cnt0.www.uz/counter/collect?'+top_r+top_rat+'" width=0 height=0 border=0 style="display: none;"/>')//-->
</SCRIPT><NOSCRIPT><IMG height=0 src="http://cnt0.www.uz/counter/collect?id=35268&pg=http%3A//uzinfocom.uz&col=340F6E&t=ffffff&p=BD6F6F" width=0 border=0 /></NOSCRIPT><!-- FINISH WWW.UZ TOP-RATING -->  
</body>
</html>