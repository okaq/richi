// okaq web server
// AQ <aq@okaq.com>
// 2017-03-18
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "iona.html"
)

func HiwaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok from server!"))
    http.ServeFile(w, r, INDEX)
}

func KitesHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("kite is flying in the sky"))
}

func main() {
    fmt.Println("starting server on localhost:8080")
    http.HandleFunc("/", HiwaHandler)
    http.HandleFunc("/a", KitesHandler)
    http.ListenAndServe(":8080", nil)
}

// bitmap decode and render 
// data in base64 json string binary
// expand to array of bits
// render as monochrome 32x32 squares
// begin library for bit manipulation
// rotate, shuffle, unshuffle

