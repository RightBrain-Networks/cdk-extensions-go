package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type OpenSearchEventSeverity interface {
	Original() *string
	Priority() *float64
	Standardized() *string
	BuildCondition(path *string) awsstepfunctions.Condition
}

// The jsii proxy struct for OpenSearchEventSeverity
type jsiiProxy_OpenSearchEventSeverity struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchEventSeverity) Original() *string {
	var returns *string
	_jsii_.Get(
		j,
		"original",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchEventSeverity) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchEventSeverity) Standardized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardized",
		&returns,
	)
	return returns
}


func OpenSearchEventSeverity_All() *[]OpenSearchEventSeverity {
	_init_.Initialize()

	var returns *[]OpenSearchEventSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func OpenSearchEventSeverity_Custom(levels ...OpenSearchEventSeverity) *[]OpenSearchEventSeverity {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range levels {
		args = append(args, a)
	}

	var returns *[]OpenSearchEventSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"custom",
		args,
		&returns,
	)

	return returns
}

func OpenSearchEventSeverity_Of(standardized *string, original *string, priority *float64) OpenSearchEventSeverity {
	_init_.Initialize()

	if err := validateOpenSearchEventSeverity_OfParameters(standardized, original, priority); err != nil {
		panic(err)
	}
	var returns OpenSearchEventSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"of",
		[]interface{}{standardized, original, priority},
		&returns,
	)

	return returns
}

func OpenSearchEventSeverity_Threshold(level OpenSearchEventSeverity) *[]OpenSearchEventSeverity {
	_init_.Initialize()

	if err := validateOpenSearchEventSeverity_ThresholdParameters(level); err != nil {
		panic(err)
	}
	var returns *[]OpenSearchEventSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"threshold",
		[]interface{}{level},
		&returns,
	)

	return returns
}

func OpenSearchEventSeverity_HIGH() OpenSearchEventSeverity {
	_init_.Initialize()
	var returns OpenSearchEventSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"HIGH",
		&returns,
	)
	return returns
}

func OpenSearchEventSeverity_INFORMATIONAL() OpenSearchEventSeverity {
	_init_.Initialize()
	var returns OpenSearchEventSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"INFORMATIONAL",
		&returns,
	)
	return returns
}

func OpenSearchEventSeverity_LOW() OpenSearchEventSeverity {
	_init_.Initialize()
	var returns OpenSearchEventSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"LOW",
		&returns,
	)
	return returns
}

func OpenSearchEventSeverity_MEDIUM() OpenSearchEventSeverity {
	_init_.Initialize()
	var returns OpenSearchEventSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		"MEDIUM",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpenSearchEventSeverity) BuildCondition(path *string) awsstepfunctions.Condition {
	if err := o.validateBuildConditionParameters(path); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Condition

	_jsii_.Invoke(
		o,
		"buildCondition",
		[]interface{}{path},
		&returns,
	)

	return returns
}

