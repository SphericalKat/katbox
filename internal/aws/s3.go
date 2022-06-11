package aws

import (
	"context"

	"github.com/SphericalKat/katbox/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
)

var S3Client *s3.Client

func Connect() {
	// s3Config := aws.Config {
	// 	Credentials: credentials.NewStaticCredentials(config.Conf.S3AccessKey, config.Conf.S3SecretKey, ""),
	// 	Endpoint: aws.String("https://s3.wasabisys.com"),
	// 	Region: aws.String("ap-southeast-2"),
	// 	S3ForcePathStyle: aws.Bool(true),
	// }

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
}
