package main

import (
	"aws-eks/helpers"
	"aws-eks/stacks"
	"log"
	"os"
	"reflect"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)
	envName := app.Node().GetContext(jsii.String("environment"))
	// envName := cdktf.NewTerraformVariable(app, jsii.String("environment"), &cdktf.TerraformVariableConfig{})
	yamlFile, err := os.ReadFile("./config/" + reflect.ValueOf(envName).String() + ".yaml")
	if err != nil {
		log.Fatal(err)
	}

	stacks.NewNetworkStack(app, "VPC", helpers.Config(yamlFile))
	// stacks.NewEksStack(app, "EksClusters", helpers.Config(yamlFile))

	app.Synth()
}
