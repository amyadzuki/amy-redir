package main

import (
    "flag"
    "fmt"
    "net/http"
    "os"
    "strconv"
    )


func main () {
    var port string
    p := *flag.Int ("port", 0, "listening port")
    if p == 0 {
        port = os.Getenv ("PORT")
        if len (port) < 1 {
            port = "8080"
        }
    } else {
        port = strconv.FormatInt (int64 (p), 10)
    }
    http.HandleFunc ("/", hello)
    fmt.Println ("Listening...")
    err := http.ListenAndServe (":" + port, nil)
    if err != nil {
        panic (err)
    }
}

func hello (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln (w, "Hello, world")
}
