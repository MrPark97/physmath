package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"text/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	physdb     = os.Getenv("DATABASE_URL")
	mathdb     = os.Getenv("DATABASE_URL")
	physdbname = "u233172724_phys"
	mathdbname = "u233172724_math"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if physdb == "" {
		physdb = "postgres://postgres:abracadabra_2016!@localhost:5432/physmath?sslmode=disable"
	}
	
	if mathdb == "" {
		mathdb = "postgres://postgres:abracadabra_2016!@localhost:5432/physmath?sslmode=disable"
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/scientist/", scientistHandler)
	http.HandleFunc("/posthandler", PostHandler)
	http.HandleFunc("/math", MathIndex)
	http.HandleFunc("/math/", MathInfo)
	http.HandleFunc("/physics", PhysIndex)
	http.HandleFunc("/physics/", PhysInfo)
	http.HandleFunc("/problems", ProblemIndex)
	http.HandleFunc("/google227af1ab1e0f26ee.html", GoogleWebMaster)
	http.HandleFunc("/problems/", ProblemInfo)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	headerOut(w)
	IndexPrint(w, r)
	footerOut(w, r)
}

func GoogleWebMaster(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/google227af1ab1e0f26ee.html")
	if err != nil {
		log.Println("template error")
	}
	data := struct {}{}
	err = t.Execute(w, data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func headerOut(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/header.tpl")
	if err != nil {
		log.Println("template error")
	}
	data := struct {
		Title string
	}{
		`PHYSMATH.uz`,
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func footerOut(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Intended, Myname, Myp, Myp1, Langh, ScientistHeader string
	}

	var Data data

	lang, err := r.Cookie("lang")
	if lang != nil && lang.Value == `en` {
		Data.Intended = `This site is for anyone looking to learn a little more about physics and mathematics. The site was created primarily for you, the users, so you will not see annoying ads and insanely jumping flash banners or GIFs. The design and program code of this site were created from scratch, without the use of various engines and copying the program code from the vastness of the worldwide network, by a student of the Tashkent Professional College of Information Technologies. I hope you enjoy this site and appreciate it. I wish you to find useful information in the bowels of this resource's web pages.`
		Data.Myname = `Park Eugene`
		Data.Myp = `Park Eugene (16.05.1997, Tashkent - ...) - programmer.`
		Data.Myp1 = `Mr.Park was born 16 May 1997. In 2004-2013 learned in specialized school 103.
			In 2013 has entered in Tashkent Professional College of Information Technologies, profession - a technician of software and multimedia systems.`
		Data.Langh = `Language`
		Data.ScientistHeader = `Scientists`
	} else {
		Data.Intended = `Этот сайт предназначен для всех, кто хочет узнать немного больше о физике и математике. Сайт создан в первую очередь для вас, пользователей, поэтому Вы не увидите надоедливой рекламы и безумно дрыгающихся флэш-баннеров или гифок. Дизайн и программный код данного сайта были созданы с нуля, без использования различных движков и копирования программного кода с просторов всемирной сети, студентом Ташкентского Профессионального Колледжа Информационных Технологий. Надеюсь, что Вам понравится данный сайт и Вы оцените его по достоинству. Желаю Вам найти полезную информацию в недрах веб-страничек данного ресурса.`
		Data.Myname = `Пак Евгений`
		Data.Myp = `Пак Евгений Герасимович (16.05.1997, Ташкент - ...) - программист.`
		Data.Myp1 = `Е. Г. Пак родился 16 мая 1997. В 2004-2013 учился в специализированной школе №103. 
			В 2013 поступил в Ташкентский Профессиональный Колледж Информационных Технологий, специальность - техник по программному обеспечению и мультимедийным системам.`
		Data.Langh = `Язык`
		Data.ScientistHeader = `Учёные`
	}

	t, err := template.ParseFiles("templates/footer.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}

	fmt.Fprint(w, `<script>`)
	scarray := `var scientists = [`

	db, err := sql.Open("postgres", physdb)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT scientist,scientist_ru,quote,quote_ru,img FROM scientists WHERE quote != '' ORDER BY scientist")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var scientist, scientist_ru, quote, quote_ru, img string

	for rows.Next() {
		err = rows.Scan(&scientist, &scientist_ru, &quote, &quote_ru, &img)

		if lang != nil && lang.Value == `en` {
			fmt.Fprint(w, scientist, " = new Scientist('", scientist, "', '/static/scientists/", img, "', '<i>\"", quote, "\"</i>', 'scientist');\n")
		} else {
			fmt.Fprint(w, scientist, " = new Scientist('", scientist_ru, "', '/static/scientists/", img, "', '<i>\"", quote_ru, "\"</i>', 'scientist');\n")
		}
		scarray += scientist + ","
	}
	rows.Close()

	scarray += `];`
	fmt.Fprint(w, scarray)
	fmt.Fprint(w, `</script>`)

	t, err = template.ParseFiles("templates/footer1.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func IndexPrint(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Mathh, Addscript, Mathp, Readmore, Physh, Physp, Problemsh, Problemsp, Geometry, Numbers, Algebraic, Algebra, Functions, Transcendental, Equations, Inequalities, Calculus, Plane, Solid, Coordinates, Transformations, Vectors, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
	}

	var (
		Data      data
		Addscript string
	)

	lang, err := r.Cookie("lang")
	if lang != nil && lang.Value == `en` {
		Data.Mathh = "Math"
		Addscript = ""
		Data.Mathp = "Mathematics lessons by M. Vygodsky. General basics and concepts of mathematics. Arithmetics, algebra, geometry and basics of calculus."
		Data.Readmore = "READ MORE"
		Data.Physh = "Physics"
		Data.Physp = "Physics lessons by L. D. Landau, A. I. Akhiezer, E. M. Lifshitz, N. Koshkin, M. Shirkevich. General basics and concepts of physics."
		Data.Problemsh = "Problems"
		Data.Problemsp = "Physics and mathematics problems. Several problems have animated canvas. Physics problems by L. Tarasov, A. Tarasova."

		Data.Geometry = "Geometry"
		Data.Numbers = "Numbers"
		Data.Algebraic = "Algebraic expressions"
		Data.Algebra = "Algebra and analysis begin"
		Data.Functions = "Functions and graphs"
		Data.Transcendental = "Transcendental expressions"
		Data.Equations = "Equations and systems of equations"
		Data.Inequalities = "Inequalities"
		Data.Calculus = "Elements of calculus"
		Data.Plane = "Geometric figures on plane"
		Data.Solid = "Solid geometry"
		Data.Coordinates = "Cartesian system of coordinates"
		Data.Transformations = "Transformations of the figures"
		Data.Vectors = "Vectors"

		Data.Mechanics = "Mechanics"
		Data.Thermodynamics = "Bases of molecular physics and thermodynamics"
		Data.Electrodynamics = "Electrodynamics"
		Data.Oscillations = "Oscillations and waves"
		Data.Atoms = "Physics of atoms and molecules"
		Data.Optics = "Optics"
		Data.Hard = "Bases of solids physics"
		Data.Nucleus = "Atomic nucleus and elementary particles physics"
	} else {
		Data.Mathh = "Математика"
		Data.Mathp = "Уроки математики от М. Выгодского. Общие начала и принципы математики. Арифметика, алгебра, геометрия и начала анализа."
		Data.Readmore = "ПОДРОБНЕЕ"
		Data.Physh = "Физика"
		Data.Physp = "Уроки физики от Л. Д. Ландау, А. И. Ахиезера, Е, М. Лифшица, Н. Кошкина, М. Ширкевича. Общие начала и принципы физики."
		Data.Problemsh = "Задачи"
		Data.Problemsp = "Задачи по физике и математике. В некоторых задачах присутствует анимированый холст. Задачи по физике от Л. Тарасова, А. Тарасовой."

		Data.Geometry = "Геометрия"
		Data.Numbers = "Числа"
		Data.Algebraic = "Алгебраические выражения"
		Data.Algebra = "Алгебра и начала анализа"
		Data.Functions = "Функции и графики"
		Data.Transcendental = "Трансцендентые выражения"
		Data.Equations = "Уравнения и системы уравнений"
		Data.Inequalities = "Неравенства"
		Data.Calculus = "Элементы математического анализа"
		Data.Plane = "Геометрические фигуры на плоскости"
		Data.Solid = "Стереометрия"
		Data.Coordinates = "Декартовы координаты"
		Data.Transformations = "Преобразования фигур"
		Data.Vectors = "Векторы"

		Data.Mechanics = "Механика"
		Data.Thermodynamics = "Основы молекулярной физики и термодинамики"
		Data.Electrodynamics = "Электродинамика"
		Data.Oscillations = "Колебания и волны"
		Data.Atoms = "Физика атомов и молекул"
		Data.Optics = "Оптика"
		Data.Hard = "Основы физики твердого тела"
		Data.Nucleus = "Физика ядра и элементарных частиц"

		Addscript = `
			<script>
			function rePosH() {
				if(window.innerWidth > 769) {
					$("#mathindexblock h1").css("right", "520px");
					$("#problemsindexblock h1").css("right", "340px");
				} else {
					$("#mathindexblock h1").css("right", "260px");
					$("#problemsindexblock h1").css("right", "170px");	
				}
			}
			rePosH();
			</script>
		`
	}

	t, err := template.ParseFiles("templates/index.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
	fmt.Fprint(w, Addscript)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if theme := r.PostFormValue("theme"); theme != "" {
		themeHandler(w, r, theme)
	}
	if name := r.PostFormValue("name"); name != "" {
		nameHandler(w, r, name, r.PostFormValue("articletheme"))
	}
	if problems := r.PostFormValue("problems"); problems != "" {
		problemHandler(w, r, problems)
	}
	if problemname := r.PostFormValue("problemname"); problemname != "" {
		problemNameHandler(w, r, problemname)
	}
}

func DbNameAndHeaderReturn(r *http.Request, theme string) (dbname, header string) {
	lang, _ := r.Cookie("lang")
	if lang != nil && lang.Value == `en` {
		switch theme {
		case `mechanics`:
			dbname = `u233172724_phys`
			header = `MECHANICS`
		case `numbers`:
			dbname = `u233172724_math`
			header = `NUMBERS`
		case `electrodynamics`:
			dbname = `u233172724_phys`
			header = `ELECTRODYNAMICS`
		case `thermodynamics`:
			dbname = `u233172724_phys`
			header = `BASES OF MOLECULAR PHYSICS AND THERMODYNAMICS`
		case `algebraic`:
			dbname = `u233172724_math`
			header = `ALGEBRAIC EXPRESSIONS`
		case `oscillations`:
			dbname = `u233172724_phys`
			header = `OSCILLATIONS AND WAVES`
		case `atoms`:
			dbname = `u233172724_phys`
			header = `PHYSICS OF ATOMS AND MOLECULES`
		case `optics`:
			dbname = `u233172724_phys`
			header = `OPTICS`
		case `hard`:
			dbname = `u233172724_phys`
			header = `BASES OF SOLIDS PHYSICS`
		case `nucleus`:
			dbname = `u233172724_phys`
			header = `ATOMIC NUCLEUS AND ELEMENTARY PARTICLES PHYSICS`
		case `plane`:
			dbname = `u233172724_math`
			header = `GEOMETRIC FIGURES ON PLANE`
		case `solidgeometry`:
			dbname = `u233172724_math`
			header = `SOLID GEOMETRY`
		case `coordinates`:
			dbname = `u233172724_math`
			header = `CARTESIAN SYSTEM OF COORDINATES`
		case `transformations`:
			dbname = `u233172724_math`
			header = `TRANSFORMATIONS OF THE FIGURES`
		case `vectors`:
			dbname = `u233172724_math`
			header = `VECTORS`
		case `functions`:
			dbname = `u233172724_math`
			header = `FUNCTIONS AND GRAPHS`
		case `transcendental`:
			dbname = `u233172724_math`
			header = `TRANSCENDENTAL EXPRESSIONS`
		case `equations`:
			dbname = `u233172724_math`
			header = `EQUATIONS AND SYSTEMS OF EQUATIONS`
		case `inequalities`:
			dbname = `u233172724_math`
			header = `INEQUALITIES`
		case `calculus`:
			dbname = `u233172724_math`
			header = `ELEMENTS OF CALCULUS`
		}
	} else {
		switch theme {
		case `mechanics`:
			dbname = `u233172724_phys`
			header = `МЕХАНИКА`
		case `numbers`:
			dbname = `u233172724_math`
			header = `ЧИСЛА`
		case `electrodynamics`:
			dbname = `u233172724_phys`
			header = `ЭЛЕКТРОДИНАМИКА`
		case `thermodynamics`:
			dbname = `u233172724_phys`
			header = `ОСНОВЫ МОЛЕКУЛЯРНОЙ ФИЗИКИ И ТЕРМОДИНАМИКИ`
		case `algebraic`:
			dbname = `u233172724_math`
			header = `АЛГЕБРАИЧЕСКИЕ ВЫРАЖЕНИЯ`
		case `oscillations`:
			dbname = `u233172724_phys`
			header = `КОЛЕБАНИЯ И ВОЛНЫ`
		case `atoms`:
			dbname = `u233172724_phys`
			header = `ФИЗИКА АТОМОВ И МОЛЕКУЛ`
		case `optics`:
			dbname = `u233172724_phys`
			header = `ОПТИКА`
		case `hard`:
			dbname = `u233172724_phys`
			header = `ОСНОВЫ ФИЗИКИ ТВЕРДОГО ТЕЛА`
		case `nucleus`:
			dbname = `u233172724_phys`
			header = `ФИЗИКА ЯДРА И ЭЛЕМЕНТАРНЫХ ЧАСТИЦ`
		case `plane`:
			dbname = `u233172724_math`
			header = `ГЕОМЕТРИЧЕСКИЕ ФИГУРЫ НА ПЛОСКОСТИ`
		case `solidgeometry`:
			dbname = `u233172724_math`
			header = `СТЕРЕОМЕТРИЯ`
		case `coordinates`:
			dbname = `u233172724_math`
			header = `ДЕКАРТОВЫ КООРДИНАТЫ`
		case `transformations`:
			dbname = `u233172724_math`
			header = `ПРЕОБРАЗОВАНИЯ ФИГУР`
		case `vectors`:
			dbname = `u233172724_math`
			header = `ВЕКТОРЫ`
		case `functions`:
			dbname = `u233172724_math`
			header = `ФУНКЦИИ И ГРАФИКИ`
		case `transcendental`:
			dbname = `u233172724_math`
			header = `ТРАНСЦЕНДЕНТНЫЕ ВЫРАЖЕНИЯ`
		case `equations`:
			dbname = `u233172724_math`
			header = `УРАВНЕНИЯ И СИСТЕМЫ УРАВНЕНИЙ`
		case `inequalities`:
			dbname = `u233172724_math`
			header = `НЕРАВЕНСТВА`
		case `calculus`:
			dbname = `u233172724_math`
			header = `ЭЛЕМЕНТЫ МАТЕМАТИЧЕСКОГО АНАЛИЗА`
		}
	}
	return
}

func themeHandler(w http.ResponseWriter, r *http.Request, theme string) {
	dbname, header := DbNameAndHeaderReturn(r, theme)
	lang, err := r.Cookie("lang")
	var namelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
	} else {
		namelg = `name`
	}

	var dbcons string
	if dbname == physdbname {
		dbcons = physdb
	} else {
		dbcons = mathdb
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + namelg + " FROM " + theme)
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	fmt.Fprint(w, `<h2><img src="/static/images/`, theme, `.png">`, header, `</h2>`)
	fmt.Fprint(w, `<ul class="sectionul">`)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Fprint(w, `<li onclick = "request.open('POST', '/posthandler', true); request.setRequestHeader('Content-type', 'application/x-www-form-urlencoded'); request.setRequestHeader('Connection', 'close'); request.send('name=`, name, `&articletheme=`, theme, `');">`, name, `</li>`)
	}
	rows.Close()
	fmt.Fprint(w, `</ul>`)
}

func themeHandler1(r *http.Request, theme string) string {
	dbname, header := DbNameAndHeaderReturn(r, theme)
	lang, err := r.Cookie("lang")
	var namelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
	} else {
		namelg = `name`
	}

	var dbcons string
	var sechref string
	if dbname == physdbname {
		dbcons = physdb
		sechref = "/physics/"
	} else {
		dbcons = mathdb
		sechref = "/math/"
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + namelg + " FROM " + theme)
	if err != nil {
		log.Println("scientists are not founded", err)
	}
	var result string
	result += `<h2><img src="/static/images/`+theme+`.png">`+header+`</h2>`
	result += `<ul class="sectionul">`
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		result+=`<li><a href="`+sechref+theme+`/`+name+`">`+name+`</li>`
	}
	rows.Close()
	result+=`</ul>`
	return result
}

func nameHandler(w http.ResponseWriter, r *http.Request, name, theme string) {
	dbname, header := DbNameAndHeaderReturn(r, theme)
	lang, err := r.Cookie("lang")
	var namelg, authorlg, articlelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		authorlg = `author_en`
		articlelg = `article_en`
	} else {
		namelg = `name`
		authorlg = `author`
		articlelg = `article`
	}

	var dbcons string
	if dbname == physdbname {
		dbcons = physdb
	} else {
		dbcons = mathdb
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + articlelg + "," + authorlg + " FROM " + theme + " WHERE " + namelg + " = '" + name + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var article, author string
	for rows.Next() {
		err = rows.Scan(&article, &author)
		fmt.Fprint(w, `<div class="articleinfo"><div class="articlesection"><h3><img src="/static/images/`, theme, `.png">`, header, `</h3></div><div class="articleauthor"><h3>`, author, `</h3></div></div>`)
		fmt.Fprint(w, "<h2>", name, "</h2>")
		fmt.Fprint(w, article)
	}
	rows.Close()
}

func nameHandler1(r *http.Request, theme, name string) string {
	dbname, header := DbNameAndHeaderReturn(r, theme)
	lang, err := r.Cookie("lang")
	var namelg, authorlg, articlelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		authorlg = `author_en`
		articlelg = `article_en`
	} else {
		namelg = `name`
		authorlg = `author`
		articlelg = `article`
	}

	var dbcons string
	if dbname == physdbname {
		dbcons = physdb
	} else {
		dbcons = mathdb
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + articlelg + "," + authorlg + " FROM " + theme + " WHERE " + namelg + " = '" + name + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var article, author, result string
	if rows.Next() {
		err = rows.Scan(&article, &author)
		result+=`<div class="articleinfo"><div class="articlesection"><h3><img src="/static/images/`+theme+`.png">`+header+`</h3></div><div class="articleauthor"><h3>`+author+`</h3></div></div>`
		result+="<h2>"+name+"</h2>"
		result+=article
	}
	rows.Close()
	return result
}

func scientistHandler(w http.ResponseWriter, r *http.Request) {
	scientist := r.URL.Path[len("/scientist/"):]
	lang, err := r.Cookie("lang")
	var namelg, scientistlg, articlelg, quotelg, scientists string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		scientistlg = `scientist`
		articlelg = `article_en`
		quotelg = `quote`
		scientists = `Scientists`
	} else {
		namelg = `name`
		scientistlg = `scientist_ru`
		articlelg = `article`
		quotelg = `quote_ru`
		scientists = `Учёные`
	}

	dbcons := physdb

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + articlelg + "," + namelg + "," + quotelg + ",img FROM scientists WHERE " + scientistlg + " = '" + scientist + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var article, name, quote, img string
	for rows.Next() {
		err = rows.Scan(&article, &name, &quote, &img)
	}
	rows.Close()

	t, err := template.ParseFiles("templates/scientist.tpl")
	if err != nil {
		log.Println("template error")
	}
	data := struct {
		Name, Article, Img, ScientistH, Quote string
	}{
		name,
		article,
		img,
		scientists,
		quote,
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func problemHandler(w http.ResponseWriter, r *http.Request, problems string) {
	dbname, header := DbNameAndHeaderReturn(r, problems)
	lang, err := r.Cookie("lang")
	var namelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
	} else {
		namelg = `name`
	}

	var dbcons string
	if dbname == physdbname {
		dbcons = physdb
	} else {
		dbcons = mathdb
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + namelg + " FROM problems WHERE category='" + problems + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	fmt.Fprint(w, `<h2><img src="/static/images/`, problems, `.png">`, header, `</h2>`)
	fmt.Fprint(w, `<ul class="sectionul">`)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Fprint(w, `<li onclick = "request.open('POST', '/posthandler', true); request.setRequestHeader('Content-type', 'application/x-www-form-urlencoded'); request.setRequestHeader('Connection', 'close'); request.send('problemname=`, name, `');">`, name, `</li>`)
	}
	rows.Close()
	fmt.Fprint(w, `</ul>`)
}

func problemHandler1(r *http.Request, problems string) string {
	dbname, header := DbNameAndHeaderReturn(r, problems)
	lang, err := r.Cookie("lang")
	var namelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
	} else {
		namelg = `name`
	}

	var dbcons string
	if dbname == physdbname {
		dbcons = physdb
	} else {
		dbcons = mathdb
	}

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT " + namelg + " FROM problems WHERE category='" + problems + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}
	var result string
	result+=`<h2><img src="/static/images/`+problems+`.png">`+header+`</h2>`
	result+=`<ul class="sectionul">`
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		result+=`<li><a href="/problems/`+problems+`/`+name+`">`+name+`</li>`
	}
	rows.Close()
	result+= `</ul>`
	return result
}

