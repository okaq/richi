// okaq web server
// AQ <aq@okaq.com>
// 2017-04-20
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    INDEX = "nopi.html"
    // palette save file
    PAL = "png/pal.json"
)

var (
    T time.Time
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // data in json array form
    // list of [r,g,b,a] values
    // ordered list of equitable pairs
    // upon click save to disk file
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b0)
    s0 := string(b0)
    fmt.Println(s0)
    s1 := fmt.Sprintf("bytes written %d", len(b0))
    b1 := []byte(s1)
    ioutil.WriteFile(PAL, b0, 0777)
    w.Write(b1)
}

func OtarHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("server started on localhost:8080")
    T = time.Now()
    fmt.Printf("time stamp: %s\n", T.String())
    http.HandleFunc("/", OtarHandler)
    http.HandleFunc("/a", SaveHandler)
    http.ListenAndServe(":8080", nil)
}
