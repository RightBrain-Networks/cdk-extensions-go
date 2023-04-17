package athena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type AnalyticsEngineOutputOptions struct {
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	Encryption IAthenaResultEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	ExpectedOwnerId *string `field:"optional" json:"expectedOwnerId" yaml:"expectedOwnerId"`
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

