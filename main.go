package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, "welcome to my web1")
        })
        fs := http.FileServer(http.Dir("static/"))
        http.Handle("/static/", http.StripPrefix("/static/", fs))
        http.ListenAndServe(":80", nil)
}
