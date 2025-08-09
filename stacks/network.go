package stacks

import (
	"aws-eks/structs"
	"strconv"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	azs "github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsavailabilityzones"
	dbSubnetGroup "github.com/cdktf/cdktf-provider-aws-go/aws/v19/dbsubnetgroup"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	awssubnet "github.com/cdktf/cdktf-provider-aws-go/aws/v19/subnet"
	awsvpc "github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NetworkStack(scope constructs.Construct, id string, AppConfig structs.AppConfig) cdktf.TerraformStack {
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
	})

	for id, _ := range AppConfig.Vpc.PublicSubnets {
		tmpSubnetConfig := subnetConfig
		tmpSubnetConfig.CidrBlock = AppConfig.Vpc.PublicSubnets[id].CidrBlock
		tmpSubnetConfig.MapPublicIpOnLaunch = bool(true)
		tmpSubnetConfig.AvailabilityZone = jsii.String(cdktf.Fn_Element(zones.Names(), jsii.Number(float64(id))).(string))
		subnet := createNewSubnet(stack, *jsii.String("public-subnet-" + strconv.Itoa(id)), tmpSubnetConfig)
		cdktf.NewTerraformOutput(stack, jsii.String("public-subnet-"+strconv.Itoa(id)+"-id"), &cdktf.TerraformOutputConfig{
			Value: subnet.Id(),
		})
	}

	for id, _ := range AppConfig.Vpc.PrivateSubnets {
		tmpSubnetConfig := subnetConfig
		tmpSubnetConfig.CidrBlock = AppConfig.Vpc.PrivateSubnets[id].CidrBlock
		tmpSubnetConfig.AvailabilityZone = jsii.String(cdktf.Fn_Element(zones.Names(), jsii.Number(float64(id))).(string))
		subnet := createNewSubnet(stack, *jsii.String("private-subnet-" + strconv.Itoa(id)), tmpSubnetConfig)
		cdktf.NewTerraformOutput(stack, jsii.String("private-subnet-"+strconv.Itoa(id)+"-id"), &cdktf.TerraformOutputConfig{
			Value: subnet.Id(),
		})
	}

	dbSubnetIds := []*string{}

	for id, _ := range AppConfig.Vpc.DbSubnets {
		tmpSubnetConfig := subnetConfig
		tmpSubnetConfig.CidrBlock = AppConfig.Vpc.DbSubnets[id].CidrBlock
		tmpSubnetConfig.AvailabilityZone = jsii.String(cdktf.Fn_Element(zones.Names(), jsii.Number(float64(id))).(string))
		subnet := createNewSubnet(stack, *jsii.String("db-subnet-" + strconv.Itoa(id)), tmpSubnetConfig)
		dbSubnetIds = append(dbSubnetIds, subnet.Id())
		cdktf.NewTerraformOutput(stack, jsii.String("db-subnet-"+strconv.Itoa(id)+"-id"), &cdktf.TerraformOutputConfig{
			Value: subnet.Id(),
		})
	}

	DbSubnetGroup(stack, "db-subnet", &dbSubnetIds)

	return stack
}

func DbSubnetGroup(scope constructs.Construct, id string, subnetIds *[]*string) cdktf.TerraformOutput {
	subnetGroup := dbSubnetGroup.NewDbSubnetGroup(scope, jsii.String(id), &dbSubnetGroup.DbSubnetGroupConfig{
		Name:      jsii.String(id),
		SubnetIds: subnetIds,
	})

	return cdktf.NewTerraformOutput(scope, jsii.String(id+"-db-subnets"), &cdktf.TerraformOutputConfig{
		Value: subnetGroup.Arn(),
	})
}
