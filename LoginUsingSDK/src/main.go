package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var svc *s3.S3
	//using environment variables
	//svc, err := usingEnvironmentVaraiables()

	//using Shared File
	svc, err := usingSharedFile()

	//Hardcoding details in code
	//svc, err := usingHardcodedValue()

	if err != nil {
		fmt.Println("Error in Authentication :", err)
	}
	getS3BucketList(svc)

}
func getS3BucketList(svc *s3.S3) {
	input := &s3.ListBucketsInput{}
	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(result)
}
func usingEnvironmentVaraiables() (*s3.S3, error) {

	//Setting Up environment variable
	os.Setenv("aws_access_key_id", "<AccessKey>")
	os.Setenv("aws_secret_access_key", "<SecretKey>")
	os.Setenv("AWS_REGION", "us-east-1")
	return s3.New(session.New()), nil
}

func usingSharedFile() (*s3.S3, error) {

	//Loading credentials from a file.
	//By default will check in C:\Users\<username>\.aws\credentials File
	//Create a file in editor with name credentials with the below data
	// [default]
	// aws_access_key_id = <accessKey>
	// aws_secret_access_key = <secretKey>
	// [key1]
	// aws_access_key_id = <accessKey>
	// aws_secret_access_key = <secretKey>
	//1st argument >> file location, If default make it as empty
	//2nd argumet >> profile name (Default Prfile name : default). In above default and key1 are profile name

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", ""),
	})
	svc := s3.New(sess)
	return svc, err
}
func usingHardcodedValue() (*s3.S3, error) {

	//this method is not recommended
	//Token not mandatory
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("<AccessKey>", "<SecretKey>", "<Token>"),
	})
	svc := s3.New(sess)
	return svc, err
}
