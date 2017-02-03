// richi web server
package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

var (
    C []string
    R *rand.Rand
)

func Cache() {
    C = make([]string, 1)
    C[0] = "hi"
    C = append(C, "lo")
    C = append(C, "open")
    C = append(C, "close")
    // load from txt file
    // one entry per line
    // init rng
}

func Rng() {
    R = rand.New(rand.NewSource(time.Now().UnixNano()))
    fmt.Println(R.Intn(len(C)))
}

func Select() string {
    f0 := R.Intn(len(C))
    return C[f0]
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
    Rng()
    http.HandleFunc("/", RichiHandler)
    http.ListenAndServe(":8080", nil)
}
