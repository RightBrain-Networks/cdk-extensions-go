//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeliveryStreamDestination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

