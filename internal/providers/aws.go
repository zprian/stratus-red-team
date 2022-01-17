package providers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

type StratusAwsService string

type AwsProvider interface {
	IsAuthenticated() bool
	S3Client() *s3.Client
	EC2Client() *ec2.Client
	STSClient() *sts.Client
	IAMClient() *iam.Client
}

type awsProviderImpl struct {
	awsConfig aws.Config
}

func (m *awsProviderImpl) IsAuthenticated() bool {
	stsClient := m.STSClient()
	_, err := stsClient.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	return err == nil
}

func (m *awsProviderImpl) EC2Client() *ec2.Client {
	return ec2.NewFromConfig(m.awsConfig)
}

func (m *awsProviderImpl) S3Client() *s3.Client {
	return s3.NewFromConfig(m.awsConfig)
}

func (m *awsProviderImpl) IAMClient() *iam.Client {
	return iam.NewFromConfig(m.awsConfig)
}

func (m *awsProviderImpl) STSClient() *sts.Client {
	return sts.NewFromConfig(m.awsConfig)
}

func NewAwsProvider() AwsProvider {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load AWS configuration, %v", err)
	}

	return &awsProviderImpl{awsConfig: cfg}
}
