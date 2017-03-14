// okaq bitmap sample
// AQ <aq@okaq.com>
// 2017-03-08
package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
)

const (
    INDEX = "goma.html"
)

var (
    P map[string]string
    J []byte
    // I map[string][]int
    // I map[string]string
    I map[string]Img
)

type Img struct {
    Name string
    // Data []int
    // Data string
    Data map[string]byte
}

func (img *Img)Save() {
    fmt.Println("saving...")
    var b0 []byte
    for i := 0; i < 128; i++ {
        s0 := strconv.Itoa(i)
        b0 = append(b0, img.Data[s0])
    }
    fmt.Println(b0)
    enc := base64.NewEncoder(base64.RawStdEncoding, os.Stdout)
    enc.Write(b0)
    enc.Close()
    // neither raw, std, or url look quite right
    // fall back on dual request save
    // keep global var of file name key on first trip
    // then obtain array buffer data on second
}

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
    dec := json.NewDecoder(r.Body)
    // terminate the socket
    defer r.Body.Close()
    var i0 Img
    err := dec.Decode(&i0)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(i0)
    // check key exists
    _, ok := I[i0.Name]
    if ok {
        w.Write([]byte("already saved!"))
        return
    }
    I[i0.Name] = i0
    // I[i0.Name] = i0.Data
    // I = append(I, i0)
    fmt.Println(I)
    i0.Save()
    s0 := fmt.Sprintf("bytes read: %d\n", len(i0.Data))
    b0 := []byte(s0)
    w.Write(b0)
    // w.Write([]byte("ok saved!"))

    // simple cache on server
    // map image name, data
    // write to file on initial post

    // raw binary send from browser
    // arraybuffer encodes as []byte
    // and golang recognizes as base64 string
    // but if array buffer is embedded in json object
    // stringify renders as {0:255,1:255...}
    // map[string]byte of fixed size
    // can either reconstruct base64 from server
    // or send name and dat in two parts
    // first text string, second binary data
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
    // I = make(map[string][]int)
    // I = make(map[string]string)
    I = make(map[string]Img)
    Load()
    Encode()
    http.HandleFunc("/", FugiHandler)
    http.HandleFunc("/a", KeysHandler)
    PngServer := http.FileServer(http.Dir("png"))
    http.Handle("/p/", http.StripPrefix("/p/", PngServer))
    http.HandleFunc("/b", SaveHandler)
    http.ListenAndServe(":8080", nil)
}
