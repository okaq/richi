// richi web server
package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

const (
    // word file, one entry per line
    BUN = "bun.txt"
    INDEX = "ciro.html"
)

var (
    C []string
    R *rand.Rand
)

func Cache() {
    C = make([]string, 1)
    C[0] = "cardinale"
    // C = append(C, "lo")
    // C = append(C, "open")
    // C = append(C, "close")
    // load from txt file
    // one entry per line
    // init rng
}

func Load() {
    // load from input bun text file
    f0, err := os.Open(BUN)
    if err != nil {
        fmt.Println(err)
    }
    defer f0.Close()
    // scan line by line
    // i := 0
    s0 := bufio.NewScanner(f0)
    for s0.Scan() {
        // i = i + 1
        // fmt.Printf("%d. meme: %s\n", i, s0.Text())
        C = append(C, s0.Text())
    }
    // report final size
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
    // fmt.Println(Select())
    // w.Write([]byte("ok richi!"))
    // select from data at random
    http.ServeFile(w, r, INDEX)
}

func MemeHandler(w http.ResponseWriter, r *http.Request) {
    // select and return as json
    // intended for xhr reqs
    fmt.Println(r)
    c0 := Select()
    // plain text
    w.Write([]byte(c0))
}

func main() {
    fmt.Println("starting richi web server on localhost:8080")
    Cache()
    Load()
    Rng()
    http.HandleFunc("/", RichiHandler)
    http.HandleFunc("/a", MemeHandler)
    http.ListenAndServe(":8080", nil)
}
