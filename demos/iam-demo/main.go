package main

import (
	"flag"
	"fmt"
	"log"

	"encoding/json"

	iamHelper "github.com/edgenda/100-dodeva-31-demos-golang/iam-demo/iamhelper"
	s3Helper "github.com/edgenda/100-dodeva-31-demos-golang/iam-demo/s3helper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
)

// The Lshortfile flag includes file name and line number in log messages.
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func exitErrorf(msg string, args ...interface{}) {
	log.Fatalf(msg+"\n", args)
}

func main() {
	bucketPrefixPtr := flag.String("prefix", "default-bucket", "bucket prefix name to create")
	flag.Parse()

	log.Printf("starting demo 2\n")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	if err != nil {
		exitErrorf("Unable to create session, %v", err)
	}

	sess.Handlers.Send.PushFront(func(r *request.Request) {
		rParamsBytes, err := json.Marshal(r.Params)
		if err != nil {
		}
		// Log every request made and its payload
		log.Printf("Request: %s/%s, Params: %s",
			r.ClientInfo.ServiceName, r.Operation.Name, string(rParamsBytes))
	})
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		exitErrorf("Credentials are invalid, %v", err)
	}
	s3Helper := s3Helper.NewS3Helper(sess)
	iamHelper := iamHelper.NewIAMHelper(sess)
	arn, err := iamHelper.GetAccountNumber()
	if err != nil {
		exitErrorf("Can't read account number, %v", err)
	}
	bucketName := fmt.Sprintf("%s-%s-ca-central-1", *bucketPrefixPtr, *arn)
	result, err := s3Helper.CreateBucket(&bucketName)

	resultBytes, err := json.Marshal(result)
	if err != nil {
		exitErrorf("Can't marshal create bucket output, %v", err)
	}
	fmt.Printf("%s", string(resultBytes))
}
