// okaq web server
// threejs / webgl
// user experience
// 2017-06-04
// AQ <aq@okaq.com>
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    // INDEX = "boli.html"
    INDEX = "coli.html"
)

func BoliHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // dump save png frame from webgl

    // using toDataURL
    // get data in form data:image/png;base64,iVBORw0KGgo...

    // send to server as json string
    // strip prefix upto comma to get png data
    // b64data := input[strings.IndexByte(input, ',')+1:]
    // then decode to bytes
    // data, err := base64.StdEncoding.DecodeString(b64data)
    // data can then be written to file as png image
}

func StitchHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // button event to begin server side video generation
}

func main() {
    fmt.Printf("started web server on localhost:8080\nat time: %s\n", time.Now().String())
    http.HandleFunc("/", BoliHandler)
    http.HandleFunc("/s", SaveHandler)
    http.HandleFunc("/a", StitchHandler)
    http.ListenAndServe(":8080", nil)
}

