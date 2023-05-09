package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type ICidrAssignment interface {
	GetCidr(scope constructs.IConstruct, id *string, options *CidrAssignmentBindOptions) *CidrAssignmentCidrDetails
	GetCidrOrIpamConfiguration(options *CidrAssignmentBindOptions) *CidrAssignmentDetails
}

// The jsii proxy for ICidrAssignment
type jsiiProxy_ICidrAssignment struct {
	_ byte // padding
}

func (i *jsiiProxy_ICidrAssignment) GetCidr(scope constructs.IConstruct, id *string, options *CidrAssignmentBindOptions) *CidrAssignmentCidrDetails {
	if err := i.validateGetCidrParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns *CidrAssignmentCidrDetails

	_jsii_.Invoke(
		i,
		"getCidr",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICidrAssignment) GetCidrOrIpamConfiguration(options *CidrAssignmentBindOptions) *CidrAssignmentDetails {
	if err := i.validateGetCidrOrIpamConfigurationParameters(options); err != nil {
		panic(err)
	}
	var returns *CidrAssignmentDetails

	_jsii_.Invoke(
		i,
		"getCidrOrIpamConfiguration",
		[]interface{}{options},
		&returns,
	)

	return returns
}

