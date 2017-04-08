// okaq web server
// 2017-04-08
package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

const (
    INDEX = "lumi.html"
)

var (
    M *Marble
    R *rand.Rand
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
    // populate list
    m.Pop()
    return &m
}

func (m *Marble) Pop() {
    for i := 0; i < m.Max; i++ {
        // get date time stamp
        i0 := time.Now().UnixNano()
        // get unique random id
        i1 := R.Int()
        s0 := fmt.Sprintf("%d+%d", i0, i1)
        m.List[i] = s0
    }
}

func (m *Marble) Inc() string {
    if m.Count >= m.Max {
        return "limit"
    } else {
        s0 := m.List[m.Count]
        m.Count = m.Count + 1
        return s0
    }
}

func MeniHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    M.Lock()
    defer M.Unlock()
    s0 := M.Inc()
    w.Write([]byte(s0))
}

func main() {
    fmt.Println("serving on localhost:8080")
    R = rand.NewRandom(rand.NewSeed(time.Now().UnixNano()))
    M = NewMarble()
    fmt.Println(M)
    http.HandleFunc("/", MeniHandler)
    http.HandleFunc("/a", FlushHandler)
    http.ListenAndServe(":8080", nil)
}

// pre-allocate unique ids
// create list and counter
// to assign to clients
// incrementing per request
// when end is reached reset

// create map with rw lock
// to store game state data

