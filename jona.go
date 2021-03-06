// okaq web server
// AQ <aq@okaq.com>
// 2017-03-30
package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

var (
    INDEX = "kona.html"
    B *Bonne
    L int
    M []byte
)

type Uniq struct {
    Date string
    Id string
}

func NewLocation() int {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    return rng.Intn(480)
}

type Bonne struct {
    Top map[string]string
    // map to game state object
    // key format: "Id+Date"
    sync.RWMutex
}

func NewBonne() *Bonne {
    b := Bonne{}
    b.Top = make(map[string]string)
    return &b
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
    s0 := fmt.Sprintf("%s+%s", u.Id, u.Date)
    fmt.Println(s0)
    // grab lock and write to cache
    // ensure unlocked before flush
    B.Lock()
    defer B.Unlock()
    B.Top[s0] = "live"
    fmt.Println(B.Top)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write([]byte("got uniq"))
}

func LocHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // xhr response type "text"
     w.Header().Set("Access-Control-Allow-Origin", "*")
     w.Write(M)
}

func main() {
    fmt.Println("starting server on localhost:8080")
    B = NewBonne()
    L = NewLocation()
    M = []byte(fmt.Sprintf("%d", L))
    http.HandleFunc("/", HiwaHandler)
    http.HandleFunc("/a", ScooHandler)
    http.HandleFunc("/b", IdHandler)
    http.HandleFunc("/c", LocHandler)
    http.ListenAndServe(":8080", nil)
}

// test peer brokerage
// xhr game state conn

// 480 byte slice describes state of grid
// each update from clients flips individual bits
// efficient optimization is to just send deltas

