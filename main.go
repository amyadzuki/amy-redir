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
    http.HandleFunc ("/",           Hello)
    http.HandleFunc ("/ddgi/",      DDGTestImage)
    http.HandleFunc ("/google/",    Google)
    http.HandleFunc ("/pkmn/",      PokemonBig)
    http.HandleFunc ("/pkmn/smol/", PokemonSmol)
    fmt.Println ("Listening...")
    err := http.ListenAndServe (":" + port, nil)
    if err != nil {
        panic (err)
    }
}

func Hello (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln (w, "Hello, world")
}

func DDGTestImage (w http.ResponseWriter, r *http.Request) {
    http.Redirect (w, r, "https://duckduckgo.com/assets/logo_icon128.v101.png", http.StatusMovedPermanently)
}

func Google (w http.ResponseWriter, r *http.Request) {
    http.Redirect (w, r, "https://google.com", http.StatusMovedPermanently)
}

func Obfuscation (w http.ResponseWriter, r *http.Request) {
    http.Redirect (w, r, "https://google.com", http.StatusMovedPermanently)
}

func PokemonBig (w http.ResponseWriter, r *http.Request) {
    num := r.URL.Path [len ("/pkmn/"):]
    idx := strconv.ParseInt (num, 0, 64)
    if idx < 0 || idx > 251 {
        fmt.Fprintf (w, "404")
        return
    }
    url := Pkmn [idx]
    http.Redirect (w, r, url, http.StatusMovedPermanently)
}

func PokemonSmol (w http.ResponseWriter, r *http.Request) {
    num := r.URL.Path [len ("/pkmn/smol/"):]
    idx := strconv.ParseInt (num, 0, 64)
    if idx < 0 || idx > 251 {
        fmt.Fprintf (w, "404")
        return
    }
    url := PkmnSmol [idx]
    http.Redirect (w, r, url, http.StatusMovedPermanently)
}

const Pkmn []string = {
    "https://duckduckgo.com/assets/logo_icon128.v101.png" // test image
}

const PkmnSmol []string = {
    "https://duckduckgo.com/assets/logo_icon128.v101.png" // test image
}
