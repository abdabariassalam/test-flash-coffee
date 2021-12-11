package aws

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (a AWS) GetUrl(fileName string) string {
	req, _ := s3.New(a.session).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(a.cfg.AWS.Bucket),
		Key:    aws.String(fileName),
	})
	rest.Build(req)
	uploadedResourceLocation := req.HTTPRequest.URL.String()
	return uploadedResourceLocation
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func (a AWS) AddFileToS3(file []byte, fileName, contentType string, size int64) error {
	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err := s3.New(a.session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(a.cfg.AWS.Bucket),
		Key:           aws.String(fileName),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(file),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(contentType),
	})

	return err
}
