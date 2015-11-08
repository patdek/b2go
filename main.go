package main

import (
	"fmt"

	"github.com/patdek/b2go"
)

func main() {
	usage()
}

func authorizeAccount(accountID, applicationKey string) {
}

func clearAccount() {
}

func createBucket(bucketName, bucketType string) {
}

func updateBucket(bucketID, bucketType string) {
}

func downloadFileByID(fileID, localFileName string) {
}

func listBuckets() {
}

func makeURL(fileID string) {
}

func uploadFile(bucketID, localFilePath, b2FileName, contentType string, fileInfo b2go.FileInfo) {

}

func usage() {
	fmt.Println(`
This program provides command-line access to the B2 service.

Usages:

    b2 authorize_account [accountId] [applicationKey]

        Prompts for Backblaze accountID and applicationKey (unless they are given
        on the command line).

        The account ID is a 12-digit hex number that you can get from
        your account page on backblaze.com.

        The application key is a 40-digit hex number that you can get from
        your account page on backblaze.com.

        Stores an account auth token in ~/.b2_account_info

    b2 clear_account

        Erases everything in ~/.b2_account_info

    b2 create_bucket <bucketName> <bucketType>

        Creates a new bucket.  Prints the ID of the bucket created.

    b2 update_bucket <bucketId> <bucketType>

        Updates the bucketType of an existing bucket.  Prints the ID
        of the bucket updated.

    b2 download_file_by_id <fileId> <localFileName>

        Downloads the given file, and stores it in the given local file.

    b2 list_buckets

        Lists all of the buckets in the current account.

    b2 make_url <fileId>

        Prints an URL that can be used to download the given file, if
        it is public.

    b2 upload_file <bucketId> <localFilePath> <b2FileName> <contentType> [<fileInfo> ...]

        Uploads one file to the given bucket.  Uploads the contents
        of the local file, and assigns the given name to the B2 file.

        Each fileInfo is of the form "a=b".

    b2 version

        Echos the version number of this program.`)
}
