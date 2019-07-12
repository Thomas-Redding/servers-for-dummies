
package main

import (
  "errors"
  "fmt"
  "io/ioutil"

  "cloud.google.com/go/storage"
  "golang.org/x/net/context"
  "google.golang.org/appengine/log"
)

type FileSystem struct {
  _ctx context.Context
  _client *storage.Client
  _bucketName string
  _bucketHandle *storage.BucketHandle
}

func connectToFileSystem(ctx context.Context, client *storage.Client, bucketName string) *FileSystem {
  bucketHandle := client.Bucket(bucketName)
  return &FileSystem{_ctx: ctx, _client: client, _bucketName: bucketName, _bucketHandle: bucketHandle}
}

/*
 Overwrite a file at `filePath` with `newContents`. If no file exists, create a new one.
 */
func (fileSystem *FileSystem) write(filePath string, newContents []byte) error {
  wc := fileSystem._bucketHandle.Object(filePath).NewWriter(fileSystem._ctx)
  wc.ContentType = "text/plain"
  wc.Metadata = map[string]string{
    "x-goog-meta-foo": "foo",
    "x-goog-meta-bar": "bar",
  }
  if _, err := wc.Write(newContents); err != nil {
    errorString := fmt.Sprintf("createFile: unable to write data to bucket %q, file %q: %v", fileSystem._bucketName, filePath, err)
    log.Errorf(fileSystem._ctx, errorString)
    return errors.New(errorString)
  }
  if err := wc.Close(); err != nil {
    errorString := fmt.Sprintf("createFile: unable to close bucket %q, file %q: %v", fileSystem._bucketName, filePath, err)
    log.Errorf(fileSystem._ctx, errorString)
    return errors.New(errorString)
  }
  return nil
}

func (fileSystem *FileSystem) read(filePath string) ([]byte, error) {
  rc, err := fileSystem._bucketHandle.Object(filePath).NewReader(fileSystem._ctx)
  if err != nil {
    errorString := fmt.Sprintf("readFile: unable to open file from bucket %q, file %q: %v", fileSystem._bucketName, filePath, err)
    log.Errorf(fileSystem._ctx, errorString)
    return nil, errors.New(errorString)
  }
  defer rc.Close()
  slurp, err := ioutil.ReadAll(rc)
  if err != nil {
    errorString := fmt.Sprintf("readFile: unable to read data from bucket %q, file %q: %v", fileSystem._bucketName, filePath, err)
    log.Errorf(fileSystem._ctx, errorString)
    return nil, errors.New(errorString)
  }
  return slurp, nil
}
