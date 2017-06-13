// okaq web server
// threejs / webgl
// user experience
// 2017-06-04
// AQ <aq@okaq.com>
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    // INDEX = "boli.html"
    INDEX = "coli.html"
)

func BoliHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // dump save png frame from webgl
}

func main() {
    fmt.Printf("started web server on localhost:8080\nat time: %s\n", time.Now().String())
    http.HandleFunc("/", BoliHandler)
    http.HandleFunc("/s", SaveHandler)
    http.ListenAndServe(":8080", nil)
}

