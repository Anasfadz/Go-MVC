package main

import (
    "net/http"
    "go-mvc/controllers"
)

func main() {
    http.HandleFunc("/", controllers.Home)
    http.ListenAndServe(":8080", nil)
}