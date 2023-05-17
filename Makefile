APP_NAME=modelo-graphql-go
VERSION=$(shell cat VERSION)
FULL_IMAGE=acrdataeng.azurecr.io/estruturante/${APP_NAME}
COMMIT_SHA=$(shell git log -1 --pretty=format:"%H")


default: clean test build

.PHONY: deps
deps:
	go mod vendor && go mod download

.PHONY: clean
clean:
	@rm -rf target

.PHONY: prepare
prepare:
	@chmod +x scripts/* && scripts/prepare.sh

.PHONY: build
build: prepare
	scripts/build-binary.sh ${VERSION} ${COMMIT_SHA}

.PHONY: build-image
build-image:
	scripts/build-image.sh ${FULL_IMAGE} ${VERSION} ${OS_ARCH} ${SUFFIX_TAG}

.PHONY: test
test: prepare
	@echo "tags $(TAGS)"
	go test $(TAGS) -mod=vendor -outputdir=target/tests -coverprofile=coverage.out -v ./...
	go tool cover -func target/tests/coverage.out



.PHONY: generate
generate: api-doc


.PHONY: qa
qa: prepare vet test
	@find ./ -type f -name "*.go" | grep -v vendor | xargs gofmt -d | tee target/format.diff
	@test ! -s target/format.diff || { echo "ERROR: the source code has not been formatted - please use 'make format'"; exit 1; }

.PHONY: format
format:
	@find ./ -type f -name "*.go" | grep -v vendor | xargs gofmt -w

.PHONY: vet
vet:
	go vet ./...


.PHONY: api-doc
api-doc:
	swag init -g cmd/${APP_NAME}/main.go  -o api/


.PHONY: update-packages
update-packages:
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod tidy


.PHONY: generate
generate:
	go generate ./...


.PHONY: config-private-repo
config-private-repo: prepare
	scripts/configure_private_repo.sh