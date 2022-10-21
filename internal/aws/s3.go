package aws

import (
	"context"
	"io"

	"github.com/SphericalKat/katbox/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/minio/minio-go/v7"
	creds "github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

var S3Client *s3.Client
var uploader *manager.Uploader
var MC *minio.Client

func Connect() {
	s3Config, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion("ap-southeast-2"),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.Conf.S3AccessKey, config.Conf.S3SecretKey, "")),
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "https://s3.ap-southeast-2.wasabisys.com",
				SigningRegion: "ap-southeast-2",
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("error connecting to s3: %v", err)
	}

	S3Client = s3.NewFromConfig(s3Config)
	log.Info("Connected to S3")

	uploader = manager.NewUploader(S3Client)
}

func ConnectMinio() {
	endpoint := "s3.ap-southeast-2.wasabisys.com"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  creds.NewStaticV4(config.Conf.S3AccessKey, config.Conf.S3SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("error connecting to s3: %v", err)
	}

	bucketName := "katbox"
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

func UploadFile(name, contentType string, file io.Reader) (string, error) {
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(config.Conf.S3BucketName),
		Key:         aws.String(name),
		ContentType: aws.String(contentType),
		Body:        file,
	})
	if err != nil {
		return "", err
	}

	return *result.Key, nil
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
