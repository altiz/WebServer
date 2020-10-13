package controller

import (
	"html/template"
	"net/http"
)

// Sayhello начинается с заглавной для того, чтобы метод был доступен вне этого файла
func Sayhello(w http.ResponseWriter, r *http.Request) {
	data := "Go Template"
	tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
	tmpl.Execute(w, data)
}
