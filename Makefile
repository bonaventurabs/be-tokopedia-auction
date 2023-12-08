# Database
PSQL_USER ?= user
PSQL_PASSWORD ?= 12345
PSQL_ADDRESS ?= postgres:5432
PSQL_DATABASE ?= tokopedia-auction

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# ~~~ Tools ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ~~ [migrate] ~~~ https://github.com/golang-migrate/migrate ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MIGRATE := $(shell command -v migrate || echo "bin/migrate")
migrate: bin/migrate ## Install migrate (database migration)

bin/migrate: VERSION := 4.14.1
bin/migrate: GITHUB  := golang-migrate/migrate
bin/migrate: ARCHIVE := migrate.$(OSTYPE)-amd64.tar.gz
bin/migrate: bin
	@ printf "Install migrate... "
	@ curl -Ls $(call github_url) | tar -zOxf - ./migrate.$(shell echo $(OSTYPE) | tr A-Z a-z)-amd64 > $@ && chmod +x $@
	@ echo "done."

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

up: dev-env         						## Startup / Spinup Docker Compose 
down: docker-stop               ## Stop Docker
destroy: docker-teardown clean  ## Teardown (removes volumes, tmp files, etc...)

install-deps: migrate## Install Development Dependencies (localy).
deps: $(MIGRATE) ## Checks for Global Development Dependencies.
deps:
	@echo "Required Tools Are Available"

dev-env: ## Bootstrap Environment (with a Docker-Compose help).
	@ docker-compose up -d --build 

dev-env-test: dev-env ## Run application (within a Docker-Compose help)
	@ $(MAKE) image-build
	docker-compose up app

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# -trimpath - will remove the filepathes from the reports, good to same money on network trafic,
#             focus on bug reports, and find issues fast.
# - race    - adds a racedetector, in case of racecondition, you can catch report with sentry.
#             https://golang.org/doc/articles/race_detector.html
#
# todo(butuzov): add additional flags to compiler to have an `version` flag.
build: ## Builds binary
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o engine \
		./app/
	@ echo "done"


build-race: ## Builds binary (with -race flag)
	@ printf "Building aplication with race flag... "
	@ go build \
		-trimpath  \
		-race      \
		-o engine \
		./app/
	@ echo "done"

# ~~~ Docker Build ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.ONESHELL:
image-build:
	@ echo "Docker Build"
	@ DOCKER_BUILDKIT=0 docker build \
		--file Dockerfile \
		--tag go-clean-arch \
			.

# ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 

PSQL_DSN := "postgres://$(PSQL_USER):$(PSQL_PASSWORD)@$(PSQL_ADDRESS)/$(PSQL_DATABASE)"

migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(PSQL_DSN) -path=misc/migrations up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(PSQL_DSN) -path=misc/migrations down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE) ## Drop everything inside the database.
	migrate  -database $(PSQL_DSN) -path=misc/migrations drop

.PHONY: migrate-create
migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir misc/migrations $${Name}

# ~~~ Cleans ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

clean: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@printf "Cleanning artifacts... "
	@rm -f *.out
	@echo "done."


clean-docker: ## Removes dangling docker images
	@ docker image prune -f


# creates a directory bin.
bin:
	@ mkdir -p $@