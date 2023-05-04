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

func AssessEncryption(region string) {
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
	volumes := []EbsVolume{}
	ebsInput := &ec2.DescribeVolumesInput{}
	volumeResult := GetVolumes(ebsInput, svc, volumes)
	countUnencryptedEBS := 0
	volID := []string{}
	for _, ev := range volumeResult {
		if ev.Encrypted == false {
			volID = append(volID, ev.VolumeID)
			countUnencryptedEBS++
		}
	}

	if len(volumeResult) > 0 {
		log.Printf("Total Unencrypted EBS Volumes in Region: %v is %v\nVolume IDs: %v \n", region, countUnencryptedEBS, volID)
	}
}

func GetVolumes(ebsInput *ec2.DescribeVolumesInput, svc *ec2.EC2, vol []EbsVolume) []EbsVolume {
	ebsOutput, err := svc.DescribeVolumes(ebsInput)
	if err != nil {
		fmt.Println(err)
	}
	if len(ebsOutput.Volumes) > 0 {
		for _, v := range ebsOutput.Volumes {
			e := EbsVolume{}
			e.VolumeID = *v.VolumeId
			e.Encrypted = *v.Encrypted
			e.Az = *v.AvailabilityZone
			e.InuseState = *v.State
			vol = append(vol, e)
		}
		if ebsInput.NextToken != nil {
			ebsInput.SetNextToken(*ebsOutput.NextToken)
			GetVolumes(ebsInput, svc, vol)
		}
	}

	return vol
}

func main() {
	reg := []string{"us-east-1"}
	for _, r := range reg {
		AssessEncryption(r)
	}
}
