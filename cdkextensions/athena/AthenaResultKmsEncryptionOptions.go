package athena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

type AthenaResultKmsEncryptionOptions struct {
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

