package storage

import (
	"context"
	"io"

	"github.com/SphericalKat/katbox/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

var MC *minio.Client

func ConnectMinio() {
	endpoint := "s3.ap-southeast-2.wasabisys.com"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.S3AccessKey, config.Conf.S3SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("error connecting to s3: %v", err)
	}

	bucketName := config.Conf.S3BucketName
	ctx := context.TODO()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "ap-southeast-2"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	MC = minioClient
}

func UploadMinio(ctx context.Context, name, contentType string, file io.Reader) (string, error) {
	info, err := MC.PutObject(
		context.TODO(),
		config.Conf.S3BucketName,
		name,
		file,
		-1,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	return info.Key, err
}
