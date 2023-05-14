package repositories

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
)

func SaveFile(mc *minio.Client, ctx context.Context, bucketName string, fileName string, file io.Reader, fileSize int64, contentType string) error {
	_, err := mc.PutObject(ctx, bucketName, fileName, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Print(err)
	}

	return err
}
