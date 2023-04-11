package aps


// Configuration for the aler manager route.
type AlertManagerReceiverProps struct {
	// Details for alerting providers where events routed to this receiver should be sent,.
	Destinations *[]IAlertManagerDestination `field:"optional" json:"destinations" yaml:"destinations"`
	// The name of the receiver which can be referenced in the other parts of the configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

