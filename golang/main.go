
package main

import (
  "fmt"
  "net/http"

  "cloud.google.com/go/storage"
  "google.golang.org/appengine"
  "google.golang.org/appengine/file"
  "google.golang.org/appengine/log"
)

func main() {
  http.HandleFunc("/", handle)
  appengine.Main()
}

func handle(writer http.ResponseWriter, request *http.Request) {
  ctx := appengine.NewContext(request)
  bucketName, err := file.DefaultBucketName(ctx)
  if err != nil {
    log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
  }
  writer.Header().Set("Content-type", "text/plain")
  writer.Write([]byte("Hello World (Golang)!\n" + request.URL.Path))

  client, err := storage.NewClient(ctx)
  if err != nil {
    log.Errorf(ctx, "failed to create client: %v", err)
    return
  }
  defer client.Close()

  fileSystem := connectToFileSystem(ctx, client, bucketName)

  writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
  fmt.Fprintf(writer, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
  fmt.Fprintf(writer, "Using bucket name: %v\n\n", bucketName)

  fileSystem.write("demo-testfile-go", []byte("lorem lorem 123"))
  fileData, _ := fileSystem.read("demo-testfile-go")
  fmt.Fprintf(writer, "File data: %v\n\n", string(fileData))
  fmt.Fprintf(writer, "temp: %v\n\n", string(fileData))
}
