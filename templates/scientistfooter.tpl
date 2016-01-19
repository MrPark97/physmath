<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width">
		<title>Admin</title>
		<link rel="stylesheet" type="text/css" href="../static/styles/style1.css">
		<link rel="stylesheet" type="text/css" href="../static/styles/medium1.css" media="(min-width:320px) and (max-width:769px)">
		<script src="/static/classes.js"></script>
	</head>
<body>
<div class="indexblock" id="physindexblock">
			<h1>{{.Physh}}</h1>
			<img src="../static/images/physics.png" id="physfrontimg">
		</div>
		<div class="mainblock" id="physmainblock">
			<h2>{{.Physh}}</h2>
			<ul>
				<li onclick="request.send('theme=mechanics');" style="background-image: url('/static/images/mechanics.png');">{{.Mechanics}}</li>
				<li onclick="request.send('theme=thermodynamics');" style="background-image: url('/static/images/thermodynamics.png');">{{.Thermodynamics}}</li>
				<li onclick="request.send('theme=electrodynamics');" style="background-image: url('/static/images/electrodynamics.png');">{{.Electrodynamics}}</li>
				<li onclick="request.send('theme=oscillations');" style="background-image: url('/static/images/oscillations.png');">{{.Oscillations}}</li>
				<li onclick="request.send('theme=optics');" style="background-image: url('/static/images/optics.png');">{{.Optics}}</li>
				<li onclick="request.send('theme=atoms');" style="background-image: url('/static/images/atoms.png');">{{.Atoms}}</li>
				<li onclick="request.send('theme=hard');" style="background-image: url('/static/images/hard.png');">{{.Hard}}</li>
				<li onclick="request.send('theme=nucleus');" style="background-image: url('/static/images/nucleus.png');">{{.Nucleus}}</li>
			</ul>
		</div>
		<div class="footer" id="physfooter"><a href="#">&copy; MrPark 2015</a></div>
</body>
</html>