// okaq web server
// AQ <aq@okaq.com>
// 2017-04-20
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "nopi.html"
)

func OtarHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("server started on localhost:8080")
    T = time.Now()
    fmt.Printf("time stamp: %s\n", T.String())
    http.HandleFunc("/", OtarHandler)
    http.ListenAndServe(":8080", nil)
}
