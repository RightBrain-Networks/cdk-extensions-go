package guardduty


// Options for configuring Kubernetes as a data source for GuardDuty.
type KubernetesOptions struct {
	// Controls whether EKS audit logs should be used as a data source for GuardDuty.
	AuditLogs *bool `field:"optional" json:"auditLogs" yaml:"auditLogs"`
}

