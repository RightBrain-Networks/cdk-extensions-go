package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type GuardDutySeverity interface {
	LowerBound() *float64
	Standardized() *string
	UpperBound() *float64
	BuildCondition(path *string) awsstepfunctions.Condition
}

// The jsii proxy struct for GuardDutySeverity
type jsiiProxy_GuardDutySeverity struct {
	_ byte // padding
}

func (j *jsiiProxy_GuardDutySeverity) LowerBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lowerBound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardDutySeverity) Standardized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardDutySeverity) UpperBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upperBound",
		&returns,
	)
	return returns
}


func GuardDutySeverity_All() *[]GuardDutySeverity {
	_init_.Initialize()

	var returns *[]GuardDutySeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.GuardDutySeverity",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func GuardDutySeverity_Custom(levels ...GuardDutySeverity) *[]GuardDutySeverity {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range levels {
		args = append(args, a)
	}

	var returns *[]GuardDutySeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.GuardDutySeverity",
		"custom",
		args,
		&returns,
	)

	return returns
}

func GuardDutySeverity_Of(standardized *string, lowerBound *float64, upperBound *float64) GuardDutySeverity {
	_init_.Initialize()

	if err := validateGuardDutySeverity_OfParameters(standardized); err != nil {
		panic(err)
	}
	var returns GuardDutySeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.GuardDutySeverity",
		"of",
		[]interface{}{standardized, lowerBound, upperBound},
		&returns,
	)

	return returns
}

func GuardDutySeverity_Threshold(level GuardDutySeverity) *[]GuardDutySeverity {
	_init_.Initialize()

	if err := validateGuardDutySeverity_ThresholdParameters(level); err != nil {
		panic(err)
	}
	var returns *[]GuardDutySeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.GuardDutySeverity",
		"threshold",
		[]interface{}{level},
		&returns,
	)

	return returns
}

func GuardDutySeverity_CRITICAL() GuardDutySeverity {
	_init_.Initialize()
	var returns GuardDutySeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.GuardDutySeverity",
		"CRITICAL",
		&returns,
	)
	return returns
}

func GuardDutySeverity_HIGH() GuardDutySeverity {
	_init_.Initialize()
	var returns GuardDutySeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.GuardDutySeverity",
		"HIGH",
		&returns,
	)
	return returns
}

func GuardDutySeverity_INFORMATIONAL() GuardDutySeverity {
	_init_.Initialize()
	var returns GuardDutySeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.GuardDutySeverity",
		"INFORMATIONAL",
		&returns,
	)
	return returns
}

func GuardDutySeverity_LOW() GuardDutySeverity {
	_init_.Initialize()
	var returns GuardDutySeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.GuardDutySeverity",
		"LOW",
		&returns,
	)
	return returns
}

func GuardDutySeverity_MEDIUM() GuardDutySeverity {
	_init_.Initialize()
	var returns GuardDutySeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.GuardDutySeverity",
		"MEDIUM",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GuardDutySeverity) BuildCondition(path *string) awsstepfunctions.Condition {
	if err := g.validateBuildConditionParameters(path); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Condition

	_jsii_.Invoke(
		g,
		"buildCondition",
		[]interface{}{path},
		&returns,
	)

	return returns
}

