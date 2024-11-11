package main

import (
    "net/http"
    "html/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, "World")
}
func main() {
    http.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", nil)
}