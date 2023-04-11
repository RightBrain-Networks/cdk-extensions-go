package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
)

type DataFormatConversionOptions struct {
	Database glue.Database `field:"required" json:"database" yaml:"database"`
	InputFormat InputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	OutputFormat OutputFormat `field:"required" json:"outputFormat" yaml:"outputFormat"`
	Table glue.Table `field:"required" json:"table" yaml:"table"`
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	Version TableVersion `field:"optional" json:"version" yaml:"version"`
}

