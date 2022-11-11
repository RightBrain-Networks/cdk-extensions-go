package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// A Fluent Bit filter that allows log records to be annotated with Kubernetes metadata based on the containers that generated them.
type FluentBitKubernetesFilter interface {
	FluentBitFilterPluginBase
	// Include Kubernetes resource annotations in the extra metadata.
	Annotations() *bool
	// Set the buffer size for HTTP client when reading responses from Kubernetes API server.
	//
	// A value of 0 results in no limit, and the buffer will expand as-needed.
	//
	// Note that if pod specifications exceed the buffer limit, the API
	// response will be discarded when retrieving metadata, and some kubernetes
	// metadata will fail to be injected to the logs.
	BufferSize() core.DataSize
	// When enabled, metadata will be fetched from K8s when docker_id is changed.
	CacheUseDockerId() *bool
	// DNS lookup retries N times until the network starts working.
	DnsRetries() *float64
	// DNS lookup interval between network status checks.
	DnsWaitTime() awscdk.Duration
	// If set, use dummy-meta data (for test/dev purposes).
	DummyMeta() *bool
	// Allow Kubernetes Pods to exclude their logs from the log processor.
	K8sLoggingExclude() *bool
	// Allow Kubernetes Pods to suggest a pre-defined Parser.
	K8sLoggingParser() *bool
	// When `keepLog` is disabled, the log field is removed from the incoming message once it has been successfully merged (`mergeLog` must be enabled as well).
	KeepLog() *bool
	// CA certificate file.
	KubeCaFile() *string
	// Absolute path to scan for certificate files.
	KubeCaPath() *string
	// Kubelet host using for HTTP request, this only works when `useKubelet` is enabled.
	KubeletHost() *string
	// Kubelet port using for HTTP request, this only works when `useKubelet` is enabled.
	KubeletPort() *float64
	// Configurable TTL for K8s cached metadata.
	//
	// By default, it is set to 0 which means TTL for cache entries is disabled
	// and cache entries are evicted at random when capacity is reached.
	//
	// In order to enable this option, you should set the number to a time
	// interval.
	KubeMetaCacheTtl() awscdk.Duration
	// If set, Kubernetes meta-data can be cached/pre-loaded from files in JSON format in this directory, named as namespace-pod.meta.
	KubeMetaPreloadCacheDir() *string
	// When the source records comes from Tail input plugin, this option allows to specify what's the prefix used in Tail configuration.
	KubeTagPrefix() *string
	// Command to get Kubernetes authorization token.
	//
	// If you want to manually choose a command to get it, you can set the
	// command here.
	//
	// For example, run running the following to get the token using aws-cli:
	//
	// ```
	// aws-iam-authenticator -i your-cluster-name token --token-only
	// ```
	//
	// This option is currently Linux-only.
	KubeTokenCommand() *string
	// Token file.
	KubeTokenFile() *string
	// Configurable 'time to live' for the K8s token.
	//
	// After this time, the token is reloaded from `kubeTokenFile` or the
	// `kubeTokenCommand`.
	KubeTokenTtl() awscdk.Duration
	// API Server end-point.
	KubeUrl() *string
	// Include Kubernetes resource labels in the extra metadata.
	Labels() *bool
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// When enabled, it checks if the `log` field content is a JSON string map, if so, it append the map fields as part of the log structure.
	MergeLog() *bool
	// When `mergeLog` is enabled, the filter tries to assume the `log` field from the incoming message is a JSON string message and make a structured representation of it at the same level of the `log` field in the map.
	//
	// Now if `mergeLogKey` is set (a string name), all the new structured
	// fields taken from the original `log` content are inserted under the new
	// key.
	MergeLogKey() *string
	// When Merge_Log is enabled, trim (remove possible \n or \r) field values.
	MergeLogTrim() *bool
	// Optional parser name to specify how to parse the data contained in the log key.
	//
	// Recommended use is for developers or testing only.
	MergeParser() *string
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Set an alternative Parser to process record Tag and extract pod_name, namespace_name, container_name and docker_id.
	//
	// The parser must be registered in a parsers file.
	// See: [Parsers File](https://github.com/fluent/fluent-bit/blob/master/conf/parsers.conf)
	//
	RegexParser() *string
	// Debug level between 0 (nothing) and 4 (every detail).
	TlsDebug() *float64
	// When enabled, turns on certificate validation when connecting to the Kubernetes API server.
	TlsVerify() *bool
	// When enabled, the filter reads logs coming in Journald format.
	UseJournal() *bool
	// This is an optional feature flag to get metadata information from kubelet instead of calling Kube Server API to enhance the log.
	// See: [Kube API heavy traffic issue for large cluster](https://docs.fluentbit.io/manual/pipeline/filters/kubernetes#optional-feature-using-kubelet-to-get-metadata)
	//
	UseKubelet() *bool
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	//
	// Returns: A configuration for the plugin that con be used by the resource
	// configuring logging.
	Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// Renders a Fluent Bit configuration file for the plugin.
	//
	// Returns: A rendered plugin configuration file.
	RenderConfigFile(config *map[string]interface{}) *string
}

// The jsii proxy struct for FluentBitKubernetesFilter
type jsiiProxy_FluentBitKubernetesFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitKubernetesFilter) Annotations() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) BufferSize() core.DataSize {
	var returns core.DataSize
	_jsii_.Get(
		j,
		"bufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) CacheUseDockerId() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cacheUseDockerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) DnsRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dnsRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) DnsWaitTime() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"dnsWaitTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) DummyMeta() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dummyMeta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) K8sLoggingExclude() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"k8sLoggingExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) K8sLoggingParser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"k8sLoggingParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KeepLog() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"keepLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeCaPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeCaPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeletHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeletPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kubeletPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeMetaCacheTtl() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"kubeMetaCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeMetaPreloadCacheDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeMetaPreloadCacheDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeTagPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeTagPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeTokenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeTokenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeTokenFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeTokenTtl() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"kubeTokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) KubeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) Labels() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) MergeLog() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"mergeLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) MergeLogKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeLogKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) MergeLogTrim() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"mergeLogTrim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) MergeParser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) RegexParser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) TlsDebug() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsDebug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) TlsVerify() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"tlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) UseJournal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useJournal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKubernetesFilter) UseKubelet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useKubelet",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitKubernetesFilter class.
func NewFluentBitKubernetesFilter(options *FluentBitKubernetesFilterOptions) FluentBitKubernetesFilter {
	_init_.Initialize()

	if err := validateNewFluentBitKubernetesFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitKubernetesFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKubernetesFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitKubernetesFilter class.
func NewFluentBitKubernetesFilter_Override(f FluentBitKubernetesFilter, options *FluentBitKubernetesFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKubernetesFilter",
		[]interface{}{options},
		f,
	)
}

func FluentBitKubernetesFilter_PLUGIN_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.FluentBitKubernetesFilter",
		"PLUGIN_NAME",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FluentBitKubernetesFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
	if err := f.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ResolvedFluentBitConfiguration

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitKubernetesFilter) RenderConfigFile(config *map[string]interface{}) *string {
	if err := f.validateRenderConfigFileParameters(config); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"renderConfigFile",
		[]interface{}{config},
		&returns,
	)

	return returns
}

