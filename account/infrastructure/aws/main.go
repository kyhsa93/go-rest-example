package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/aws/s3"
	"github.com/kyhsa93/gin-rest-example/config"
)

// AWS aws struct
type AWS struct {
	secretID   string
	secretKey  string
	token      string
	S3         *s3.S3
	s3Endpoint string
	s3Region   string
}

// New create AWS instance
func New(config *config.Config) *AWS {
	awsInfra := &AWS{
		s3Endpoint: config.AWS.S3.Endpoint,
		s3Region:   config.AWS.S3.Region,
		secretID:   config.AWS.EnvironmentValue.SecretID,
		secretKey:  config.AWS.EnvironmentValue.SecretKey,
		token:      config.AWS.EnvironmentValue.Token,
	}
	awsInfra.S3 = s3.New(config, awsInfra.getAWSSession())
	return awsInfra
}

func (awsInfra *AWS) awsEndpointResolver(
	service,
	region string,
	optFns ...func(*endpoints.Options),
) (endpoints.ResolvedEndpoint, error) {
	if service == endpoints.S3ServiceID {
		return endpoints.ResolvedEndpoint{
			URL:           awsInfra.s3Endpoint,
			SigningRegion: awsInfra.s3Region,
		}, nil
	}

	return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
}

func (awsInfra *AWS) getAWSSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(endpoints.ApNortheast2RegionID),
		EndpointResolver: endpoints.ResolverFunc(awsInfra.awsEndpointResolver),
		Credentials: credentials.NewStaticCredentials(
			awsInfra.secretID,
			awsInfra.secretKey,
			awsInfra.token,
		),
		S3ForcePathStyle: aws.Bool(true),
	}))
}