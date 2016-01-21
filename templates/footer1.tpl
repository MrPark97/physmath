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
		gracefulScroll(0);
		sectionOpenAnimation("math", browser, "/math");
	}
	
	function PhysicsOpen() {
		if(window.innerWidth > 769) {
			gracefulScroll(450);
		} else {
			gracefulScroll(234);
		}
		sectionOpenAnimation("phys", browser, "/physics");
	}

	function ProblemsOpen() {
		if(window.innerWidth > 769) {
			gracefulScroll(900);
		} else {
			gracefulScroll(460);
		}
		sectionOpenAnimation("problems", browser, "/problems");
	}

	function ScientistOpen() {
		if(window.innerWidth > 769) {
			gracefulScroll(1350);
		} else {
			gracefulScroll(693);
		}
		clearInterval(scientistChange);
		openScientists("scientist", browser);
	}
</script>
<script src="/static/ajax.js"></script>
</body>
</html>