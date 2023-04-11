package k8saws


type LiftOptions struct {
	// Lift records nested under the this key.
	NestedUnder *string `field:"required" json:"nestedUnder" yaml:"nestedUnder"`
}

