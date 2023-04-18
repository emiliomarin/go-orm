mod: export GO111MODULE=on
mod: export GOPROXY=direct
mod: export GOSUMDB=off

migration-run: export POSTGRESQL_URL=postgresql://arexdb_dev:arexdb_dev@localhost:5432/go_orm?sslmode=disable

## Install project dependencies
.PHONY: mod
mod:
	@go mod tidy
	@go mod vendor

## Run project
run: ; $(info running code…)
	go run  .

.PHONY: migration-create
migration-create: ## Creates a new migration usage: `migration-create name=<migration name>`
	@migrate create -dir ./migrations -ext sql $(name)

.PHONY: migration-run
migration-run: ## Running migrations: `migration-run `
	$(info Running migrations...)
	@migrate -database ${POSTGRESQL_URL} -path ./migrations $(dir) $(count)
