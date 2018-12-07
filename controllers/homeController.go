package controllers

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type DataPage struct {
	PageTitle string
	Todos     []Todo
}

func HomeIndexHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:       "Homepage",
		Description: "Halaman the shonet",
		Data: DataPage{
			PageTitle: "TO DO LIST",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		},
	}

	tmpl := template.Must(template.ParseFiles("views/layout/main.tpl", "views/element/title.tpl", "views/home/home.tpl"))
	tmpl.Execute(w, data)
}

func TitlePage() string {
	return "Title Homepage"
}
