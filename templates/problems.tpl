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
<div class="indexblock" id="problemsindexblock">
			<h1>{{.Problemsh}}</h1>
			<img src="/static/images/problems.png" id="problemsfrontimg">
		</div>
		<div class="mainblock" id="problemsmainblock">
		{{if (len .Article) gt 2 }}

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
</body>
</html>