// okaq web server
// geo math
// AQ <aq@okaq.com>
// 2017-12-20

package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    // INDEX = "zayo.html"
    // INDEX = "yayo.html"
    // INDEX = "xayo.html"
    INDEX = "wayo.html"
    BIN = "node"
)

func AayoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func Bin() {
    // static assets from build
    pre := fmt.Sprintf("/%s/", BIN)
    fs := http.FileServer(http.Dir(BIN))
    http.Handle(pre, http.StripPrefix(pre,fs))
}

func main() {
    fmt.Printf("okaq geo web start\n%s\n", time.Now().String())
    http.HandleFunc("/", AayoHandler)
    Bin()
    http.ListenAndServe(":8080", nil)
}
