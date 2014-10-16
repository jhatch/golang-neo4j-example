package main

import (
    "fmt"
    "net/http"
    "log"
)

const PORT = "8001"

func main() {
    // mount the api
    http.HandleFunc("/", renderWelcome)

    // start 'er up
    log.Printf("STARTING golang rest api on port %s\n", PORT)
    err := http.ListenAndServe(":"+PORT, nil)
    if err != nil {
        log.Fatalf("ERROR STARTING golang rest api on port %s\n", PORT, err)
        os.Exit(1)
    }
}

func renderWelcome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the golang-neo4j-example REST api!")
}