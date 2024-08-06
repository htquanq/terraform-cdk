package helpers

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func SubnetType(subnetType string) awsec2.SubnetType {
	switch subnetType {
	case "public":
		{
			return awsec2.SubnetType_PUBLIC
		}
	case "private_nat":
		{
			return awsec2.SubnetType_PRIVATE_WITH_NAT
		}
	case "private_egress":
		{
			return awsec2.SubnetType_PRIVATE_WITH_EGRESS
		}
	case "private_ioslated":
		{
			return awsec2.SubnetType_PRIVATE_ISOLATED
		}
	default:
		return awsec2.SubnetType_PUBLIC
	}
}

func CreateInternetGateway(CreateInternetGateway bool) bool {
	return CreateInternetGateway
}

func MapIPPublicOnLaunch(MapIPPublicOnLaunch bool) bool {
	return MapIPPublicOnLaunch
}

func InstanceTenancyVerifier(InstanceTenancy string) *string {
	switch InstanceTenancy {
	case strings.ToLower("default"):
		{
			return jsii.String("default")
		}
	case strings.ToLower("dedicated"):
		{
			return jsii.String("dedicated")
		}
	default:
		{
			return jsii.String("default")
		}
	}
}
