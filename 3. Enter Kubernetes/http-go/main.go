package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "log"
    "net/http"
    "os"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Fprint(w, "Welcome! "+hostname+"\n")
    } else {
        fmt.Fprint(w, "Welcome! ERROR\n")
    }
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)

    log.Fatal(http.ListenAndServe(":8080", router))
}
