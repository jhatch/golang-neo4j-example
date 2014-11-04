package main

import (
  "net/http"
  "log"
  "strings"
  "io/ioutil"
  "os"
)

const PORT = "8000"

func main() {
  // serve static files
  fs := http.FileServer(http.Dir("client"))
  http.Handle("/", fs)

  // proxy the api
  http.HandleFunc("/api", proxyAPI)

  // start 'er up
  log.Printf("STARTING golang server on port %s\n", PORT)
  err := http.ListenAndServe(":"+PORT, nil)
  if err != nil {
    log.Fatalf("ERROR STARTING golang server on port %s\n", PORT, err)
    os.Exit(1)
  }
}

func proxyAPI(w http.ResponseWriter, r *http.Request) {
  response, err := http.Get("http://localhost:8001" + strings.Replace(r.URL.Path, "/api", "", -1))
  if err != nil {
    log.Fatalf("ERROR proxying to api server http://localhost:8001", err)
  } else {
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Printf("%s", err)
      os.Exit(1)
    }
    w.Write([]byte(string(contents)))
  }
}