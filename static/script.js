function gracefulScroll(y) {
			var y0 = window.pageYOffset;
			if (y0 > y) {
				var ended = false;
				var a = 0.01;
				var frequency = 4;
				var n = 0;
				var interval = window.setInterval(function(){
					var t = n*frequency;
					s = a*t*t/2;
					if((y0 - s) < y) {scrollTo(0, y); window.clearInterval(interval);}
					else window.scrollTo(0, y0 - s);
					n++;
				}, frequency);
			} else if (y0 < y) {
				var ended = false;
				var a = 0.01;
				var frequency = 4;
				var n = 0;
				var interval = window.setInterval(function(){
					var t = n*frequency;
					s = a*t*t/2;
					if((y0 + s) > y) {scrollTo(0, y); window.clearInterval(interval);}
					else window.scrollTo(0, y0 + s);
					n++;
				}, frequency);
			}
		}

function sectionOpenAnimation(section, browser, href) {
		if(browser == 'moz') {
			setTimeout(function(){
				window.location.pathname = href;
			}, 1000);
		} else if(browser == 'webkit') {
			$('#'+section+'indexblock p').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'indexblock .button').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'frontimg').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "imgResize",
			});
			$('#'+section+'indexblock h1').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "hResize",
			});
			setTimeout(function(){
				$('#'+section+'indexblock p').remove();
				$('#'+section+'indexblock .button').remove();
				$('#'+section+'indexblock').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "openSection",
				});
				setTimeout(function(){
					$('#'+section+'mainblock').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "mainShow",
					});
					$('#'+section+'footer').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "footerShow",
					});
					setTimeout(function(){
						$(".indexblock").not("#"+section+"indexblock").remove();
						window.scrollTo(0, 0);
						window.location.pathname = href;
					}, 1000);
				}, 1000);
			}, 1000);
		}
		else {
			setTimeout(function(){
				window.location.pathname = href;
			}, 1000);
		}
}

function openScientists(section, browser) {
		if(browser == 'moz') {
					setTimeout(function(){
						window.location.pathname = "/scientist/"+scientist.name;
					}, 1000);
		} else if(browser == 'webkit') {
			$('#'+section+'indexblock p').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'indexblock .button').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'frontimg').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'indexblock h1').css({
				"-webkit-animation-duration" : "1s",
				"-webkit-animation-fill-mode" : "forwards",
				"-webkit-animation-name" : "hResize1",
			});
			setTimeout(function(){
				$('#'+section+'indexblock p').remove();
				$('#'+section+'indexblock .button').remove();
				$('#'+section+'frontimg').remove();
				$('#'+section+'indexblock').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "openSection",
				});
				$('#'+section+'indexblock h1').text(scientistHeader);
				$('#'+section+'indexblock h1').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeIn1",
				});
				setTimeout(function(){
					$('#'+section+'mainblock').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "mainShow",
					});
					$('#'+section+'footer').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "footerShow",
					});
					setTimeout(function(){
						$(".indexblock").not("#"+section+"indexblock").remove();
						window.scrollTo(0, 0);
						window.location.pathname = "/scientist/"+scientist.name;
					}, 1000);
				}, 1000);
			}, 1000);
		} else {
			setTimeout(function(){
				window.location.pathname = "/scientist/"+scientist.name;
			}, 1000);
		}
	}

function nextScientist(scientist, section, browser) {
			if(browser == 'moz') {
			$('#'+section+'indexblock h1').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeOut",
			});
			$('#'+section+'frontimg').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeOut",
			});
			$('#'+section+'indexblock p').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeOut",
			});
			setTimeout(function() {
				$('#'+section+'indexblock h1').text(scientist.name);
				$('#'+section+'indexblock p').html(scientist.quote);
				$('#'+section+'frontimg').attr("src", scientist.image);
				$('#'+section+'indexblock h1').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeIn",
				});
				$('#'+section+'frontimg').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeIn",
				});
				$('#'+section+'indexblock p').css({
					"-moz-animation-duration" : "1s",
					"-moz-animation-fill-mode" : "forwards",
					"-moz-animation-name" : "fadeIn",
				});
			}, 1000);
		} else if(browser == 'webkit') {
			$('#'+section+'indexblock h1').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'frontimg').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeOut",
			});
			$('#'+section+'indexblock p').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeOut",
			});
			setTimeout(function() {
				$('#'+section+'indexblock h1').text(scientist.name);
				$('#'+section+'indexblock p').html(scientist.quote);
				$('#'+section+'frontimg').attr("src", scientist.image);
				$('#'+section+'indexblock h1').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeIn",
				});
				$('#'+section+'frontimg').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeIn",
				});
				$('#'+section+'indexblock p').css({
					"-webkit-animation-duration" : "1s",
					"-webkit-animation-fill-mode" : "forwards",
					"-webkit-animation-name" : "fadeIn",
				});
			}, 1000);

} else {
	setTimeout(function() {
				$('#'+section+'indexblock h1').text(scientist.name);
				$('#'+section+'indexblock p').html(scientist.quote);
				$('#'+section+'frontimg').attr("src", scientist.image);
			}, 1000);
}
}

window.location.re