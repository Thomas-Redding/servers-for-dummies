
package main

import (
  "fmt"
  "net/http"
  "strconv"

  "cloud.google.com/go/storage"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
)

func main() {
  http.HandleFunc("/", handle)
  appengine.Main()
}

func handle(writer http.ResponseWriter, request *http.Request) {
  ctx := appengine.NewContext(request)
  client, err := storage.NewClient(ctx)
  if err != nil {
    log.Errorf(ctx, "failed to create client: %v", err)
    sendError(writer, 500, "Internal Server Error")
  	return
  }
  defer client.Close()

  fileSystem, err := connectToDefaultFileSystem(ctx, client)
  if err != nil {
    sendError(writer, 500, "Internal Server Error")
    return
  }
  fileSystem.write("test-file-golang.txt", []byte("Lorem ipsum dol..."))
  fileData, _ := fileSystem.read("test-file-golang.txt")

  writer.Header().Set("Content-type", "text/plain")
  writer.Write([]byte("Hello World (Golang)!\n"))
  writer.Write([]byte(request.URL.Path + "\n"))
  fmt.Fprintf(writer, "%v", string(fileData))
}

func sendError(writer http.ResponseWriter, errorCode int, message string) {
  writer.WriteHeader(errorCode)
  writer.Write([]byte("Error " + strconv.Itoa(errorCode) + ": " + message))
}
