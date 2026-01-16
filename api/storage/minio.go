package storage

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Minio *minio.Client

func InitMinio() {
	endpoint := os.Getenv("MINIO_ENDPOINT") // localhost:9000 หรือ minio:9000
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // dev
	})
	if err != nil {
		log.Fatal(err)
	}

	Minio = client
	log.Println("✅ MinIO connected")
}
