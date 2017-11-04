// okaq web server
// IEX IPO SIM
// https://iextrading.com/
// AQ <aq@okaq.com>
// 2017-11-04

package main

import (
    "fmt"
    "time"
)

const (
    INDEX = "zopi.html"
)

func AopiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("stock application started at: %s\n", time.Now().String())
    http.HandleFunc("/", AopiHandler)
    http.ListenAndServe(":8080", nil)
}
