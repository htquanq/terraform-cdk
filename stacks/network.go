package stacks

import (
	"aws-eks/structs"
	"strconv"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	azs "github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsavailabilityzones"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	awssubnet "github.com/cdktf/cdktf-provider-aws-go/aws/v19/subnet"
	awsvpc "github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewNetworkStack(scope constructs.Construct, id string, AppConfig structs.AppConfig) cdktf.TerraformStack {

	stack := cdktf.NewTerraformStack(scope, &id)
	provider := awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region:  &AppConfig.AWSProfileRegion,
		Profile: &AppConfig.AWSProfileName,
	})
	vpc := awsvpc.NewVpc(stack, &AppConfig.Vpc.Name, &AppConfig.Vpc.Config)

	generateVpcOutput(stack, vpc)

	subnetConfig := awssubnet.SubnetConfig{
		VpcId:     vpc.Id(),
		DependsOn: &[]cdktf.ITerraformDependable{vpc},
	}

	zones := azs.NewDataAwsAvailabilityZones(stack, jsii.String("zones"), &azs.DataAwsAvailabilityZonesConfig{
		Provider: provider,
		State:    jsii.String("available"),
		// DependsOn: &[]cdktf.ITerraformDependable{provider},
	})

	for i := 0; i < len(AppConfig.Vpc.PublicSubnets); i++ {
		tmpSubnetConfig := subnetConfig
		tmpSubnetConfig.CidrBlock = AppConfig.Vpc.PublicSubnets[i].CidrBlock
		tmpSubnetConfig.MapPublicIpOnLaunch = bool(true)
		tmpSubnetConfig.AvailabilityZone = jsii.String(cdktf.Fn_Element(cdktf.Token_AsAny(zones.Names()), jsii.Number(generateAzIndex(i, int(*cdktf.Fn_LengthOf(cdktf.Token_AsAny(zones.Names())))))).(string))
		subnet := createNewSubnet(stack, *jsii.String("public-subnet-" + strconv.Itoa(i)), tmpSubnetConfig)
		cdktf.NewTerraformOutput(stack, jsii.String("public-subnet-"+strconv.Itoa(i)+"-id"), &cdktf.TerraformOutputConfig{
			Value: subnet.Id(),
		})
	}

	for i := 0; i < len(AppConfig.Vpc.PrivateSubnets); i++ {
		tmpSubnetConfig := subnetConfig
		tmpSubnetConfig.CidrBlock = AppConfig.Vpc.PrivateSubnets[i].CidrBlock
		tmpSubnetConfig.AvailabilityZone = jsii.String(cdktf.Fn_Element(cdktf.Token_AsAny(zones.Names()), jsii.Number(generateAzIndex(i, int(*cdktf.Fn_LengthOf(cdktf.Token_AsAny(zones.Names())))))).(string))
		subnet := createNewSubnet(stack, *jsii.String("private-subnet-" + strconv.Itoa(i)), tmpSubnetConfig)
		cdktf.NewTerraformOutput(stack, jsii.String("private-subnet-"+strconv.Itoa(i)+"-id"), &cdktf.TerraformOutputConfig{
			Value: subnet.Id(),
		})
	}

	return stack
}
