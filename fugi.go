// okaq bitmap sample
// AQ <aq@okaq.com>
// 2017-03-08
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    INDEX = "goma.html"
)

var (
    P map[string]string
    J []byte
)

func FugiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func KeysHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Header().Set("Content-Type", "application/json")
    w.Write(J)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // json from req body
    // png name, sample bit array data

    // simple cache on server
    // map image name, data
    // write to file on initial post
}

func Load() {
    // read files from directory
    // populate list object / json
    f0, err := ioutil.ReadDir("png")
    if err != nil {
        fmt.Println(err)
    }
    for i := 0; i < len(f0); i++ {
        fmt.Println(f0[i])
        n0 := f0[i].Name()
        url0 := fmt.Sprintf("/p/%s", n0)
        P[n0] = url0
    }
    fmt.Println(P)
}

func Encode() {
    var err error
    J, err = json.Marshal(P)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(J)
    fmt.Println(len(J))
    fmt.Println(string(J))
}

func main() {
    fmt.Println("start goma bitmap sample on localhost:8080")
    P = make(map[string]string)
    Load()
    Encode()
    http.HandleFunc("/", FugiHandler)
    http.HandleFunc("/a", KeysHandler)
    PngServer := http.FileServer(http.Dir("png"))
    http.Handle("/p/", http.StripPrefix("/p/", PngServer))
    http.HandleFunc("/b", SaveHandler)
    http.ListenAndServe(":8080", nil)
}
