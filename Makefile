.PHONY: build       # Build the Docker image
build:
	docker-compose -f deploy/docker-compose.yml build --no-cache

.PHONY: up          # Starts the Docker containers
up:
	docker-compose -f deploy/docker-compose.yml up -d

.PHONY: down        # Stops the Docker containers
down:
	docker-compose -f deploy/docker-compose.yml down

.PHONY: fmt         # Go format the files
fmt:
	docker-compose -f deploy/docker-compose.yml exec golang gofmt -w ./..

.PHONY: test        # Runs the tests inside the container
test:
	docker-compose -f deploy/docker-compose.yml exec golang go test -cover ./... -coverprofile=coverage.out

.PHONY: ssh         # Access to bash and tty
ssh:
	docker-compose -f deploy/docker-compose.yml exec golang sh

.PHONY: lint        # Runs the lints with few options enabled
lint:
	docker-compose -f deploy/docker-compose.yml exec golang golangci-lint run \
	    --enable=golint \
	    --enable=gofmt \
	    --enable=gosec  \
	    --enable=misspell \
	    --enable=maligned \
	    --enable=interfacer \
	    --enable=stylecheck  \
	    --enable=unconvert \
	    --enable=maligned \
	    --enable=bodyclose \
	    --enable=goconst \
	    --enable=nakedret \
	    --enable=unparam \
	    --enable=lll \
	    --enable=gochecknoinits \
	    --enable=dupl \
	    --enable=depguard \
	    --enable=gocritic \
	    --enable=prealloc \
	    --enable=gocyclo \
	    --enable=funlen