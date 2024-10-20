# function to check if argument is empty
error_if_empty = $(if $(strip $1),,echo "error: required: $2 is empty" && exit 1)

ifeq ($(strip $(EC2_CERT)),)
	EC2_CERT_OPT =
else
	EC2_CERT_OPT = -i ${EC2_CERT}
endif

BUILD_DIR = .deploy
BINARY_NAME = api
OUT_BINARY_PATH = ${BUILD_DIR}/${BINARY_NAME}

EC2_USER ?= ubuntu
EC2_HOST ?= ec2-23-20-111-38.compute-1.amazonaws.com
APPS_HOME = /home/${EC2_USER}/apps/
REMOTE_DEPLOY_PATH = ${APPS_HOME}${BINARY_NAME}/

dev:
	docker compose up --build

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

schema:
	$(call error_if_empty, ${name}, "name")
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}

migrate:
	make generate

	$(call error_if_empty, ${name}, "name")

	atlas migrate diff ${name} \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/16/dev?search_path=public"

ssh:
	ssh ubuntu@ec2-23-20-111-38.compute-1.amazonaws.com

build:
	GOOS=linux GOARCH=amd64 go build -o ${OUT_BINARY_PATH} ./cmd/api

deploy:
	make build
	@if [ -f .deploy/.env ]; then \
		scp ${EC2_CERT_OPT} .deploy/.env ${EC2_USER}@${EC2_HOST}:${REMOTE_DEPLOY_PATH}; \
	fi
	scp ${EC2_CERT_OPT} ${OUT_BINARY_PATH} ${EC2_USER}@${EC2_HOST}:${REMOTE_DEPLOY_PATH}
	scp ${EC2_CERT_OPT} shell/ec2-deploy.sh ${EC2_USER}@${EC2_HOST}:${APPS_HOME}
	ssh ${EC2_CERT_OPT} ${EC2_USER}@${EC2_HOST} "bash ${APPS_HOME}ec2-deploy.sh"
