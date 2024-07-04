package main

import (
    "net/http"
    "github.com/Anas/Go-MVC/controllers"
)

func main() {
    http.HandleFunc("/", controllers.Home)
    http.ListenAndServe(":8080", nil)
}