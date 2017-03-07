// okaq bitmap sample
// AQ <aq@okaq.com>
// 2017-03-08
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "goma.html"
)

func FugiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("start goma bitmap sample on localhost:8080")
    http.HandleFunc("/", FugiHandler)
    http.ListenAndServe(":8080", nil)
}
