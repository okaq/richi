// okaq nano game
// "surprisingly popular"
// AQ <aq@okaq.com>
// 2017-08-04
package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

const (
    ZIN = "zin.html"
    INDEX = "xin.html"
    BITMAP = "win.html"
)

var (
    C *Cache
)

type Cache struct {
    sync.Mutex
    Count int
}

func (c *Cache) Increment() {
    c.Lock()
    defer c.Unlock()
    c.Count = c.Count + 1
}

func (c *Cache) Json() string {
    c.Lock()
    defer c.Unlock()
    s0 := fmt.Sprintf("{\"count\":%d}", c.Count)
    return s0
}
func NewCache() *Cache {
    c := Cache{}
    c.Count = 0
    return &c
}

func ZinHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,ZIN)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func AbHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("greetings from ain!"))
    C.Increment()
    s0 := fmt.Sprintf("{\"count\":%d}", C.Count)
    b0 := []byte(s0)
    w.Write(b0)
    // good habit
    // stats page and cache
    // hold request and data transfered count
}

func BitHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,BITMAP)
}

func main() {
    fmt.Printf("localhost:8080 zin web start: %s\n", time.Now().String())
    C = NewCache()
    // http.HandleFunc("/", ZinHandler)
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/a", AbHandler)
    http.HandleFunc("/b", BitHandler)
    http.ListenAndServe(":8080", nil)
}
