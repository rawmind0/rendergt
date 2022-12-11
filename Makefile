GOFMT_FILES?=$$(find . -name '*.go')
BIN_NAME="rendergt"

default: build

build: validate
	@sh -c "'$(CURDIR)/scripts/build'"

ci: validate
	@sh -c "'$(CURDIR)/scripts/ci'"

validate: fmtcheck lint vet

vet:
	@echo "==> Checking that code complies with go vet requirements..."
	@go vet $$(go list ./...); if [ $$? -gt 0 ]; then \
		echo ""; \
		echo "If vet reported more suspicious constructs, please check and"; \
		echo "fix them if necessary, before submitting the code for review."; \
	fi

lint:
	@echo "==> Checking that code complies with golint requirements..."
	@GO111MODULE=off go get -u golang.org/x/lint/golint
	@if [ -n "$$(golint $$(go list ./...) | grep -v 'should have comment.*or be unexported' | tee /dev/stderr)" ]; then \
		echo ""; \
		echo "golint found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review."; \
    	exit 1; \
	fi

bin:
	go build -tags urfave_cli_no_docs -o $(BIN_NAME)

mod:
	@echo "==> Updating go modules..."
	go mod tidy

fmt:
	@echo "==> Updating code to complies with gofmt requirements..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/fmtcheck'"

package:
	@sh -c "'$(CURDIR)/scripts/package'"

.PHONY: build bin fmt vet lint fmtcheck package ci
