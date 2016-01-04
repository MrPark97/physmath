package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	physdb     = "postgres://usqhjiqrwslumu:Z_PorYcYUUEdwgNBUFhWk3vMVZ@ec2-54-204-8-138.compute-1.amazonaws.com:5432/datuej4qjvf8bm"
	mathdb     = "postgres://usqhjiqrwslumu:Z_PorYcYUUEdwgNBUFhWk3vMVZ@ec2-54-204-8-138.compute-1.amazonaws.com:5432/datuej4qjvf8bm"
	physdbname = "u233172724_phys"
	mathdbname = "u233172724_math"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/posthandler", PostHandler)
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
		Data.Intended = `This site is intended for all, who wants to hear little more about physics and mathematics. The site is created in the first place for you, users so you do not see the boring advertisment and madly moving flash-banners or gifs. The design and programme code of this site were created without different CMS and copying programme code from elbowroom of the worldwide network by student of the Tashkent Professional College of Information Technologies. Hope that you like the site and you value it on value. Want you to find useful information in the site depths.`
		Data.Myname = `Park Eugene`
		Data.Myp = `Park Eugene (16.05.1997, Tashkent - ...) - programmer.`
		Data.Myp1 = `Mr.Park was born 16 May 1997. In 2004-2013 learned in specialized school 103.
			In 2013 has entered in Tashkent Professional College of Information Technologies, profession - a technician of software and multimedia systems.`
		Data.Langh = `Language`
		Data.ScientistHeader = `Scientists`
	} else {
		Data.Intended = `Этот сайт предназначен для всех, кто хочет узнать немного больше о физике и математике. Сайт создан в первую очередь для вас, пользователей, поэтому вы не увидите надоедливой рекламы и безумно дрыгающихся флэш-баннеров или гифок. Дизайн и программный код данного сайта были созданы с нуля, без использования различных движков и копирования программного кода с просторов всемирной сети, студентом Ташкентского Профессионального Колледжа Информационных Технологий. Надеюсь, что вам понравится данный сайт и вы оцените его по достоинству. Желаю вам найти полезную информацию в недрах веб-страничек данного ресурса.`
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
		Mathh, Addscript, Mathp, Readmore, Physh, Physp, Problemsh, Problemsp, Mathh1, Numbers, Algebraic, Geometry, Algebra, Functions, Transcendental, Equations, Inequalities, Calculus, Mechanics, Thermodynamics, Electrodynamics, Oscillations, Atoms, Optics, Hard, Nucleus string
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

		Data.Mathh1 = "Algebra and analysis begin"
		Data.Numbers = "Numbers"
		Data.Algebraic = "Algebraic expressions"
		Data.Geometry = "Geometry"
		Data.Algebra = "Algebra and analysis begin"
		Data.Functions = "Functions and graphs"
		Data.Transcendental = "Transcendental expressions"
		Data.Equations = "Equations and systems of equations"
		Data.Inequalities = "Inequalities"
		Data.Calculus = "Elements of calculus"

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

		Data.Mathh1 = "Алгебра и начала анализа"
		Data.Numbers = "Числа"
		Data.Algebraic = "Алгебраические выражения"
		Data.Algebra = "Алгебра и начала анализа"
		Data.Functions = "Функции и графики"
		Data.Transcendental = "Трансцендентые выражения"
		Data.Equations = "Уравнения и системы уравнений"
		Data.Inequalities = "Неравенства"
		Data.Calculus = "Элементы математического анализа"

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
	if scientist := r.PostFormValue("scientist"); scientist != "" {
		scientistHandler(w, r, scientist)
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

func scientistHandler(w http.ResponseWriter, r *http.Request, scientist string) {
	lang, err := r.Cookie("lang")
	var namelg, scientistlg, articlelg, quotelg string
	if lang != nil && lang.Value == `en` {
		namelg = `name_en`
		scientistlg = `scientist`
		articlelg = `article_en`
		quotelg = `quote`
	} else {
		namelg = `name`
		scientistlg = `scientist_ru`
		articlelg = `article`
		quotelg = `quote_ru`
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
		fmt.Fprint(w, `<h2>`, name, `</h2>`)
		fmt.Fprint(w, `<blockquote>"`, quote, `"</blockquote>`)
		fmt.Fprint(w, `<img src="/static/scientists/`, img, `">`)
		fmt.Fprint(w, article)
	}
	rows.Close()
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
