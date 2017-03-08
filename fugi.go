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

func KeysHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func Load() {
    // read files from directory
    // populate list object / json
}

func main() {
    fmt.Println("start goma bitmap sample on localhost:8080")
    Load()
    http.HandleFunc("/", FugiHandler)
    http.HandleFunc("/a", KeysHandler)
    http.ListenAndServe(":8080", nil)
}
