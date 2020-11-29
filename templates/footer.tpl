<div id="indexpage" class="indexblock">
	<div class="indexblock-info">
		<h1>Physmath.uk</h1>
		<p>{{.Intended}}</p>
	</div>
</div>
<div class="indexblock" id="mrpark">
	<div class="indexblock-info">
		<h1>{{.Myname}}</h1>
		<p>{{.Myp}}</p>
		<p>{{.Myp1}}</p>
		<p><b>e-mail:</b> evgeniy.pak.97@gmail.com</p>
	</div>
</div>
<div class="indexblock" id="langs">
	<div class="indexblock-info">
		<h1>{{.Langh}}</h1>
		<div id="lang">
		<div id="RU" class="lang" onclick="document.cookie = 'lang=ru'; window.open(document.URL, '_self');"><h4>RU</h4></div>
		<div id="EN" class="lang" onclick="document.cookie = 'lang=en'; window.open(document.URL, '_self');"><h4>EN</h4></div>
		</div>
	</div>
</div>
<script src="/static/script.js"></script>
<script>
	var browser;
	var scientistHeader = "{{.ScientistHeader}}";
	if(navigator["mozGetUserMedia"]) browser = "moz";
	else if(navigator["webkitGetUserMedia"]) browser = "webkit";

	function Scientist(name, image, quote) {
		this.name = name;
		this.image = image;
		this.quote = quote;
	}
</script>
