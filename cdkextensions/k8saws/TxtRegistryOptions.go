package k8saws


// Configuration options for setting up a TXT registry for ExternalDNS.
type TxtRegistryOptions struct {
	// A unique identifier that is used to establish ownership of managed DNS records.
	//
	// Prevents conflicts in the event of multiple clusters running external-dns.
	// Default: Unique address of the owning CDK node.
	//
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// A prefix to be added top TXT ownership records.
	//
	// By default, the ownership record is a TXT record with the same name as the
	// managed record that was created. This causes issues as some record types
	// (CNAME's) do not allow duplicate records of a different type.
	//
	// This prefix is used to prevent such name collissions while still allowing
	// DNS ownership records to be created.
	// Default: 'edns.''
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

