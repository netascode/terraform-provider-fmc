HOSTNAME=registry.terraform.io
NAMESPACE=netascode
NAME=fmc
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64
PROVIDER_DIR := ./internal/provider/.

temp = $(subst /, ,$@)
OS = $(word 1, $(temp))
ARCH = $(word 2, $(temp))

default: build


build:
	go build -o ${BINARY}_${VERSION}_${OS_ARCH}


install: build
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	rm -rf vendor
	go mod vendor
	go mod tidy
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY}_${VERSION}_${OS_ARCH} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/terraform-provider-fmc
	@echo "Install Successfully Completed"


# Generate new provider
gen:
	go generate


# Run all acceptance tests
# Example: make testacc
.PHONY: testacc
testacc: gen
	TF_ACC=1 go test -v ${PROVIDER_DIR} -timeout 120m


# Run acceptance tests with specific test name
# Example: make testspec ARGS="NetworkHost"
.PHONY: testspec
testspec: gen
	TF_ACC=1 go test ${PROVIDER_DIR} -v -run ^TestAccDataSourceFmc$(ARGS)$$ -timeout 120m -count=1
	TF_ACC=1 go test ${PROVIDER_DIR} -v -run ^TestAccFmc$(ARGS)$$ -timeout 120m -count=1
