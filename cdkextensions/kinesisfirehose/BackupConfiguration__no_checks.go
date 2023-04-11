//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewBackupConfigurationParameters(options *BackupConfigurationOptions) error {
	return nil
}

