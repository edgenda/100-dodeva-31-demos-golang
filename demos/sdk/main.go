package main

import (
	"fmt"
	"log"
	"os"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// The Lshortfile flag includes file name and line number in log messages.
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Printf("starting demo 1\n")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	if err != nil {
		exitErrorf("Unable to create session, %v", err)
	}

	sess.Handlers.Send.PushFront(func(r *request.Request) {
		rParamsBytes, err := json.Marshal(r.Params)
		if err != nil {}
		// Log every request made and its payload
		log.Printf("Request: %s/%s, Params: %s",
			r.ClientInfo.ServiceName, r.Operation.Name, string(rParamsBytes))
	})
	_, err = sess.Config.Credentials.Get()
	if err != nil {
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
