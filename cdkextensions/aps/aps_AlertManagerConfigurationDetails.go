package aps


// The details that are needed to configure alert manager for an Amazon APS workspace.
type AlertManagerConfigurationDetails struct {
	// The rendered alert manager configuration in the format expected by an Amazon APS workspace.
	Contents *string `field:"required" json:"contents" yaml:"contents"`
}