func problemNameHandler(w http.ResponseWriter, r *http.Request, problemname string) {
	lang, err := r.Cookie("lang")
	var namelg, authorlg, problemlg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		authorlg = `author_en`
		problemlg = `problem_en`
	} else {
		namelg = `name`
		authorlg = `author`
		problemlg = `problem`
	}

	dbcons := physdb

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT category," + authorlg + "," + problemlg + " FROM problems WHERE " + namelg + " = '" + problemname + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var category, problem, author string
	for rows.Next() {
		err = rows.Scan(&category, &author, &problem)
		_, header := DbNameAndHeaderReturn(r, category)
		fmt.Fprint(w, `<div class="articleinfo"><div class="articlesection"><h3><img src="/static/images/`, category, `.png">`, header, `</h3></div><div class="articleauthor"><h3>`, author, `</h3></div></div>`)
		fmt.Fprint(w, "<h2>", problemname, "</h2>")
		fmt.Fprint(w, problem)
	}
	rows.Close()
}

func problemNameHandler1(r *http.Request, problemname string) string {
	lang, err := r.Cookie("lang")
	var namelg, authorlg, problemlg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		authorlg = `author_en`
		problemlg = `problem_en`
	} else {
		namelg = `name`
		authorlg = `author`
		problemlg = `problem`
	}

	dbcons := physdb

	db, err := sql.Open("postgres", dbcons)
	if err != nil {
		log.Println("database connection error", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT category," + authorlg + "," + problemlg + " FROM problems WHERE " + namelg + " = '" + problemname + "'")
	if err != nil {
		log.Println("scientists are not founded", err)
	}

	var category, problem, author, result string
	for rows.Next() {
		err = rows.Scan(&category, &author, &problem)
		_, header := DbNameAndHeaderReturn(r, category)
		result += `<div class="articleinfo"><div class="articlesection"><h3><img src="/static/images/`+ category+`.png">`+header+`</h3></div><div class="articleauthor"><h3>`+author+`</h3></div></div>`
		result += "<h2>"+problemname+"</h2>"
		result += problem
	}
	rows.Close()
	return result
}

func MathIndex(w http.ResponseWriter, r *http.Request) {
	lang, err := r.Cookie("lang")
	type data struct {
		Addscript, Article, Mathh, Geometry, Numbers, Algebraic, Algebra, Functions, Transcendental, Equations, Inequalities, Calculus, Plane, Solid, Coordinates, Transformations, Vectors string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Geometry = "Geometry"
		Data.Numbers = "Numbers"
		Data.Algebraic = "Algebraic expressions"
		Data.Algebra = "Algebra and analysis begin"
		Data.Functions = "Functions and graphs"
		Data.Transcendental = "Transcendental expressions"
		Data.Equations = "Equations and systems of equations"
		Data.Inequalities = "Inequalities"
		Data.Calculus = "Elements of calculus"
		Data.Plane = "Geometric figures on plane"
		Data.Solid = "Solid geometry"
		Data.Coordinates = "Cartesian system of coordinates"
		Data.Transformations = "Transformations of the figures"
		Data.Vectors = "Vectors"

		Data.Mathh = "Math"
	} else {
		Data.Geometry = "Геометрия"
		Data.Numbers = "Числа"
		Data.Algebraic = "Алгебраические выражения"
		Data.Algebra = "Алгебра и начала анализа"
		Data.Functions = "Функции и графики"
		Data.Transcendental = "Трансцендентые выражения"
		Data.Equations = "Уравнения и системы уравнений"
		Data.Inequalities = "Неравенства"
		Data.Calculus = "Элементы математического анализа"
		Data.Plane = "Геометрические фигуры на плоскости"
		Data.Solid = "Стереометрия"
		Data.Coordinates = "Декартовы координаты"
		Data.Transformations = "Преобразования фигур"
		Data.Vectors = "Векторы"
		Data.Addscript = `
			<script>
			function rePosH() {
				if(window.innerWidth > 769) {
					$("#mathindexblock h1").css("right", "520px");
					$("#problemsindexblock h1").css("right", "340px");
				} else {
					$("#mathindexblock h1").css("right", "260px");
					$("#problemsindexblock h1").css("right", "170px");	
				}
			}
			rePosH();
			</script>
		`
		Data.Mathh = "Математика"
	}

	t, err := template.ParseFiles("templates/math.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func MathInfo(w http.ResponseWriter, r *http.Request) {
	infoSlice := strings.Split(r.URL.Path[len("/math/"):], "/")
	var name, theme string
	var Article string
	if len(infoSlice) == 1 {
		theme = r.URL.Path[len("/math/"):]
		Article = themeHandler1(r, theme)
	} else {
		theme = infoSlice[0]
		name = infoSlice[1]
		Article = nameHandler1(r, theme, name)
	}

	lang, err := r.Cookie("lang")
	type data struct {
		Addscript, Article, Mathh, Geometry, Numbers, Algebraic, Algebra, Functions, Transcendental, Equations, Inequalities, Calculus, Plane, Solid, Coordinates, Transformations, Vectors string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Mathh = "Math"
		Data.Article = Article
	} else {
		Data.Addscript = `
			<script>
			function rePosH() {
				if(window.innerWidth > 769) {
					$("#mathindexblock h1").css("right", "520px");
					$("#problemsindexblock h1").css("right", "340px");
				} else {
					$("#mathindexblock h1").css("right", "260px");
					$("#problemsindexblock h1").css("right", "170px");	
				}
			}
			rePosH();
			</script>
		`
		Data.Mathh = "Математика"
		Data.Article = Article
	}

	t, err := template.ParseFiles("templates/math.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func PhysIndex(w http.ResponseWriter, r *http.Request) {
	lang, err := r.Cookie("lang")
	type data struct {
		Article, Physh, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Physh = "Physics"

		Data.Mechanics = "Mechanics"
		Data.Thermodynamics = "Bases of molecular physics and thermodynamics"
		Data.Electrodynamics = "Electrodynamics"
		Data.Oscillations = "Oscillations and waves"
		Data.Atoms = "Physics of atoms and molecules"
		Data.Optics = "Optics"
		Data.Hard = "Bases of solids physics"
		Data.Nucleus = "Atomic nucleus and elementary particles physics"
	} else {
		Data.Physh = "Физика"

		Data.Mechanics = "Механика"
		Data.Thermodynamics = "Основы молекулярной физики и термодинамики"
		Data.Electrodynamics = "Электродинамика"
		Data.Oscillations = "Колебания и волны"
		Data.Atoms = "Физика атомов и молекул"
		Data.Optics = "Оптика"
		Data.Hard = "Основы физики твердого тела"
		Data.Nucleus = "Физика ядра и элементарных частиц"
	}

	t, err := template.ParseFiles("templates/physics.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func PhysInfo(w http.ResponseWriter, r *http.Request) {
	infoSlice := strings.Split(r.URL.Path[len("/physics/"):], "/")
	var name, theme string
	var Article string
	if len(infoSlice) == 1 {
		theme = r.URL.Path[len("/physics/"):]
		Article = themeHandler1(r, theme)
	} else {
		theme = infoSlice[0]
		name = infoSlice[1]
		Article = nameHandler1(r, theme, name)
	}

	lang, err := r.Cookie("lang")
	type data struct {
		Article, Physh, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Physh = "Physics"
		Data.Article = Article
	} else {
		Data.Physh = "Физика"
		Data.Article = Article
	}

	t, err := template.ParseFiles("templates/physics.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func ProblemIndex(w http.ResponseWriter, r *http.Request) {
	lang, err := r.Cookie("lang")
	type data struct {
		Article, Problemsh, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Problemsh = "Problems"

		Data.Mechanics = "Mechanics"
		Data.Thermodynamics = "Bases of molecular physics and thermodynamics"
		Data.Electrodynamics = "Electrodynamics"/*
		Data.Oscillations = "Oscillations and waves"
		Data.Atoms = "Physics of atoms and molecules"
		Data.Optics = "Optics"
		Data.Hard = "Bases of solids physics"
		Data.Nucleus = "Atomic nucleus and elementary particles physics"*/
	} else {
		Data.Problemsh = "Задачи"

		Data.Mechanics = "Механика"
		Data.Thermodynamics = "Основы молекулярной физики и термодинамики"
		Data.Electrodynamics = "Электродинамика"/*
		Data.Oscillations = "Колебания и волны"
		Data.Atoms = "Физика атомов и молекул"
		Data.Optics = "Оптика"
		Data.Hard = "Основы физики твердого тела"
		Data.Nucleus = "Физика ядра и элементарных частиц"*/
	}

	t, err := template.ParseFiles("templates/problems.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}

func ProblemInfo(w http.ResponseWriter, r *http.Request) {
	infoSlice := strings.Split(r.URL.Path[len("/problems/"):], "/")
	var name, theme string
	var Article string
	if len(infoSlice) == 1 {
		theme = r.URL.Path[len("/problems/"):]
		Article = problemHandler1(r, theme)
	} else {
		name = infoSlice[1]
		Article = problemNameHandler1(r, name)
	}

	lang, err := r.Cookie("lang")
	type data struct {
		Article, Problemsh, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
	}

	var Data data
	if lang != nil && lang.Value == `en` {
		Data.Problemsh = "Problems"

		Data.Article = Article
	} else {
		Data.Problemsh = "Задачи"

		Data.Article = Article
	}

	t, err := template.ParseFiles("templates/problems.tpl")
	if err != nil {
		log.Println("template error")
	}
	err = t.Execute(w, Data)
	if err != nil {
		log.Println("template print error", err)
	}
}