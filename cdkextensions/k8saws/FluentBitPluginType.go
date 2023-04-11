package k8saws


// The types of Fluent Bit plugins that can be configured.
type FluentBitPluginType string

const (
	// A plugin that transforms or filters records.
	FluentBitPluginType_FILTER FluentBitPluginType = "FILTER"
	// A plugin that configures where output should be sent.
	FluentBitPluginType_OUTPUT FluentBitPluginType = "OUTPUT"
	// A plugin that read data from input objects into structured objects.
	FluentBitPluginType_PARSER FluentBitPluginType = "PARSER"
)

