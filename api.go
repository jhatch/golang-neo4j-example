package main

import (
  "fmt"
  "net/http"
  "net/url"
  "bytes"
  "log"
  "os"
)

const PORT = "8001"

func main() {
  // mount the api
  http.HandleFunc("/", renderWelcome)
  http.HandleFunc("/cypher/", runCypher)

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

// run a raw cypher against the local neo4j instance
// curl http://localhost:8001/cypher/MATCH%20(all:Database)%20RETURN%20all
func runCypher(w http.ResponseWriter, r *http.Request) {
  uri, _ := url.Parse(r.RequestURI)
  cypher := uri.Path[8:len(uri.Path)]
  fmt.Fprintln(w, "Executing cypher on neo4j", cypher)

  body := "{\"statements\" : [ {\"statement\" : \""+cypher+"\", \"parameters\" : {\"props\" : {\"name\" : \"My Node\"}}} ]}}"
  b := bytes.NewBufferString(body)
  
  neo4jRequest, _ := http.NewRequest("POST", "http://localhost:7474/db/data/transaction", b)
  neo4jRequest.Header.Add("Accept", "application/json; charset=UTF-8")
  neo4jRequest.Header.Add("Content-Type", "application/json")

  client := &http.Client{}
  res, err := client.Do(neo4jRequest)

  if err != nil {
    fmt.Fprintln(w, "error running cypher", err)
  } else {
    buf := new(bytes.Buffer)
    buf.ReadFrom(res.Body)
    s := buf.String()
    fmt.Fprintln(w, "Response:", s)
  }
}

