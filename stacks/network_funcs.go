package stacks

import (
	"github.com/aws/jsii-runtime-go"
	awssubnet "github.com/cdktf/cdktf-provider-aws-go/aws/v19/subnet"
	awsvpc "github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func createNewSubnet(stack cdktf.TerraformStack, id string, subnetConfig awssubnet.SubnetConfig) awssubnet.Subnet {
	subnet := awssubnet.NewSubnet(stack, &id, &subnetConfig)
	generateSubnetOutput(stack, id, subnet)
	return subnet
}

func generateSubnetOutput(stack cdktf.TerraformStack, id string, subnet awssubnet.Subnet) {
	cdktf.NewTerraformOutput(stack, jsii.String(id+"Id"), &cdktf.TerraformOutputConfig{
		Value: subnet.Id(),
	})
}
func generateVpcOutput(stack cdktf.TerraformStack, vpc awsvpc.Vpc) {
	cdktf.NewTerraformOutput(stack, jsii.String("VpcArn"), &cdktf.TerraformOutputConfig{
		Value: vpc.Arn(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("VpcId"), &cdktf.TerraformOutputConfig{
		Value: vpc.Id(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("FrienlyId"), &cdktf.TerraformOutputConfig{
		Value: vpc.FriendlyUniqueId(),
	})
}
