package glue


// Encryption mode for CloudWatch Logs.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_CloudWatchEncryption.html#Glue-Type-CloudWatchEncryption-CloudWatchEncryptionMode
//
type CloudWatchEncryptionMode string

const (
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	CloudWatchEncryptionMode_KMS CloudWatchEncryptionMode = "KMS"
)

