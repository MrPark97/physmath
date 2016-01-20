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
<body onresize="rePosH();">
<div class="indexblock" id="mathindexblock">
<h1>{{.Mathh}}</h1>
<img src="/static/images/math.png" id="mathfrontimg">
</div>
<div class="mainblock" id="mathmainblock">
		{{if (len .Article) gt 2 }}

			{{.Article}}
		{{else}}
			<h2>{{.Algebra}}</h2>
			<ul>
				<li style="background-image: url('/static/images/numbers.png');"><a href="/math/numbers">{{.Numbers}}</a></li>
				<li style="background-image: url('/static/images/algebraic.png');"><a href="math/algebraic">{{.Algebraic}}</a></li>
				<li style="background-image: url('/static/images/functions.png');"><a href="math/functions">{{.Functions}}</a></li>
				<li style="background-image: url('/static/images/transcendental.png');"><a href="math/transcendental">{{.Transcendental}}</a></li>
				<li style="background-image: url('/static/images/equations.png');"><a href="math/equations">{{.Equations}}</a></li>
				<li style="background-image: url('/static/images/inequalities.png');"><a href="math/inequalities">{{.Inequalities}}</a></li>
				<li style="background-image: url('/static/images/calculus.png');"><a href="math/calculus">{{.Calculus}}</a></li>
			</ul>
			<h2>{{.Geometry}}</h2>
			<ul>
				<li style="background-image: url('/static/images/plane.png');"><a href="math/plane">{{.Plane}}</a></li>
				<li style="background-image: url('/static/images/solidgeometry.png');"><a href="math/solidgeometry">{{.Solid}}</a></li>
				<li style="background-image: url('/static/images/coordinates.png');"><a href="math/coordinates">{{.Coordinates}}</a></li>
				<li style="background-image: url('/static/images/transformations.png');"><a href="math/transformations">{{.Transformations}}</a></li>
				<li style="background-image: url('/static/images/vectors.png');"><a href="math/vectors">{{.Vectors}}</a></li>
			</ul>
		{{end}}
</div>
<div class="footer" id="mathfooter"><a href="#">&copy; MrPark 2015</a></div>
{{.Addscript}}
</body>
</html>