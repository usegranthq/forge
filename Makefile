# function to check if argument is empty
error_if_empty = $(if $(strip $1),,echo "error: required: $2 is empty" && exit 1)

dev:
	air -c .air.toml

build:
	go build -o bin/main ./cmd/api

generate:
	go generate ./ent

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
