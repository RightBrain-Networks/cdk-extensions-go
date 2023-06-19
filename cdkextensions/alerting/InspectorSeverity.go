package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type InspectorSeverity interface {
	Original() *string
	Priority() *float64
	Standardized() *string
	BuildCondition(path *string) awsstepfunctions.Condition
}

// The jsii proxy struct for InspectorSeverity
type jsiiProxy_InspectorSeverity struct {
	_ byte // padding
}

func (j *jsiiProxy_InspectorSeverity) Original() *string {
	var returns *string
	_jsii_.Get(
		j,
		"original",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorSeverity) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InspectorSeverity) Standardized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardized",
		&returns,
	)
	return returns
}


func InspectorSeverity_All() *[]InspectorSeverity {
	_init_.Initialize()

	var returns *[]InspectorSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.InspectorSeverity",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func InspectorSeverity_Custom(levels ...InspectorSeverity) *[]InspectorSeverity {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range levels {
		args = append(args, a)
	}

	var returns *[]InspectorSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.InspectorSeverity",
		"custom",
		args,
		&returns,
	)

	return returns
}

func InspectorSeverity_Of(standardized *string, original *string, priority *float64) InspectorSeverity {
	_init_.Initialize()

	if err := validateInspectorSeverity_OfParameters(standardized, original, priority); err != nil {
		panic(err)
	}
	var returns InspectorSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.InspectorSeverity",
		"of",
		[]interface{}{standardized, original, priority},
		&returns,
	)

	return returns
}

func InspectorSeverity_Threshold(level InspectorSeverity) *[]InspectorSeverity {
	_init_.Initialize()

	if err := validateInspectorSeverity_ThresholdParameters(level); err != nil {
		panic(err)
	}
	var returns *[]InspectorSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.InspectorSeverity",
		"threshold",
		[]interface{}{level},
		&returns,
	)

	return returns
}

func InspectorSeverity_CRITICAL() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"CRITICAL",
		&returns,
	)
	return returns
}

func InspectorSeverity_HIGH() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"HIGH",
		&returns,
	)
	return returns
}

func InspectorSeverity_INFORMATIONAL() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"INFORMATIONAL",
		&returns,
	)
	return returns
}

func InspectorSeverity_LOW() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"LOW",
		&returns,
	)
	return returns
}

func InspectorSeverity_MEDIUM() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"MEDIUM",
		&returns,
	)
	return returns
}

func InspectorSeverity_UNTRIAGED() InspectorSeverity {
	_init_.Initialize()
	var returns InspectorSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.InspectorSeverity",
		"UNTRIAGED",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InspectorSeverity) BuildCondition(path *string) awsstepfunctions.Condition {
	if err := i.validateBuildConditionParameters(path); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Condition

	_jsii_.Invoke(
		i,
		"buildCondition",
		[]interface{}{path},
		&returns,
	)

	return returns
}

