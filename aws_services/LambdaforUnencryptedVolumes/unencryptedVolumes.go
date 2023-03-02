package main

/*
	Goal => Create a lambda which looks for all the EBS volumes and share encryption status
		1. List down all the EBS volumes
		2. Share all the status of the EBS Volumes
*/

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
)

type EbsVolume struct {
	VolumeID   string
	Az         string
	Encrypted  bool
	InuseState string
}

// InitialiseAWSCreds - This is to initialise the AWS Session
func InitialiseAWSCreds() (*ec2.EC2, error) {
	awsAccessKey, awsSecretAccessKey := os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY")
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretAccessKey, "")
	config := aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	}
	sess, err := session.NewSession(&config)
	if err != nil {
		fmt.Println(err)
	}
	svc := ec2.New(sess)
	return svc, nil
}

func AssessEncryption(region string, Client *ec2.EC2) {
	volumes := []EbsVolume{}
	//sess, err := InitialiseAWSCreds()
	//if err != nil {
	//	fmt.Println(err)
	//}
	ebsInput := &ec2.DescribeVolumesInput{}
	volumeResult := GetVolumes(ebsInput, Client, volumes)
	fmt.Println(volumeResult)
}

func GetVolumes(ebsInput *ec2.DescribeVolumesInput, svc *ec2.EC2, vol []EbsVolume) []EbsVolume {
	ebsOutput, err := svc.DescribeVolumes(ebsInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ebsOutput)
	return nil
}

func handler() {
	log.Println("Hello")

}

func main() {
	sess, err := InitialiseAWSCreds()
	if err != nil {
		fmt.Println(err)
	}
	region := []string{"us-east-1", "us-west-2"}
	for _, particularRegion := range region {
		AssessEncryption(particularRegion, sess)
	}

	//lambda.Start(handler)
}
