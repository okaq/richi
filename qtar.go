// okaq web server
// 2017-04-24
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "popi.html"
)

func QtarHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("web serving on localhost:8080")
    t0 := time.Now()
    fmt.Printf("start time: %s\n", t0.String())
    http.HandleFunc("/", QtarHandler)
    http.ListenAndServe(":8080", nil)
}
