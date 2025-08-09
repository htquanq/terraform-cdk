package structs

import (
	awsekscluster "github.com/cdktf/cdktf-provider-aws-go/aws/v19/ekscluster"
	awsiamrole "github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamrole"
	awssubnet "github.com/cdktf/cdktf-provider-aws-go/aws/v19/subnet"
	awsvpc "github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpc"
)

type AppConfig struct {
	AWSAccountID     string                           `yaml:"accountId"`
	AWSProfileName   string                           `yaml:"awsProfile"`
	AWSProfileRegion string                           `yaml:"region"`
	Vpc              VpcConfigRoot                    `yaml:"vpc"`
	EksClusters      []awsekscluster.EksClusterConfig `yaml:"eksClusters"`
	IamRoles         []awsiamrole.IamRoleConfig       `yaml:"iamRoles"`
}

type VpcConfigRoot struct {
	Name           string                   `yaml:"name"`
	Config         awsvpc.VpcConfig         `yaml:"config"`
	PublicSubnets  []awssubnet.SubnetConfig `yaml:"publicSubnets"`
	PrivateSubnets []awssubnet.SubnetConfig `yaml:"privateSubnets"`
	DbSubnets      []awssubnet.SubnetConfig `yaml:"dbSubnets"`
}
