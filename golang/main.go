
package main

import (
  "google.golang.org/appengine"
  "net/http"
)

func main() {
  http.HandleFunc("/", handle)
  appengine.Main()
}

func handle(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-type", "text/plain")
  writer.Write([]byte("Hello World (Golang)!\n" + request.URL.Path))
}
