package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var ec2client *ec2.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error," + err.Error())

	}

	ec2client = ec2.NewFromConfig(cfg)
}

// GettingTag - The Function Queries all the ec2 instances and ebs volumes about Tags
func GettingTag() {
	params := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}
	result, err := ec2client.DescribeInstances(context.TODO(), params)
	if err != nil {
		fmt.Println("Error Calling ec2:", err)
		return
	}
	count := len(result.Reservations)
	fmt.Println("Instances:", count)

	for i, reservation := range result.Reservations {
		for k, instance := range reservation.Instances {
			fmt.Println("Instance number: ", i, "-", k, "Id: ", *instance.InstanceId)
			for _, tag := range instance.Tags {
				fmt.Println("Instance Tags key:", *tag.Key)
				fmt.Println("Instance Tags Value:", *tag.Value)
			}
		}
	}
}

// TagNotProper - This function checks whether tags which are not proper matching SSM
func TagNotProper() {

}

// ApplyTag - Applies Tag to those Resource which doesn't match proper tagging and apply as per SSM
func ApplyTag() {

}

func main() {
	GettingTag()

}
