package platform

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

func ConnectGoogleStorage() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()
	BacketName := "rental-car-app"
	bucket := client.Bucket(BacketName)

	log.Printf("Bucket: %v", bucket.BucketName())
}
