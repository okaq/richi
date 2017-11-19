// okaq web server
// csp nano game
// AQ <aq@okaq.com>
// 2017-11-18

package main

import (
    "fmt"
    "time"
    "net/http"
    "sync"
)

const (
    INDEX = "xevi.html"
)

var (
    // cache, atomic index, pid list
    C int
    P []string
)

type Pids struct {
    // counter
    Count uint64
    Id []string
    sync.RWMutex
}

func NewPids *Pids {
    p := new(Pids)
    return &Pids
}

func Cache() {
    C = 0
    P = make([]int, 100)
}

func AeviHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // read json from request data
}

func main() {
    fmt.Printf("csp web serve start %s\n", time.Now().String())
    Cache()
    http.HandleFunc("/", AeviHandler)
    http.HandleFunc("/b", PidHandler)
    http.ListenAndServe(":8080", nil)
}
