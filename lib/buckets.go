package b2go

// Buckets
// The storage for your account is grouped into buckets. Each bucket is a container that holds files. You can think of
// buckets as the top-level folders in your B2 Cloud Storage account. There is no limit to the number of files in a bucket.

// CreateBucket creates a new bucket in your account of name and type
func CreateBucket(bucketName, bucketType string) {
}

// DeleteBucket deletes the bucket specified. Only buckets that contain no version of any files can be deleted.
func DeleteBucket(bucketID string) {
}

// ListBuckets lists buckets associated with an account, in alphabetical order by bucket ID.
func ListBuckets() {
}

// UpdateBucket modifies the bucketType of an existing bucket. Can be used to allow everyone to download the contents of
// the bucket without providing any authorization, or to prevent anyone from downloading the contents of the
// bucket without providing a bucket auth token.
func UpdateBucket(bucketName, bucketType string) {
}

// DeleteBucketAndFiles will blow out a bucket and all the files required to delete it
func DeleteBucketAndFiles(bucketID string, areYourReallySure bool) {
}
