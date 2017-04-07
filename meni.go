// okaq web server
// 2017-04-08
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "lumi.html"
)

func MeniHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("serving on localhost:8080")
    http.HandleFunc("/", MeniHandler)
    http.ListenAndServe(":8080", nil)
}

// pre-allocate unique ids
// create list and counter
// to assign to clients
// incrementing per request
// when end is reached reset

// create map with rw lock
// to store game state data

