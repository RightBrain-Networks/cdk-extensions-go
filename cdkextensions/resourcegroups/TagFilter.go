package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type TagFilter interface {
	IGroupConfiguration
	// Collection of resource types that are allowed to be in the Resource Group being configured.
	ResourceTypes() *[]*string
	// Collection of filters to be used to determine the resources that belong to the Resource Group.
	TagFilters() *map[string]*[]*string
	// Add a resource type to the resource group.
	//
	// If no resource types are registered in the configuration then all resource
	// types are allowed.
	//
	// Returns: The Resource Group configuration object to which the type was
	// registered.
	AddResourceType(typeId *string) TagFilter
	// Adds a new tag filter that should be used for controlling the resources in the Resource Group.
	//
	// Returns: The Resource Group configuration object to which the type was
	// registered.
	AddTagFilter(key *string, values ...*string) TagFilter
	// Associates this configuration with a construct that is handling the creation of a resource group.
	//
	// Returns: The configuration to be used for the Resource Group.
	Bind(_scope constructs.IConstruct) *BoundGroupConfiguration
}

// The jsii proxy struct for TagFilter
type jsiiProxy_TagFilter struct {
	jsiiProxy_IGroupConfiguration
}

func (j *jsiiProxy_TagFilter) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagFilter) TagFilters() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}


// Creates a new instance of the TagFilter class.
func NewTagFilter(props TagFilterProps) TagFilter {
	_init_.Initialize()

	j := jsiiProxy_TagFilter{}

	_jsii_.Create(
		"cdk-extensions.resourcegroups.TagFilter",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of the TagFilter class.
func NewTagFilter_Override(t TagFilter, props TagFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.resourcegroups.TagFilter",
		[]interface{}{props},
		t,
	)
}

func (t *jsiiProxy_TagFilter) AddResourceType(typeId *string) TagFilter {
	if err := t.validateAddResourceTypeParameters(typeId); err != nil {
		panic(err)
	}
	var returns TagFilter

	_jsii_.Invoke(
		t,
		"addResourceType",
		[]interface{}{typeId},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagFilter) AddTagFilter(key *string, values ...*string) TagFilter {
	if err := t.validateAddTagFilterParameters(key); err != nil {
		panic(err)
	}
	args := []interface{}{key}
	for _, a := range values {
		args = append(args, a)
	}

	var returns TagFilter

	_jsii_.Invoke(
		t,
		"addTagFilter",
		args,
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagFilter) Bind(_scope constructs.IConstruct) *BoundGroupConfiguration {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *BoundGroupConfiguration

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

