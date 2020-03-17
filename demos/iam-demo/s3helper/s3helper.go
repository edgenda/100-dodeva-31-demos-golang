package s3helper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Helper ; AWS S3 service helper
type S3Helper struct {
	session *session.Session
	s3svc *s3.S3
}

// NewS3Helper : Instantiates an S3Helper
func NewS3Helper(session *session.Session) *S3Helper {
	svc := s3.New(session)
	return &S3Helper{
		session: session,
		s3svc: svc,
	}
}

// CreateBucket : Creates a bucket with name
func (s3helper *S3Helper) CreateBucket(bucket *string) (*s3.CreateBucketOutput, error) {
	createBucketInput := &s3.CreateBucketInput{
		Bucket: aws.String(*bucket),
	}
	result, err := s3helper.s3svc.CreateBucket(createBucketInput)
	return result, err
}
