// okaq web server
// threejs / webgl
// user experience
// 2017-06-04
// AQ <aq@okaq.com>
package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

const (
    // INDEX = "boli.html"
    // INDEX = "coli.html"
    // INDEX = "doli.html"
    INDEX = "eoli.html"
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
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    s := string(b)
    fmt.Println(b,s)
    a := strings.IndexByte(s, ',')
    fmt.Println(a)
    // if its a set byte offset, index without string cast
    d := s[a+1:]
    f, err := base64.StdEncoding.DecodeString(d)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
    // save bytes to png format file
    ioutil.WriteFile("vrox_01.png",f,0444)
}

func StitchHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // button event to begin server side video generation
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // let's see if we can view html5 video webm
}

func ChronoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // server clock to sync client time
    // just in time unix epoch 1.5M
}

func main() {
    fmt.Printf("started web server on localhost:8080\nat time: %s\n", time.Now().String())
    http.HandleFunc("/", BoliHandler)
    http.HandleFunc("/a", SaveHandler)
    http.HandleFunc("/b", StitchHandler)
    http.HandleFunc("/c", ViewHandler)
    http.HandleFunc("/d", ChronoHandler)
    http.ListenAndServe(":8080", nil)
}

