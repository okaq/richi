// okaq web server
// AQ <aq@okaq.com>
// 2017-03-30
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

var (
    INDEX = "kona.html"
)

type Uniq struct {
    Date string
    Id string
}

func HiwaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func ScooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write([]byte("scoobah party!"))
}

func IdHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    var u Uniq
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(u)
}

func main() {
    fmt.Println("starting server on localhost:8080")
    http.HandleFunc("/", HiwaHandler)
    http.HandleFunc("/a", ScooHandler)
    http.HandleFunc("/b", IdHandler)
    http.ListenAndServe(":8080", nil)
}

// test peer brokerage
// xhr game state conn
