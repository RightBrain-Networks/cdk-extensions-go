package networkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISite interface {
	SiteArn() *string
	SiteId() *string
}

// The jsii proxy for ISite
type jsiiProxy_ISite struct {
	_ byte // padding
}

func (j *jsiiProxy_ISite) SiteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISite) SiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteId",
		&returns,
	)
	return returns
}

