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
    INDEX = "yayo.html"
)

func AayoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("okaq geo web start\n%s\n", time.Now().String())
    http.HandleFunc("/", AayoHandler)
    http.ListenAndServe(":8080", nil)
}
