//go:build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func validateAthenaSqlOutputEncryption_CseKmsParameters(options *AthenaResultKmsEncryptionOptions) error {
	return nil
}

func validateAthenaSqlOutputEncryption_SseKmsParameters(options *AthenaResultKmsEncryptionOptions) error {
	return nil
}

