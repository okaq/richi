// okaq eye splash
// AQ <aq@okaq.com>
// 2017-10-26
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    DEBUG = true
    EYE = "zye.html"
)

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

func main() {
    if DEBUG {
        fmt.Println("okaq web eye live...")
        fmt.Printf("localhost:8080 start time: %s\n", time.Now().String())
    }
    http.HandleFunc("/", EyeHandler)
    http.HandleFunc("/a", PidHandler)
    http.ListenAndServe(":8080", nil)
}
