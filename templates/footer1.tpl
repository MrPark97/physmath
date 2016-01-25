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
</body>
</html>