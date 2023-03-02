package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

/*
	Goal => Remove UnEncrypted EBS Volume from ec2 Instance and attach Encrypted EBS Volume
		1. Query EBS Volume against the ec2 instance => Done
		2. Find the status of the EBS Volume ( Encrypted or Non Encrypted )
		3. Create Snapshot of the EBS Volume if not encrypted
		4. Remove the EBS Non Encrypted volume from the ec2 instance
		5. Attach the New EBS volume Encrypted volume to the ec2 instance
*/

// GetEc2EBSVolume -
func GetEc2EBSVolume(ec2Client *ec2.EC2) {
	ec2Client = InitialiseAWS()
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String("i-0c098929127d8b28f"),
		},
	}
	result, err := ec2Client.DescribeInstances(input)
	if err != nil {
		fmt.Println(err)
	}
	for _, ec2Output := range result.Reservations {
		for _, instances := range ec2Output.Instances {
			for _, ebsBlock := range instances.BlockDeviceMappings {
				log.Println("EBS Volume ID:", *ebsBlock.Ebs.VolumeId)
			}
		}
	}
}

// QueryEBSEncryptionStatus -
func QueryEBSEncryptionStatus() {

}

func InitialiseAWS() *ec2.EC2 {
	awsAccessKey := "AKIA4NHZ4F7V35ARB7NX"
	awsSecretAccessKey := "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"
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretAccessKey, "")
	config := aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	}
	sess, err := session.NewSession(&config)
	if err != nil {
		fmt.Println(err)
	}
	ec2Client := ec2.New(sess)

	return ec2Client
}

func main() {
	session := InitialiseAWS()
	GetEc2EBSVolume(session)

}
