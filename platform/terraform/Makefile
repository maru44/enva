.PHONY: plan apply

ENVIRON:=prod
# if dev ENV=dev
ENV:=$(ENVIRON) # prod or dev

VAR_FILE:=main.tfvars

init:
	@echo "terraform init"
	@cd ${ENV} && terraform init

plan:
	@echo "terraform plan -var-file ${VAR_FILE}"
	@cd ${ENV} && terraform plan -var-file ${VAR_FILE}

apply:
	@echo "terraform apply -var-file ${VAR_FILE}"
	@cd ${ENV} && terraform apply -var-file ${VAR_FILE}

destroy:
	@echo "terraform destroy -var-file ${VAR_FILE}"
	@cd ${ENV} && terraform destroy -var-file ${VAR_FILE}
