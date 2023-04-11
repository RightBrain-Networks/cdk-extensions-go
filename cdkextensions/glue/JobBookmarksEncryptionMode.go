package glue


// Encryption mode for Job Bookmarks.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_JobBookmarksEncryption.html#Glue-Type-JobBookmarksEncryption-JobBookmarksEncryptionMode
//
type JobBookmarksEncryptionMode string

const (
	// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
	//
	JobBookmarksEncryptionMode_CLIENT_SIDE_KMS JobBookmarksEncryptionMode = "CLIENT_SIDE_KMS"
)

