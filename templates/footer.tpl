<div id="indexpage" class="indexblock">
	<h1>Physmath.uz</h1>
	<p>{{.Intended}}</p>
</div>
<div class="indexblock" id="mrpark">
	<h1>{{.Myname}}</h1>
	<p>{{.Myp}}</p>
	<p>{{.Myp1}}</p>
	<p><b>e-mail:</b> evgeniy.pak.97@mail.ru</p>
</div>
<div class="indexblock" id="langs">
	<h1>{{.Langh}}</h1>
	<div id="lang">
	<div id="RU" class="lang" onclick="document.cookie = 'lang=ru'; window.open(document.URL, '_self');"><h4>RU</h4></div>
	<div id="EN" class="lang" onclick="document.cookie = 'lang=en'; window.open(document.URL, '_self');"><h4>EN</h4></div>
	</div>
</div>
<script src="/static/script.js"></script>
<script>
	var browser;
	var scientistHeader = {{.ScientistHeader}};
	if(navigator["mozGetUserMedia"]) browser = "moz";
	else if(navigator["webkitGetUserMedia"]) browser = "webkit";

	function Scientist(name, image, quote) {
		this.name = name;
		this.image = image;
		this.quote = quote;
	}
</script>
