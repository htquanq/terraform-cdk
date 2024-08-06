package helpers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

func EksVersionMapping(version float64) awseks.KubernetesVersion {
	switch version {
	case 1.14:
		{
			return awseks.KubernetesVersion_V1_14()
		}
	case 1.15:
		{
			return awseks.KubernetesVersion_V1_15()
		}
	case 1.17:
		{
			return awseks.KubernetesVersion_V1_17()
		}
	case 1.18:
		{
			return awseks.KubernetesVersion_V1_18()
		}
	case 1.19:
		{
			return awseks.KubernetesVersion_V1_19()
		}
	case 1.20:
		{
			return awseks.KubernetesVersion_V1_20()
		}
	case 1.22:
		{
			return awseks.KubernetesVersion_V1_22()
		}
	case 1.23:
		{
			return awseks.KubernetesVersion_V1_23()
		}
	case 1.24:
		{
			return awseks.KubernetesVersion_V1_24()
		}
	case 1.25:
		{
			return awseks.KubernetesVersion_V1_25()
		}
	case 1.26:
		{
			return awseks.KubernetesVersion_V1_26()
		}
	case 1.27:
		{
			return awseks.KubernetesVersion_V1_27()
		}
	case 1.28:
		{
			return awseks.KubernetesVersion_V1_28()
		}
	case 1.29:
		{
			return awseks.KubernetesVersion_V1_29()
		}
	default:
		return awseks.KubernetesVersion_V1_30()
	}
}
