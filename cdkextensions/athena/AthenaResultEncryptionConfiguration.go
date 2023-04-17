package athena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type AthenaResultEncryptionConfiguration struct {
	BucketEncryption awss3.BucketEncryption `field:"required" json:"bucketEncryption" yaml:"bucketEncryption"`
	EncryptionLabel *string `field:"required" json:"encryptionLabel" yaml:"encryptionLabel"`
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

