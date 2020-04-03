package main

import (
    "github.com/acra5y/go-dilation"
    "github.com/acra5y/n-dilation-computer/handler"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/dilation", handler.DilationHandler(godilation.UnitaryNDilation))
    port := ":8080"
    log.Printf("Listening on port %s...", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
