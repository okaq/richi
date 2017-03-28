// okaq web server
// AQ <aq@okaq.com>
// 2017-03-30
package main

import (
    "fmt"
    "net/http"
)

var (
    INDEX = "kona.html"
)


func HiwaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func ScooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("scoobah party!"))
}

func main() {
    fmt.Println("starting server on localhost:8080")
    http.HandleFunc("/", HiwaHandler)
    http.HandleFunc("/a", ScooHandler)
    http.ListenAndServe(":8080", nil)
}

// test peer brokerage
// xhr game state conn
