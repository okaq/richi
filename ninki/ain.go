// okaq nano game
// "surprisingly popular"
// AQ <aq@okaq.com>
// 2017-08-04
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    ZIN = "zin.html"
)

func ZinHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,ZIN)
}

func AbHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // good habit
    // stats page and cache
    // hold request and data transfered count
}

func main() {
    fmt.Printf("localhost:8080 zin web start: %s\n", time.Now().String())
    http.HandleFunc("/", ZinHandler)
    http.HandleFunc("/a", AbHandler)
    http.ListenAndServe(":8080", nil)
}
