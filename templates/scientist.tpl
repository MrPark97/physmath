<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width">
		<title>PHYSMATH.uz</title>
		<link rel="stylesheet" type="text/css" href="../static/styles/style1.css">
		<link rel="stylesheet" type="text/css" href="../static/styles/medium1.css" media="(min-width:320px) and (max-width:769px)">
		<script src="/static/classes.js"></script>
	</head>
<body>
<div class="indexblock" id="scientistindexblock">
			<h1>{{.ScientistH}}</h1>
		</div>
		<div class="mainblock" id="scientistmainblock">
		<h2>{{.Name}}</h2>
		<blockquote>{{.Quote}}</blockquote>
		<img src="/static/scientists/{{.Img}}">
		{{.Article}}
		</div>
		<div class="footer" id="scientistfooter"><a href="#">&copy; MrPark 2015</a></div>
<script id="MathJax" defer type="text/javascript" src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
</script>
<script>
	if(navigator["mozGetUserMedia"]) {
		var mj = document.getElementById('MathJax');
		document.body.removeChild(mj);
	}
</script>
</body>
</html>