package msgstruct

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//SendEmail use reusable s3 session
func SendEmail(sess *session.Session, bucket string, key string, email StandardEmailStructure) error {
	return SendBatchEmail(sess, bucket, key, []StandardEmailStructure{email})
}

//SendEmailDefault sends email without s3 session
func SendEmailDefault(region string, bucket string, key string, email StandardEmailStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendEmail(sess, bucket, key, email)
}
