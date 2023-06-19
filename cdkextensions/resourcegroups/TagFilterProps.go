package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Configuration options for configuring a Resource Group containing resources based on a set of matching tags.
type TagFilterProps interface {
	// The resource types that are allowed to be in the Resource Group being configured.
	ResourceTypes() *[]*string
	// The filters that should be used to determine the resources that belong to the resource group.
	TagFilters() *map[string]*[]*string
}

// The jsii proxy struct for TagFilterProps
type jsiiProxy_TagFilterProps struct {
	_ byte // padding
}

func (j *jsiiProxy_TagFilterProps) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagFilterProps) TagFilters() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}


func NewTagFilterProps() TagFilterProps {
	_init_.Initialize()

	j := jsiiProxy_TagFilterProps{}

	_jsii_.Create(
		"cdk-extensions.resourcegroups.TagFilterProps",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTagFilterProps_Override(t TagFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.resourcegroups.TagFilterProps",
		nil, // no parameters
		t,
	)
}

