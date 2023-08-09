package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type RuleSet interface {
	Bind(scope constructs.IConstruct) *ScopedRuleSet
}

// The jsii proxy struct for RuleSet
type jsiiProxy_RuleSet struct {
	_ byte // padding
}

func RuleSet_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

func RuleSet_CIS_FOUNDATIONS_1_2_0() RuleSet {
	_init_.Initialize()
	var returns RuleSet
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"CIS_FOUNDATIONS_1_2_0",
		&returns,
	)
	return returns
}

func RuleSet_CIS_FOUNDATIONS_1_4_0() RuleSet {
	_init_.Initialize()
	var returns RuleSet
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"CIS_FOUNDATIONS_1_4_0",
		&returns,
	)
	return returns
}

func RuleSet_FOUNDATIONAL_BEST_PRACTICES_1_0_0() RuleSet {
	_init_.Initialize()
	var returns RuleSet
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"FOUNDATIONAL_BEST_PRACTICES_1_0_0",
		&returns,
	)
	return returns
}

func RuleSet_NIST_800_53_5_0_0() RuleSet {
	_init_.Initialize()
	var returns RuleSet
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"NIST_800_53_5_0_0",
		&returns,
	)
	return returns
}

func RuleSet_PCI_DSS_3_2_1() RuleSet {
	_init_.Initialize()
	var returns RuleSet
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.RuleSet",
		"PCI_DSS_3_2_1",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RuleSet) Bind(scope constructs.IConstruct) *ScopedRuleSet {
	if err := r.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ScopedRuleSet

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

