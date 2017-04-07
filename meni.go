// okaq web server
// 2017-04-08
package main

import (
    "fmt"
    "net/http"
    "sync"
)

const (
    INDEX = "lumi.html"
)

var (
    M *Marble
)

type Marble struct {
    Count int
    Max int
    List []string
    sync.RWMutex
}

func NewMarble() *Marble {
    m := Marble{}
    m.Count = 0
    m.Max = 64
    m.List = make([]string, m.Max)
    return &m
}

func MeniHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("serving on localhost:8080")
    M = NewMarble()
    fmt.Println(M)
    http.HandleFunc("/", MeniHandler)
    http.ListenAndServe(":8080", nil)
}

// pre-allocate unique ids
// create list and counter
// to assign to clients
// incrementing per request
// when end is reached reset

// create map with rw lock
// to store game state data

