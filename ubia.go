// okaq web server
// 2017-05-22
package main

import (
    "bufio"
    "fmt"
    "net/http"
    "sync"
    "time"
)

const (
    INDEX = "tola.html"
)

var (
    C *Cache
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

func NewCache() *Cache {
    c := Cache{}
    c.Players = make(map[string]*Player)
    return &c
}

func (c *Cache) Print() {
    for k,v := range c.Players {
        fmt.Println(k)
        fmt.Printf("pixels: %d\n", v.Pixels)
    }
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
    // create chan receiver to lock write to cache
    // Fprintf format string for output
}

func Pid2Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    s0 := bufio.NewScanner(r.Body)
    defer r.Body.Close()
    s0.Scan()
    // send text
    // lock here
    p0 := NewPlayer()
    t0 := s0.Text()
    b0 := s0.Bytes()
    C.Lock()
    defer C.Unlock()
    C.Players[t0] = p0
    fmt.Printf("Length (in bytes) of the request body is: %d.\n", len(b0))
    fmt.Printf("Request body data: %s.\n", t0)
    C.Print()
    w.Header().Set("Content-type","text/plain")
    fmt.Fprintf(w, "%s+%d", s0.Text(), time.Now().UnixNano())
}

func PixHandler(w http.ResponseWriter, r *http.Request) {
    // process pixel data from typed array
    // update cache data
}

func HexHandler(w http.ResponseWriter, r *http.Request) {
    // hex response test
    s0 := `have a glorious du jour!`
    b0 := []byte(s0)
    e0 := fmt.Sprintf("%x", b0)
    b1 := []byte(e0)
    fmt.Println(s0,b0,e0,b1)
}

func B64Handler(w http.ResponseWriter, r *http.Request) {
    // base64 response test
    s0 := `bonne frisson du plaisir`
    b0 := []byte(s0)
}

func main() {
    fmt.Println("web server started on localhost:8080")
    t0 := time.Now()
    fmt.Printf("begin time: %s\n", t0.String())
    C = NewCache()
    http.HandleFunc("/", UbiaHandler)
    http.HandleFunc("/a", Pid2Handler)
    http.HandleFunc("/b", PixHandler)
    http.HandleFunc("/c", HexHandler)
    http.HandleFunc("/d", B64Handler)
    http.ListenAndServe(":8080", nil)
}
// logging
// version number maps to directory
// open files for:
// metadata (version, start, stats)
// request logs

