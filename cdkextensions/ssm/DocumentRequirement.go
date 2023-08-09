package ssm


type DocumentRequirement struct {
	Document IDocument `field:"required" json:"document" yaml:"document"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

