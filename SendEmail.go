package msgstruct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//SendEmail use reusable s3 session
func SendEmail(sess *session.Session, bucket string, key string, email StandardEmailStructure) error {
	if len(key) <= 0 {
		return errors.New("key must not be empty")
	}
	svc := s3.New(sess)
	strEmail, err := json.Marshal(email)
	if err != nil {
		return err
	}
	fmt.Println("sending to s3")
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("email/sending/" + key + ".json"),
		Body:   bytes.NewReader(strEmail),
	})
	return err
}

//SendEmailDefault sends email without s3 session
func SendEmailDefault(region string, bucket string, key string, email StandardEmailStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendEmail(sess, bucket, key, email)
}
