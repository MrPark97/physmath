var g = 9.81;
var scale = 20;
var particle1, particle2;
var problem;
function Particle(r, name, x, y, v, alpha) {
	this.t = 0;
	this.alphadeg = alpha;
	if(alpha == 90) {
	this.v0 = new Vector(0, v);
	this.v = new Vector(0, v);
	}
	else {
	this.alpharad = alpha*0.017453293;
	this.v0 = new Vector(v * Math.cos(this.alpharad), v * Math.sin(this.alpharad));
	this.v = new Vector(v * Math.cos(this.alpharad), v * Math.sin(this.alpharad));
	}
	this.radius = r;
	this.name = name;
	if(this.v0.y > 0) this.dy = true;
	else this.dy = false;
	this.pos = new Vector(x, y);
	this.pos0 = new Vector(x, y);
	this.t0 = new Date().getTime();
}

function Problem(a) {
 this.particles = a;
 this.timecreate = new Date().getTime();
 this.status = this.finishDelay = null;
}

var maxStep = 0.05;
Problem.prototype.animate = function(step) {
 if (this.status != null) this.finishDelay -= step;
 while (step > 0) {
 var thisStep = Math.min(step, maxStep);
 this.particles.forEach(function(particle) {
 particle.act(this);
 }, this);
 step -= thisStep;
 }
};

Particle.prototype.act = function(problem) {
	var t = (new Date().getTime() - this.t0)/1000;
	if(this.dy) {
		this.v.y = this.v0.y - g*t;
		if(this.v.y < 0) {
			this.dy = false;

			this.pos0 = new Vector(this.pos.x, this.pos.y);
			this.t0 = new Date().getTime();
			t = (new Date().getTime() - this.t0)/1000;

			this.v0 = new Vector(this.v.x, this.v.y);

 			this.pos.y = this.pos0.y + this.v0.y*t - g*t*t/2;
 		} else {
 			this.pos.y = this.pos0.y + this.v0.y*t - g*t*t/2;
 		}
	} else {
 		this.v.y = this.v0.y - g*t;

 		if(this.v.y>0) {
	 		this.dy = true;

 			this.pos0 = new Vector(this.pos.x, this.pos.y);
 			this.t0 = new Date().getTime();
 			t = (new Date().getTime() - this.t0)/1000;

 			this.v0 = new Vector(this.v.x, this.v.y);

 			this.pos.y = this.pos0.y + this.v0.y*t - g*t*t/2;
 		} else {
 			this.pos.y = this.pos0.y + this.v0.y*t - g*t*t/2;
 		}
	}
	this.pos.x = this.pos0.x + this.v0.x*t;

if(this.pos.y <= 0 && problem.status===null) {problem.status = "lost"; problem.finishDelay = 1; /*if(this.t>10) problem.finishDelay = -1;*/}
//if(this.name == "B") document.getElementById("log").innerHTML += this.name+ " vy = "+ this.v.y +this.dy+ /*" vx = " + this.v.x + */" y =" + this.pos.y + /*" x = " + this.pos.x+*/ "<br>";
}

function Vector(x, y) {
 this.x = x; this.y = y;
}
Vector.prototype.plus = function(other) {
 return new Vector(this.x + other.x, this.y + other.y);
};
Vector.prototype.times = function(factor) {
 return new Vector(this.x * factor, this.y * factor);
};

Problem.prototype.isFinished = function() {
 return this.status != null && this.finishDelay < 0;
};

function runAnimation(frameFunc) {
 var lastTime = null;
 function frame(time) {
 var stop = false;
 if (lastTime != null) {
 var timeStep = Math.min(time - lastTime, 100) / 1000;
 stop = frameFunc(timeStep) === false;
 }
 lastTime = time;
 if (!stop)
 requestAnimationFrame(frame);
 }
 requestAnimationFrame(frame);
}

function runProblem(Display, problem) {
 var display = new Display(document.querySelector('#problemsmainblock'), problem);
 runAnimation(function(step) {
 problem.animate(step);
 display.drawFrame(step);

 if (problem.isFinished()) {
 	display.clear();
 	return false;
 }
 });
}

function CanvasDisplay(parent, problem) {
	this.canvas = document.createElement("canvas");
 	this.canvas.width = 600;
 	this.canvas.height = 450;
 	parent.appendChild(this.canvas);
	this.cx = this.canvas.getContext("2d");
	this.problem = problem;
	this.animationTime = 0;
	this.viewport = {
 		left: 0,
 		top: 0,
 		width: this.canvas.width / scale,
 		height: this.canvas.height / scale
 	};
 	this.drawFrame(0);
}

CanvasDisplay.prototype.clear = function() {
 this.canvas.parentNode.removeChild(this.canvas);
};

CanvasDisplay.prototype.drawFrame = function(step) {
 this.animationTime += step;
 this.clearDisplay();
 this.drawParticles();
};

CanvasDisplay.prototype.clearDisplay = function() {
 if (this.problem.status == "lost")
 this.cx.fillStyle = "deepskyblue";
 else this.cx.fillStyle = "white";
 this.cx.fillRect(0, 0, this.canvas.width, this.canvas.height);
};

CanvasDisplay.prototype.drawParticles = function() {
 this.problem.particles.forEach(function(particle) {
 var x = (particle.pos.x - this.viewport.left) * scale;
 var y = (this.canvas.height - particle.pos.y*scale - this.viewport.top);
if(y>0){ this.cx.beginPath();
 this.cx.arc(x, y, particle.radius*scale, 0, 7);
 this.cx.fillStyle = "black";
 this.cx.fill();
 this.cx.font = "16px Segoe UI";
 this.cx.fillStyle = "darkgreen";
 this.cx.fillText(particle.name, x+16, y-16);
 this.cx.strokeStyle = "red";
 this.cx.beginPath();
 this.cx.moveTo(x, y);
 var x1, y1;
 if(particle.dy && particle.v.y>=0) {x1 = x+particle.v.x*10; y1 = y-particle.v.y*10;}
 else if(!particle.dy && particle.v.y<=0) {x1 = x+particle.v.x*10; y1 = y-particle.v.y*10;}
 //if(particle.name=="А") alert(particle.dy);
 this.cx.lineTo(x1, y1);
 this.cx.font = "12px Segoe UI";
 this.cx.fillText("v", x1+12, y1-12);
 this.cx.stroke();
}
 }, this);
};