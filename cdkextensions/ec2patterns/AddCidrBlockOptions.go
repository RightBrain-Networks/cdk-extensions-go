package ec2patterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type AddCidrBlockOptions struct {
	CidrAssignment ec2.ICidrAssignment `field:"required" json:"cidrAssignment" yaml:"cidrAssignment"`
}

