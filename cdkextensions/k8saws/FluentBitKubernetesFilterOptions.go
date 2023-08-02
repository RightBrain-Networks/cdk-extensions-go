package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// Options for configuring the Kubernetes Fluent Bit filter plugin.
// See: [Kubernetes Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/kubernetes)
//
type FluentBitKubernetesFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Include Kubernetes resource annotations in the extra metadata.
	// Default: true.
	//
	Annotations *bool `field:"optional" json:"annotations" yaml:"annotations"`
	// Set the buffer size for HTTP client when reading responses from Kubernetes API server.
	//
	// A value of 0 results in no limit, and the buffer will expand as-needed.
	//
	// Note that if pod specifications exceed the buffer limit, the API
	// response will be discarded when retrieving metadata, and some kubernetes
	// metadata will fail to be injected to the logs.
	// Default: 32k.
	//
	BufferSize core.DataSize `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// When enabled, metadata will be fetched from K8s when docker_id is changed.
	// Default: false.
	//
	CacheUseDockerId *bool `field:"optional" json:"cacheUseDockerId" yaml:"cacheUseDockerId"`
	// DNS lookup retries N times until the network starts working.
	// Default: 6.
	//
	DnsRetries *float64 `field:"optional" json:"dnsRetries" yaml:"dnsRetries"`
	// DNS lookup interval between network status checks.
	// Default: 30 seconds.
	//
	DnsWaitTime awscdk.Duration `field:"optional" json:"dnsWaitTime" yaml:"dnsWaitTime"`
	// If set, use dummy-meta data (for test/dev purposes).
	// Default: false.
	//
	DummyMeta *bool `field:"optional" json:"dummyMeta" yaml:"dummyMeta"`
	// Allow Kubernetes Pods to exclude their logs from the log processor.
	// Default: false.
	//
	K8sLoggingExclude *bool `field:"optional" json:"k8sLoggingExclude" yaml:"k8sLoggingExclude"`
	// Allow Kubernetes Pods to suggest a pre-defined Parser.
	// Default: false.
	//
	K8sLoggingParser *bool `field:"optional" json:"k8sLoggingParser" yaml:"k8sLoggingParser"`
	// When `keepLog` is disabled, the log field is removed from the incoming message once it has been successfully merged (`mergeLog` must be enabled as well).
	// Default: true.
	//
	KeepLog *bool `field:"optional" json:"keepLog" yaml:"keepLog"`
	// CA certificate file.
	// Default: '/var/run/secrets/kubernetes.io/serviceaccount/ca.crt'
	//
	KubeCaFile *string `field:"optional" json:"kubeCaFile" yaml:"kubeCaFile"`
	// Absolute path to scan for certificate files.
	KubeCaPath *string `field:"optional" json:"kubeCaPath" yaml:"kubeCaPath"`
	// Kubelet host using for HTTP request, this only works when `useKubelet` is enabled.
	KubeletHost *string `field:"optional" json:"kubeletHost" yaml:"kubeletHost"`
	// Kubelet port using for HTTP request, this only works when `useKubelet` is enabled.
	// Default: 10250.
	//
	KubeletPort *float64 `field:"optional" json:"kubeletPort" yaml:"kubeletPort"`
	// Configurable TTL for K8s cached metadata.
	//
	// By default, it is set to 0 which means TTL for cache entries is disabled
	// and cache entries are evicted at random when capacity is reached.
	//
	// In order to enable this option, you should set the number to a time
	// interval.
	// Default: 0.
	//
	KubeMetaCacheTtl awscdk.Duration `field:"optional" json:"kubeMetaCacheTtl" yaml:"kubeMetaCacheTtl"`
	// If set, Kubernetes meta-data can be cached/pre-loaded from files in JSON format in this directory, named as namespace-pod.meta.
	KubeMetaPreloadCacheDir *string `field:"optional" json:"kubeMetaPreloadCacheDir" yaml:"kubeMetaPreloadCacheDir"`
	// When the source records comes from Tail input plugin, this option allows to specify what's the prefix used in Tail configuration.
	// Default: 'kube.var.log.containers.'
	//
	KubeTagPrefix *string `field:"optional" json:"kubeTagPrefix" yaml:"kubeTagPrefix"`
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
	KubeTokenCommand *string `field:"optional" json:"kubeTokenCommand" yaml:"kubeTokenCommand"`
	// Token file.
	// Default: '/var/run/secrets/kubernetes.io/serviceaccount/token'
	//
	KubeTokenFile *string `field:"optional" json:"kubeTokenFile" yaml:"kubeTokenFile"`
	// Configurable 'time to live' for the K8s token.
	//
	// After this time, the token is reloaded from `kubeTokenFile` or the
	// `kubeTokenCommand`.
	// Default: 10 minutes.
	//
	KubeTokenTtl awscdk.Duration `field:"optional" json:"kubeTokenTtl" yaml:"kubeTokenTtl"`
	// API Server end-point.
	// Default: 'https://kubernetes.default.svc/'
	//
	KubeUrl *string `field:"optional" json:"kubeUrl" yaml:"kubeUrl"`
	// Include Kubernetes resource labels in the extra metadata.
	// Default: true.
	//
	Labels *bool `field:"optional" json:"labels" yaml:"labels"`
	// When enabled, it checks if the `log` field content is a JSON string map, if so, it append the map fields as part of the log structure.
	// Default: false.
	//
	MergeLog *bool `field:"optional" json:"mergeLog" yaml:"mergeLog"`
	// When `mergeLog` is enabled, the filter tries to assume the `log` field from the incoming message is a JSON string message and make a structured representation of it at the same level of the `log` field in the map.
	//
	// Now if `mergeLogKey` is set (a string name), all the new structured
	// fields taken from the original `log` content are inserted under the new
	// key.
	MergeLogKey *string `field:"optional" json:"mergeLogKey" yaml:"mergeLogKey"`
	// When Merge_Log is enabled, trim (remove possible \n or \r) field values.
	// Default: true.
	//
	MergeLogTrim *bool `field:"optional" json:"mergeLogTrim" yaml:"mergeLogTrim"`
	// Optional parser name to specify how to parse the data contained in the log key.
	//
	// Recommended use is for developers or testing only.
	MergeParser *string `field:"optional" json:"mergeParser" yaml:"mergeParser"`
	// Set an alternative Parser to process record Tag and extract pod_name, namespace_name, container_name and docker_id.
	//
	// The parser must be registered in a parsers file.
	// See: [Parsers File](https://github.com/fluent/fluent-bit/blob/master/conf/parsers.conf)
	//
	RegexParser *string `field:"optional" json:"regexParser" yaml:"regexParser"`
	// Debug level between 0 (nothing) and 4 (every detail).
	// Default: -1.
	//
	TlsDebug *float64 `field:"optional" json:"tlsDebug" yaml:"tlsDebug"`
	// When enabled, turns on certificate validation when connecting to the Kubernetes API server.
	// Default: true.
	//
	TlsVerify *bool `field:"optional" json:"tlsVerify" yaml:"tlsVerify"`
	// When enabled, the filter reads logs coming in Journald format.
	// Default: false.
	//
	UseJournal *bool `field:"optional" json:"useJournal" yaml:"useJournal"`
	// This is an optional feature flag to get metadata information from kubelet instead of calling Kube Server API to enhance the log.
	// See: [Kube API heavy traffic issue for large cluster](https://docs.fluentbit.io/manual/pipeline/filters/kubernetes#optional-feature-using-kubelet-to-get-metadata)
	//
	// Default: false.
	//
	UseKubelet *bool `field:"optional" json:"useKubelet" yaml:"useKubelet"`
}

