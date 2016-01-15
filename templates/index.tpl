<div class="indexblock" id="mathindexblock">
	<h1>{{.Mathh}}</h1>
			<p>{{.Mathp}}</p>
			<div class="button" onclick="MathOpen();"><h6>{{.Readmore}}</h6></div>
			<img src="/static/images/math.png" id="mathfrontimg">
		</div>
		<div class="mainblock" id="mathmainblock">
			<h2>{{.Algebra}}</h2>
			<ul>
				<li onclick="request.send('theme=numbers');" style="background-image: url('/static/images/numbers.png');">{{.Numbers}}</li>
				<li onclick="request.send('theme=algebraic'); "style="background-image: url('/static/images/algebraic.png');">{{.Algebraic}}</li>
				<li onclick="request.send('theme=functions');" style="background-image: url('/static/images/functions.png');">{{.Functions}}</li>
				<li onclick="request.send('theme=transcendental');" style="background-image: url('/static/images/transcendental.png');">{{.Transcendental}}</li>
				<li onclick="request.send('theme=equations');" style="background-image: url('/static/images/equations.png');">{{.Equations}}</li>
				<li onclick="request.send('theme=inequalities');" style="background-image: url('/static/images/inequalities.png');">{{.Inequalities}}</li>
				<li onclick="request.send('theme=calculus');" style="background-image: url('/static/images/calculus.png');">{{.Calculus}}</li>
			</ul>
			<h2>{{.Geometry}}</h2>
			<ul>
				<li onclick="request.send('theme=plane');" style="background-image: url('/static/images/plane.png');">{{.Plane}}</li>
				<li onclick="request.send('theme=solidgeometry');" style="background-image: url('/static/images/solidgeometry.png');">{{.Solid}}</li>
				<li onclick="request.send('theme=coordinates');" style="background-image: url('/static/images/coordinates.png');">{{.Coordinates}}</li>
				<li onclick="request.send('theme=transformations');" style="background-image: url('/static/images/transformations.png');">{{.Transformations}}</li>
				<li onclick="request.send('theme=vectors');" style="background-image: url('/static/images/vectors.png');">{{.Vectors}}</li>
			</ul>
		</div>
		<div class="footer" id="mathfooter"><a href="#">&copy; MrPark 2015</a></div>
		<div class="indexblock" id="physindexblock">
			<h1>{{.Physh}}</h1>
			<p>{{.Physp}}</p>
			<div class="button" onclick="PhysicsOpen();"><h6>{{.Readmore}}</h6></div>
			<img src="/static/images/physics.png" id="physfrontimg">
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
		<div class="indexblock" id="problemsindexblock">
			<h1>{{.Problemsh}}</h1>
			<p>{{.Problemsp}}</p>
			<div class="button" onclick="ProblemsOpen();"><h6>{{.Readmore}}</h6></div>
			<img src="/static/images/problems.png" id="problemsfrontimg">
		</div>
		<div class="mainblock" id="problemsmainblock">
			<h2>{{.Problemsh}}</h2>
			<ul>
				<li onclick="request.send('problems=mechanics');" style="background-image: url('/static/images/mechanics.png');">{{.Mechanics}}</li>
				<li onclick="request.send('problems=thermodynamics');" style="background-image: url('/static/images/thermodynamics.png');">{{.Thermodynamics}}</li>
				<li onclick="request.send('problems=electrodynamics');" style="background-image: url('/static/images/electrodynamics.png');">{{.Electrodynamics}}</li>
			</ul>
		</div>
		<div class="footer" id="problemsfooter"><a href="#">&copy; MrPark 2015</a></div>
		<div class="indexblock" id="scientistindexblock">
			<h1>Einstein</h1>
			<p><i>"As far as the laws of mathematics refer to reality, they are not certain, and as far as they are certain, they do not refer to reality."</i></p>
			<div class="button" onclick="ScientistOpen(); request.send('scientist='+scientist.name);"><h6>{{.Readmore}}</h6></div>
			<img src="/static/scientists/einstein.jpg" id="scientistfrontimg">
		</div>
		<div class="mainblock" id="scientistmainblock">
		</div>
		<div class="footer" id="scientistfooter"><a href="#">&copy; MrPark 2015</a></div>