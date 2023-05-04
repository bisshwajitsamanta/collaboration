package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"os"
)

/*
	Goal =>
		1. List out the Load balancers in the region => Done
		2. List out the instances under the load balancers Target Group.
		3. Restart the instances under the load balancers one by one.
*/

type LoadBalancer struct {
	Name string
}

// InitialiseAWS - Initialise AWS Session
func InitialiseAWS() *elbv2.ELBV2 {
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
	svc := elbv2.New(sess)
	return svc
}

// RetrieveInstances - Retrieve Instance Details
func RetrieveInstances(svc *elbv2.ELBV2) {
	targetGroupArn := TargetGroups(svc)
	fmt.Println(targetGroupArn)
}

// TargetGroups - List out the Targets Groups
func TargetGroups(svc *elbv2.ELBV2) string {
	loadbalancerarn := ListLoadBalancer(svc)
	var targetGroupArn string
	input := elbv2.DescribeTargetGroupsInput{
		LoadBalancerArn: aws.String(loadbalancerarn),
	}
	result, err := svc.DescribeTargetGroups(&input)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result.TargetGroups {
		targetGroupArn = *v.TargetGroupArn
	}
	return targetGroupArn
}

// ListLoadBalancer - List Load Balancers
func ListLoadBalancer(svc *elbv2.ELBV2) string {
	var loadBalancerArn string
	input := elbv2.DescribeLoadBalancersInput{
		Names: []*string{
			aws.String("my-alb"),
		},
	}
	result, err := svc.DescribeLoadBalancers(&input)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result.LoadBalancers {
		loadBalancerArn = *v.LoadBalancerArn
		fmt.Println("Load Balancer Name:", *v.LoadBalancerName)
		fmt.Println("Load Balancer ARN:", *v.LoadBalancerArn)
		fmt.Println("Load Balancer IP:", *v.IpAddressType)
		fmt.Println("Load Balancer Availability Zone:", v.AvailabilityZones)
	}
	return loadBalancerArn
}

func main() {
	svc := InitialiseAWS()
	RetrieveInstances(svc)
}
