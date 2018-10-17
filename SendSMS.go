package msgstruct

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//SendSMS use reusable s3 session
func SendSMS(sess *session.Session, bucket string, key string, sms StandardSMSStructure) error {
	return SendBatchSMS(sess, bucket, key, []StandardSMSStructure{sms})
}

//SendSMSDefault sends sms without s3 session
func SendSMSDefault(region string, bucket string, key string, sms StandardSMSStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendSMS(sess, bucket, key, sms)
}
