package gui

import (
	"net/http"
)

func InitGUIMux() *http.ServeMux {
	GUIMux := http.NewServeMux()
	GUIMux.HandleFunc("/", handler)
	return GUIMux
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	User  string
	Todos []Todo
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		User: "User",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	testTemplate.Execute(w, data)
}
