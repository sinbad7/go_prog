package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("home/grigorovich/go/src/github.com/go_prog/staticweb/public"))
    mux.Handle("/", fs)
    http.ListenAndServe(":8080", mux)
}