package stacks

import (
	"aws-eks/structs"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	awsiamrole "github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamrole"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewIamStack(scope constructs.Construct, id string, AppConfig structs.AppConfig) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region:  &AppConfig.AWSProfileRegion,
		Profile: &AppConfig.AWSProfileName,
	})

	for i := 0; i < len(AppConfig.IamRoles); i++ {
		roleConfig := AppConfig.IamRoles[i]
		role := awsiamrole.NewIamRole(stack, roleConfig.Name, &roleConfig)
		generateIamOutput(stack, role)
	}

	return stack
}

func generateIamOutput(stack cdktf.TerraformStack, iamRole awsiamrole.IamRole) {
	cdktf.NewTerraformOutput(stack, jsii.String("VpcArn"), &cdktf.TerraformOutputConfig{
		Value: iamRole.Arn(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("VpcId"), &cdktf.TerraformOutputConfig{
		Value: iamRole.Id(),
	})
}
