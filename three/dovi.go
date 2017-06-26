// golang static dir web server
package main

import (
    "fmt"
    "net/http"
)

const (
    DIR = "dovi"
)

func main() {
    fmt.Println("static dovi file serving...")
    fs := http.FileServer(http.Dir(DIR))
    http.Handle("/", fs)
    http.ListenAndServe(":8080", nil)
}
