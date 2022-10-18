package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// S3 encryption configuration.
type S3Encryption struct {
	// Encryption mode.
	Mode S3EncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

