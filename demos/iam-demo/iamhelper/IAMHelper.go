package iamhelper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"

	"regexp"
)

// IAMHelper ; AWS IAM service helper
type IAMHelper struct {
	session *session.Session
	iamsvc  *iam.IAM
}

// NewIAMHelper : Instantiates an IAMHelper
func NewIAMHelper(session *session.Session) *IAMHelper {
	iamSvc := iam.New(session)
	return &IAMHelper{
		iamsvc:  iamSvc,
		session: session,
	}
}

// GetUser : Returns session associated user
func(helper *IAMHelper) GetUser() (*iam.GetUserOutput, error){
	input := &iam.GetUserInput{
	}
	result, err := helper.iamsvc.GetUser(input)
	return result, err
}

// GetAccountNumber : Returns current user account number
func(helper *IAMHelper) GetAccountNumber() (*string, error) {
	user, err := helper.GetUser()
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	re := regexp.MustCompile(`\d{10,16}`)
	found := re.FindAllString(*user.User.Arn, 1)[0]
	return &found, nil
}