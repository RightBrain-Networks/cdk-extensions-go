//go:build !no_runtime_type_checking

package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateApacheSparkOutputEncryption_SseKmsParameters(options *AthenaResultKmsEncryptionOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

