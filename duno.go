// okaq swift web server
// AQ <aq@okaq.com>
// 2017-02-28
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "eliu.html"
)

func DonuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting on localhost:8080")
    http.HandleFunc("/", DonuHandler)
    http.ListenAndServe(":8080", nil)
}
