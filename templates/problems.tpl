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
	</head>
<body>
</div>
<div class="indexblock" id="problemsindexblock">
			<img src="/static/images/problems.png" id="problemsfrontimg">
			<h1>{{.Problemsh}}</h1>
		</div>
		<div class="mainblock" id="problemsmainblock">
		{{if gt (len .Article) 2 }}

			{{.Article}}
		{{else}}
			<h2>{{.Problemsh}}</h2>
			<ul>
				<li style="background-image: url('/static/images/mechanics.png');"><a href="problems/mechanics">{{.Mechanics}}</a></li>
				<li style="background-image: url('/static/images/thermodynamics.png');"><a href="problems/thermodynamics">{{.Thermodynamics}}</a></li>
				<li style="background-image: url('/static/images/electrodynamics.png');"><a href="problems/electrodynamics">{{.Electrodynamics}}</a></li>
			</ul>
			{{end}}
		</div>
		<div class="footer" id="problemsfooter"><a href="#">&copy; MrPark 2015</a></div>
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