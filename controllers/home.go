package controllers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
    test := "test"
    tmpl, _ := template.ParseFiles(filepath.Join("views", "layout.html"), filepath.Join("views", "home.html"))
    tmpl.Execute(w, test)
}
