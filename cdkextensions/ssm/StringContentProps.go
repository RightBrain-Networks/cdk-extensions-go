package ssm


type StringContentProps struct {
	Content *string `field:"required" json:"content" yaml:"content"`
	DocumentFormat DocumentFormat `field:"required" json:"documentFormat" yaml:"documentFormat"`
}

