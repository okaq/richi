// okaq eye splash
// AQ <aq@okaq.com>
// 2017-10-26
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "sync"
    "time"
)

const (
    DEBUG = true
    EYE = "zye.html"
)

var (
    C *Cache
)

type Cache struct {
    Data map[string]string
    *sync.Mutex
}

func NewCache() *Cache {
    c := Cache{}
    c.Data = make(map[string]string)
    return &c
}

func EyeHandler(w http.ResponseWriter, r *http.Request) {
    if DEBUG {
        fmt.Println(r)
    }
    http.ServeFile(w,r,EYE)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    // ready player id
    if DEBUG {
        fmt.Println(r)
    }
    w.Header().Set("Content-type", "text/plain")
    w.Write([]byte("ok pid one"))
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
    // cache data sent
    // update global state
    // write details to client
    if DEBUG {
        fmt.Println(r)
    }
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.Write([]byte("data unread"))
        return
    }
    s0 := string(b0)
    if DEBUG {
        fmt.Println(s0)
    }
    C.Data[s0] = "001"
    if DEBUG {
        fmt.Println(C)
    }
    j0 := []byte(`{"id":0,"state":"zero"}`)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j0)
}

func main() {
    if DEBUG {
        fmt.Println("okaq web eye live...")
        fmt.Printf("localhost:8080 start time: %s\n", time.Now().String())
    }
    C = NewCache()
    http.HandleFunc("/", EyeHandler)
    http.HandleFunc("/a", PidHandler)
    http.HandleFunc("/b", DataHandler)
    http.ListenAndServe(":8080", nil)
}
