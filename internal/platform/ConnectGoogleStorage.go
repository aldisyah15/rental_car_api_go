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
	backetName := "rental-car-app"
	bucket := client.Bucket(backetName)

	log.Printf("Bucket: %v", bucket.BucketName())
}
