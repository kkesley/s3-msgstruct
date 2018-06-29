package msgstruct

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//SendSMS use reusable s3 session
func SendSMS(sess *session.Session, bucket string, key *string, sms StandardSMSStructure) error {
	svc := s3.New(sess)
	strSMS, err := json.Marshal(sms)
	if err != nil {
		return err
	}
	fmt.Println("sending to s3")
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    key,
		Body:   bytes.NewReader(strSMS),
	})
	return err
}

//SendSMSDefault sends sms without s3 session
func SendSMSDefault(region string, bucket string, key *string, sms StandardSMSStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendSMS(sess, bucket, key, sms)
}
