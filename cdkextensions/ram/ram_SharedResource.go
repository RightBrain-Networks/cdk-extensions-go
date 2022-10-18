package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

type SharedResource interface {
	ISharedResource
	Bind(_scope constructs.IConstruct) *string
}

// The jsii proxy struct for SharedResource
type jsiiProxy_SharedResource struct {
	jsiiProxy_ISharedResource
}

func SharedResource_FromArn(arn *string) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromArnParameters(arn); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedResource",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func SharedResource_FromProject(project awscodebuild.IProject) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromProjectParameters(project); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedResource",
		"fromProject",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func SharedResource_FromSubnet(subnet awsec2.ISubnet) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromSubnetParameters(subnet); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedResource",
		"fromSubnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedResource) Bind(_scope constructs.IConstruct) *string {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

