// okaq eye splash
// AQ <aq@okaq.com>
// 2017-10-26
package main

import (
    "fmt"
//     "net/http"
    "time"
)

const (
    DEBUG = true
    EYE = "zye.html"
)

func main() {
    fmt.Println("okaq web eye live...")
    fmt.Printf("localhost:8080 start time: %s\n", time.Now().String())
}
