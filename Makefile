PROJECT_NAME=testcntrns
CURRENT_DIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

up:
	docker build -t $(PROJECT_NAME) .
	docker run --rm $(PROJECT_NAME)

test:
	go test ./... -short

lint:
	docker run --rm -v $(CURRENT_DIR):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run --timeout=3m0s -E bodyclose  -E deadcode -E depguard -E dogsled -E dupl -E errcheck -E gochecknoinits -E goconst -E gocritic -E gocyclo -E gofmt -E gosec -E goprintffuncname -E gosimple -E govet -E ineffassign -E interfacer -E misspell -E nakedret -E rowserrcheck -E staticcheck -E structcheck -E stylecheck -E typecheck -E unconvert -E unparam -E varcheck -E whitespace -E prealloc