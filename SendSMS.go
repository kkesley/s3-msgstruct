package msgstruct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//SendSMS use reusable s3 session
func SendSMS(sess *session.Session, bucket string, key string, sms StandardSMSStructure) error {
	if len(key) <= 0 {
		return errors.New("key must not be empty")
	}
	svc := s3.New(sess)
	strSMS, err := json.Marshal(sms)
	if err != nil {
		return err
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9_-]+")
	if err != nil {
		fmt.Println(err)
		return err
	}
	filteredKey := reg.ReplaceAllString(key, "_")
	fmt.Println("sending to s3")
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("sms/sending/" + filteredKey + ".json"),
		Body:   bytes.NewReader(strSMS),
	})
	return err
}

//SendSMSDefault sends sms without s3 session
func SendSMSDefault(region string, bucket string, key string, sms StandardSMSStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendSMS(sess, bucket, key, sms)
}
