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
	Goal is to understand the status of ec2-instance ( Running or Stopped )
*/

func handler() {
	log.Println("Hello Lambda")
}

func GetEc2Instance(client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return result, err
}

func main() {
	awsAccesskey := "AKIA6MJ4CNSY2PQZPC5Q"
	awsSecretKey := "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"
	creds := credentials.NewStaticCredentials(awsAccesskey, awsSecretKey, "")
	config := aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	}
	sess, err := session.NewSession(&config)
	if err != nil {
		fmt.Println(err)
	}
	ec2Client := ec2.New(sess)
	runningInstances, err := GetEc2Instance(ec2Client)
	if err != nil {
		fmt.Printf("Couldn't retrieve running instances: %v", err)
		return
	}

	for _, reservation := range runningInstances.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Found running instance: %s\n", *instance.InstanceId)
		}
	}

	//lambda.Start(handler)
}
