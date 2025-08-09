SHELL := /bin/bash
define go_mod
	@echo "GO MOD TIDY"
	go mod tidy
endef

synth-network:
	shopt -s expand_aliases;
	$(call go_mod)
	sed -i "s/environment_to_replace/demo/g" cdktf.json;
	STACK_NAME="network" cdktf synth;

deploy-network:
	sed -i "s/environment_to_replace/demo/g" cdktf.json;
	STACK_NAME="network" cdktf deploy;

synth-rds:
	shopt -s expand_aliases;
	$(call go_mod)
	sed -i "s/environment_to_replace/demo/g" cdktf.json;
	STACK_NAME="rds" cdktf synth ;

deploy-rds:
	sed -i "s/environment_to_replace/demo/g" cdktf.json;
	STACK_NAME="rds" cdktf deploy;