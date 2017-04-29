// okaq web server
// 2017-04-30
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "rola.html"
)

func SiboHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("web server started localhost:8080")
    t0 := time.Now()
    fmt.Printf("start time: %s\n", t0.String())
    http.HandleFunc("/", SiboHandler)
    http.ListenAndServe(":8080", nil)
}
