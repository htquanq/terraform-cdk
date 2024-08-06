package helpers

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

func ImageMapper(image string) awseks.NodegroupAmiType {
	switch image {
	case strings.ToUpper("WINDOWS_FULL"):
		{
			return awseks.NodegroupAmiType_WINDOWS_FULL_2022_X86_64
		}
	case strings.ToUpper("WINDOWS_CORE"):
		{
			return awseks.NodegroupAmiType_WINDOWS_CORE_2022_X86_64
		}
	case strings.ToUpper("GPU"):
		{
			return awseks.NodegroupAmiType_AL2_X86_64_GPU
		}
	case strings.ToUpper("LINUX_ARM"):
		{
			return awseks.NodegroupAmiType_AL2023_ARM_64_STANDARD
		}
	default:
		return awseks.NodegroupAmiType_AL2023_X86_64_STANDARD
	}
}
