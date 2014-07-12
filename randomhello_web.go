package main

import (
    "net/http"
    "io"
    "github.com/n10o/samgol"
)

func Hello(w http.ResponseWriter, req *http.Request) {
    var word string = "<h1>Hello " + samgol.GetPlace() + "</h1>"
    io.WriteString(w, word)
}

func main() {
    http.HandleFunc("/", Hello)
    http.ListenAndServe(":4000", nil)
}
