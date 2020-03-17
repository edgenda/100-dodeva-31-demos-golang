package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	fmt.Printf("starting demo 1\n")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	
	if err != nil{
		exitErrorf("Unable to create session, %v", err)
	}
	_, err = sess.Config.Credentials.Get()

	if err != nil{
		exitErrorf("Credentials are invalid, %v", err)
	}
	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	
	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
