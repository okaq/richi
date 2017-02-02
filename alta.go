// richi web server
package main

import (
    "fmt"
    "net/http"
)

var (
    C []string
)

func Cache() {
    C = make([]string, 1)
    C[0] = "hi"
    C = append(C, "lo")
    // load from txt file
    // one entry per line
    // init rng
}

func Select() string {
    return C[1]
}

func RichiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Println(Select())
    w.Write([]byte("ok richi!"))
    // select from data at random
}

func main() {
    fmt.Println("starting richi web server on localhost:8080")
    Cache()
    http.HandleFunc("/", RichiHandler)
    http.ListenAndServe(":8080", nil)
}
