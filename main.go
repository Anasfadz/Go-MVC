package main

import (
    "net/http"
    "go-mvc/controllers"
)

func main() {
    http.HandleFunc("/", controllers.Home)
    http.HandleFunc("/test", controllers.Test)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
    http.ListenAndServe(":8080", nil)
}