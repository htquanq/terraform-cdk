SHELL := /bin/bash

synth-demo:
	shopt -s expand_aliases;
	go mod tidy;
	gsed -i "s/environment_to_replace/demo/g" cdktf.json;
	cdktf synth;

deploy-demo:
	go mod tidy;
	gsed -i "s/environment_to_replace/demo/g" cdktf.json;
	cdktf deploy;