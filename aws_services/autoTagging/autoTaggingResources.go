package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Instance struct {
	InstanceID string    `json:"instance_id"`
	Tags       []TagType `json:"Tags"`
}
type TagType struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Instances struct {
	Instances []*Instance `json:"instances"`
}

var (
	ec2client *ec2.Client
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error," + err.Error())

	}

	ec2client = ec2.NewFromConfig(cfg)
}

// GettingTag - The Function Queries all the ec2 instances and ebs volumes about Tags
func GettingTag() {
	var ins []*Instance
	params := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}
	result, err := ec2client.DescribeInstances(context.TODO(), params)
	if err != nil {
		fmt.Println("Error Calling ec2:", err)
		return
	}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			tags := []TagType{}
			for _, tag := range instance.Tags {
				tagstype := TagType{Key: *tag.Key, Value: *tag.Value}
				tags = append(tags, tagstype)
			}
			ec2Tag := &Instance{
				InstanceID: *instance.InstanceId,
				Tags:       tags}
			ins = append(ins, ec2Tag)
		}
	}
	p, _ := json.MarshalIndent(Instances{Instances: ins}, " ", "\t")
	fmt.Printf("%s\n", p)
}

//TODO - Need to store the tags and later compare if needed more tags

// TagNotProper - This function checks whether tags which are not proper matching SSM
func TagNotProper() {

}

// ApplyTag - Applies Tag to those Resource which doesn't match proper tagging and apply as per SSM
func ApplyTag() {

}

func main() {
	GettingTag()

}
