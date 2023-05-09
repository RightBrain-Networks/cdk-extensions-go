package ec2


type CidrAssignmentBindOptions struct {
	MaxNetmask *float64 `field:"optional" json:"maxNetmask" yaml:"maxNetmask"`
	MinNetmask *float64 `field:"optional" json:"minNetmask" yaml:"minNetmask"`
}

