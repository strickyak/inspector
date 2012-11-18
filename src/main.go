package main

import (
  "os"
  "fmt"
  "net/http"
  "strconv"
)

import "inspect"

var _ = strconv.Atoi

func home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
  inspect.Register(r)
  inspect.Register(w)
}

func main() {
  port := os.Args[1]

  inspect.Root["args"] = os.Args
  inspect.Root["homeFunc"] = home

  http.HandleFunc("/", home)
  http.ListenAndServe("127.0.0.1:" + port, nil)
}
