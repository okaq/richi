// okaq web server
// csp nano game
// AQ <aq@okaq.com>
// 2017-11-06

package main

import (
    "fmt"
    "time"
    "net/http"
)

const (
    // INDEX = "zevi.html"
    // INDEX = "yevi.html"
    INDEX = "xevi.html"
)

func AeviHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("csp web serve start %s\n", time.Now().String())
    http.HandleFunc("/", AeviHandler)
    http.ListenAndServe(":8080", nil)
}
