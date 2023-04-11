//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose"
)

func validateFluentBitOutput_CloudwatchLogsParameters(match FluentBitMatch, logGroup awslogs.ILogGroup) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

func validateFluentBitOutput_KinesisParameters(match FluentBitMatch, stream awskinesis.IStream) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}

func validateFluentBitOutput_KinesisFirehoseParameters(match FluentBitMatch, deliveryStream kinesisfirehose.IDeliveryStream) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if deliveryStream == nil {
		return fmt.Errorf("parameter deliveryStream is required, but nil was provided")
	}

	return nil
}

func validateFluentBitOutput_OpensearchParameters(match FluentBitMatch, domain awsopensearchservice.IDomain) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

