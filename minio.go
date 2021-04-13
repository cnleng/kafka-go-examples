package main

import (
	"os"
    "log"
    "fmt"
    "io"
    "context"
"github.com/minio/minio-go/v7"
 //   "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "100.80.250.216:9001"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%#v\n", minioClient) // minioClient is now setup
    //minio/test-bucket/minio-file-object.txt
    object, err := minioClient.GetObject(context.Background(), "testbucket", "minio.go", minio.GetObjectOptions{})
    if err != nil {
      fmt.Println(err)
      return
    }
    localFile, err := os.Create("/tmp/local-import.txt")
    if err != nil {
       fmt.Println(err)
       return
    }
    log.Println(localFile)
    if _, err = io.Copy(localFile, object); err != nil {
       fmt.Println(err)
       return
    }
    log.Println(object)
}

