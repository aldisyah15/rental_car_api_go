package platform

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

func UploadFile(ctx context.Context, fileHeader *multipart.FileHeader, bucket string) (string, error) {
	// 1. Buka file dari FileHeader
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file: %w", err)
	}
	defer file.Close()

	// 2. Set timeout untuk upload
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	objectName := fmt.Sprintf("cars/%d_%s", time.Now().Unix(), fileHeader.Filename)

	o := client.Bucket(bucket).Object(objectName)
	wc := o.NewWriter(ctx)

	wc.ContentType = fileHeader.Header.Get("Content-Type")

	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %w", err)
	}

	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, objectName)
	return publicURL, nil
}
