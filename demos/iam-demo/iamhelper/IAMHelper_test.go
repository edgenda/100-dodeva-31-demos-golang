package iamhelper

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// The Lshortfile flag includes file name and line number in log messages.
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

// TestsScaffold : Test target scaffold structure
type TestsScaffold struct {
	session *session.Session
	target  *IAMHelper
}

func NewTestsScaffold(session *session.Session) *TestsScaffold {
	testTarget := NewIAMHelper(session)
	return &TestsScaffold{
		session: session,
		target:  testTarget,
	}
}

// TestCanGetUser : Verifies that helper is able to get user
func TestCanGetUser(t *testing.T) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	if err != nil {
		t.Errorf("Unable to create session, %v", err)
	}
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		t.Errorf("Credentials are invalid, %v", err)
	}
	testScaffold := NewTestsScaffold(sess)
	result, err := testScaffold.target.GetUser()
	if err != nil {
		t.Errorf("Failed at getting users, %v", err)
	}
	if result == nil {
		t.Error("user fetched cannot be nil")
	}
	resultBytes, err := json.Marshal(result)
	log.Printf("result: %s", string(resultBytes))
}

// TestCanGetAccountNumber : Test verifies that ARN can be read from user
func TestCanGetAccountNumber(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	if err != nil {
		t.Errorf("Unable to create session, %v", err)
	}
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		t.Errorf("Credentials are invalid, %v", err)
	}
	testScaffold := NewTestsScaffold(sess)
	result, err := testScaffold.target.GetUser()
	if err != nil {
		t.Errorf("Failed at getting users, %v", err)
	}
	if result == nil {
		t.Error("user fetched cannot be nil")
	}
	// can get {{acount_arn}} out of arn:aws:iam::{{acount_arn}}:user/ebrouillard
	arn, err := testScaffold.target.GetAccountNumber()
	if result == nil {
		t.Errorf("can't fetch arn, %v", err)
	}
	if len(*arn) != 12 {
		t.Errorf("arn found doesn't match expectation %s", *arn)
	}
}
