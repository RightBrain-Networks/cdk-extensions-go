package athena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

type AnalyticsEngineConfiguration struct {
	EncrpytionKey awskms.IKey `field:"optional" json:"encrpytionKey" yaml:"encrpytionKey"`
	EnforceConfiguration *bool `field:"optional" json:"enforceConfiguration" yaml:"enforceConfiguration"`
	EngineVersion AnalyticsEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	PublishMetrics *bool `field:"optional" json:"publishMetrics" yaml:"publishMetrics"`
	QueryScannedBytesLimit core.DataSize `field:"optional" json:"queryScannedBytesLimit" yaml:"queryScannedBytesLimit"`
	RequesterPays *bool `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	ResultsBucket awss3.IBucket `field:"optional" json:"resultsBucket" yaml:"resultsBucket"`
	ResultsBucketEncryptionKey awskms.IKey `field:"optional" json:"resultsBucketEncryptionKey" yaml:"resultsBucketEncryptionKey"`
	ResultsBucketEncryptionType *string `field:"optional" json:"resultsBucketEncryptionType" yaml:"resultsBucketEncryptionType"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

