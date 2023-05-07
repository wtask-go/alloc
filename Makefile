description := "alloc - allocation \& dereference helper package for Go"

.PHONY: all
all: about
	@echo
	@printf "%-12s %b\n" "tools" "\e[0;90mInstall required binaries locally\e[0m"
	@printf "%-12s %b\n" "format" "\e[0;90mAuto format source code\e[0m"
	@printf "%-12s %b\n" "lint" "\e[0;90mRun golangci-lint\e[0m"
	@printf "%-12s %b\n" "test" "\e[0;90mRun all tests\e[0m"
	@printf "%-12s %b\n" "outdated" "\e[0;90mCheck outdated dependencies\e[0m"
	@echo

.PHONY: about
about:
	@echo "$(description)"

.PHONY: tools
tools:
	go install "golang.org/x/tools/cmd/goimports@latest"
	go install "github.com/psampaz/go-mod-outdated@v0.9.0"
	go install "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2"
	@echo

.PHONY: format
format:
	goimports -l -w -local github.com/webtask-go/ .
	@echo

.PHONY: lint
lint:
	golangci-lint run --out-format colored-line-number
	@echo

.PHONY: test
test:
	go test -count=1 -failfast -v ./...
	@echo

.PHONY: outdated
outdated:
	@echo -n "Checking outdated dependencies ..."
	@echo "\r$$(go list -u -m -json all | go-mod-outdated -update -direct)"
	@echo
