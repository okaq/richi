// okaq web server
// 2017-05-22
package main

import (
    "bufio"
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "tola.html"
)

type Player struct {
    Id string
    Pixels []int
}

func NewPlayer() *Player {
    return &Player{}
}

type Cache struct {
    Players map[string]*Player
    sync.Mutex
}

func UbiaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    s0 := bufio.NewScanner(r.Body)
    s0.Scan()
    t0 := s0.Text()
    b0 := s0.Bytes()
    fmt.Printf("Length (in bytes) of the request body is: %d.\n", len(b0))
    fmt.Printf("Request body data: %s.\n", t0)
    w.Header().Set("Content-type","text/plain")
    w.Write([]byte("oko"))
}

func main() {
    fmt.Println("web server started on localhost:8080")
    t0 := time.Now()
    fmt.Printf("begin time: %s\n", t0.String())
    http.HandleFunc("/", UbiaHandler)
    http.HandleFunc("/a", PidHandler)
    http.ListenAndServe(":8080", nil)
}
// logging
// version number maps to directory
// open files for:
// metadata (version, start, stats)
// request logs

