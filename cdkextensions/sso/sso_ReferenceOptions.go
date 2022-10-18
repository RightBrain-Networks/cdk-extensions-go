package sso


// Configuration options for creating a referenced customer managed policy.
type ReferenceOptions struct {
	// The name of the customer managed policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path for the policy.
	//
	// For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User
	// Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows a string of characters consisting of either a
	// forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from
	// the ! (`\u0021`) through the DEL character (`\u007F`), including most
	// punctuation characters, digits, and upper and lowercased letters.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

