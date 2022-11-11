package k8saws


// Options that are applicable to all Fluent Bit Plugins regardless of type.
type FluentBitPluginCommonOptions struct {
	// The name of the fluent bit plugin.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	PluginType FluentBitPluginType `field:"required" json:"pluginType" yaml:"pluginType"`
}

