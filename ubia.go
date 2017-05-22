// okaq web server
// 2017-05-22
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "tola.html"
)

func UbiaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("web server started on localhost:8080")
    t0 := time.Now()
    fmt.Printf("begin time: %s\n", t0.String())
    http.HandleFunc("/", UbiaHandler)
    http.ListenAndServe(":8080", nil)
}
// logging
// version number maps to directory
// open files for:
// metadata (version, start, stats)
// request logs

