//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IIpam) validateAddScopeParameters(id *string, options *PrivateIpamScopeOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IIpam) validateAssociateResourceDiscoveryParameters(resourceDiscovery IIpamResourceDiscovery) error {
	if resourceDiscovery == nil {
		return fmt.Errorf("parameter resourceDiscovery is required, but nil was provided")
	}

	return nil
}

