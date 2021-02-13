package handler

import (
	"golang_web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to Home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "I'm learning golang website",
	// 	"content": "I'm learning golang website by agung setiawan",
	// }

	data := []entity.Product{
	{ID:1,Name:"mobilio",Price:220000000,Stock:3},
	{ID:2,Name:"xpander",Price:270000000,Stock:8},
	{ID:3,Name:"pajero sport",Price:500000000,Stock:1},
}


	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world,saya sedang belajar golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mario from nitendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product page : %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter,r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w,"Error is happening",http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
}

	return

	}

	http.Error(w, "Eror is happening", http.StatusBadRequest)

}

func Process(w http.ResponseWriter,r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
}

	name := r.Form.Get("name")
	message := r.Form.Get("message")

	data := map[string]interface{}{
		"name": name,
		"message": message,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
	return
	}

	
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror is happening", http.StatusInternalServerError)
		return
	}

return
	}
http.Error(w, "Eror is happening", http.StatusBadRequest)
}