package stacks

// import (
// 	"aws-eks/structs"

// 	"github.com/aws/constructs-go/constructs/v10"
// 	"github.com/aws/jsii-runtime-go"
// 	awsekscluster "github.com/cdktf/cdktf-provider-aws-go/aws/v19/ekscluster"
// 	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
// 	"github.com/hashicorp/terraform-cdk-go/cdktf"
// )

// func NewEksStack(scope constructs.Construct, id string, AppConfig structs.AppConfig) cdktf.TerraformStack {

// 	stack := cdktf.NewTerraformStack(scope, &id)
// 	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
// 		Region:  &AppConfig.AWSProfileRegion,
// 		Profile: &AppConfig.AWSProfileName,
// 	})

// 	for i := 0; i < len(AppConfig.EksClusters); i++ {
// 		clusterConfig := AppConfig.EksClusters[i]
// 		clusterConfig.VpcConfig.SubnetIds = cdktf.NewDataTerraformRemoteState(stack, jsii())
// 		cluster := awsekscluster.NewEksCluster(stack, clusterConfig.Name, &clusterConfig)
// 		generateEksOutput(stack, cluster)
// 	}
// 	return stack
// }

// func generateEksOutput(stack cdktf.TerraformStack, cluster awsekscluster.EksCluster) {
// 	cdktf.NewTerraformOutput(stack, jsii.String("ClusterArn"), &cdktf.TerraformOutputConfig{
// 		Value: cluster.Arn(),
// 	})
// 	cdktf.NewTerraformOutput(stack, jsii.String("ClusterId"), &cdktf.TerraformOutputConfig{
// 		Value: cluster.Id(),
// 	})
// 	cdktf.NewTerraformOutput(stack, jsii.String("ClusterFriendlyId"), &cdktf.TerraformOutputConfig{
// 		Value: cluster.FriendlyUniqueId(),
// 	})
// }
