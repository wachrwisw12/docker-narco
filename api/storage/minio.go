package storage

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	Minio  *minio.Client
	Bucket string
)

func InitMinio() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	Bucket = os.Getenv("MINIO_BUCKET")
	secure := os.Getenv("MINIO_SECURE") == "true"

	if endpoint == "" || accessKey == "" || secretKey == "" || Bucket == "" {
		log.Fatal("❌ MinIO env not set")
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, Bucket)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		if err := client.MakeBucket(ctx, Bucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatal(err)
		}
		log.Println("🪣 Bucket created:", Bucket)
	}

	Minio = client
	log.Println("✅ MinIO connected:", endpoint)
}
