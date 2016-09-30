package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request)  {
	//For debugging sake, parse template in request.
	templ, err := template.ParseFiles("templates/layout.gohtml", "templates/index.gohtml")
	//Execute index template
	//err := templates.ExecuteTemplate(res, "index.gohtml", nil)
	err = templ.Execute(res, nil)
	if err != nil {
		log.Fatal(err)
	}
}

var templates *template.Template
func init(){
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main()  {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}