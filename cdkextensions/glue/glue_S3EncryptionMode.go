package glue


// Encryption mode for S3.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_S3Encryption.html#Glue-Type-S3Encryption-S3EncryptionMode
//
type S3EncryptionMode string

const (
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	S3EncryptionMode_KMS S3EncryptionMode = "KMS"
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	S3EncryptionMode_S3_MANAGED S3EncryptionMode = "S3_MANAGED"
)

