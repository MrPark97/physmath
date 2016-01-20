<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width">
		<title>PHYSMATH.uz</title>
		<link rel="stylesheet" type="text/css" href="/static/styles/style1.css">
		<link rel="stylesheet" type="text/css" href="/static/styles/medium1.css" media="(min-width:320px) and (max-width:769px)">
		<script src="/static/classes.js"></script>
		<script src="/static/jquery.js"></script>
	</head>
<body>
</div>
<div class="indexblock" id="physindexblock">
	<h1>{{.Physh}}</h1>
	<img src="/static/images/physics.png" id="physfrontimg">
</div>
<div class="mainblock" id="physmainblock">
		{{if (len .Article) gt 2 }}

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
<div class="footer" id="physfooter"><a href="#">&copy; MrPark 2015</a></div>
</body>
</html>