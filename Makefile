GO           ?= go
GOFMT        ?= $(GO)fmt
pkgs          = ./...
BINARY_NAME=beat
BINARY_LINUX=releases/$(BINARY_NAME)_linux_amd64
BINARY_MAC=releases/$(BINARY_NAME)_darwin_amd64


help: Makefile
	@echo
	@echo " Choose a command run in Beat:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


## install_revive: Install revive for linting.
install_revive:
	@echo ">> ============= Install Revive ============= <<"
	$(GO) get github.com/mgechev/revive


## install_golint: Install golint for linting.
install_golint:
	@echo ">> ============= Install Revive ============= <<"
	$(GO) get -u golang.org/x/lint/golint


## style: Check code style.
style:
	@echo ">> ============= Checking Code Style ============= <<"
	@fmtRes=$$($(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print)); \
	if [ -n "$${fmtRes}" ]; then \
		echo "gofmt checking failed!"; echo "$${fmtRes}"; echo; \
		echo "Please ensure you are using $$($(GO) version) for formatting code."; \
		exit 1; \
	fi


## check_license: Check if license header on all files.
check_license:
	@echo ">> ============= Checking License Header ============= <<"
	@licRes=$$(for file in $$(find . -type f -iname '*.go' ! -path './vendor/*') ; do \
               awk 'NR<=3' $$file | grep -Eq "(Copyright|generated|GENERATED)" || echo $$file; \
       done); \
       if [ -n "$${licRes}" ]; then \
               echo "license header checking failed:"; echo "$${licRes}"; \
               exit 1; \
       fi


## test_short: Run test cases with short flag.
test_short:
	@echo ">> ============= Running Short Tests ============= <<"
	$(GO) test -short $(pkgs)


## test: Run test cases.
test:
	@echo ">> ============= Running All Tests ============= <<"
	$(GO) test -v -cover $(pkgs)


## lint: Lint the code.
lint:
	@echo ">> ============= Lint All Files ============= <<"
	revive -config config.toml -exclude vendor/... -formatter friendly ./...
	golint ./...


## verify: Verify dependencies
verify:
	@echo ">> ============= List Dependencies ============= <<"
	$(GO) list -m all
	@echo ">> ============= Verify Dependencies ============= <<"
	$(GO) mod verify


## format: Format the code.
format:
	@echo ">> ============= Formatting Code ============= <<"
	$(GO) fmt $(pkgs)


## vet: Examines source code and reports suspicious constructs.
vet:
	@echo ">> ============= Vetting Code ============= <<"
	$(GO) vet $(pkgs)


## coverage: Create HTML coverage report
coverage:
	@echo ">> ============= Coverage ============= <<"
	rm -f coverage.html cover.out
	$(GO) test -coverprofile=cover.out $(pkgs)
	go tool cover -html=cover.out -o coverage.html


## build: Build go binaries for Linux and Mac
build:
	mkdir releases
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_LINUX) -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -o $(BINARY_MAC) -v


## ci: Run all CI tests.
ci: style check_license test vet lint
	@echo "\n==> All quality checks passed"


.PHONY: help
