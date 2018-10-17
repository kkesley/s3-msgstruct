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

//SendBatchEmail use reusable s3 session
func SendBatchEmail(sess *session.Session, bucket string, key string, emails []StandardEmailStructure) error {
	if len(key) <= 0 {
		return errors.New("key must not be empty")
	}

	svc := s3.New(sess)
	strEmail, err := json.Marshal(emails)
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
		Key:    aws.String("email/sending/" + filteredKey + ".json"),
		Body:   bytes.NewReader(strEmail),
	})
	return err
}

//SendBatchEmailDefault sends batch emails without s3 session
func SendBatchEmailDefault(region string, bucket string, key string, emails []StandardEmailStructure) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return SendBatchEmail(sess, bucket, key, emails)
}
