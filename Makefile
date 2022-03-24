.DEFAULT_GOAL = build
BUILD_FOLDER  = $(shell pwd)/build
FLAGS_LINUX   = GOOS=linux GOARCH=amd64 CGO_ENABLED=1

.PHONY: build
build:
	@echo "[builder] Building tips executable"
	go get
	$(FLAGS_LINUX) go build -o $(BUILD_FOLDER)/tips
	@echo "[builder] Done!"

.PHONY: lint
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

.PHONY: fmt
fmt:
	@echo "[gofmt] Formatting code"
	gofmt -s -w .

.PHONY: clean
clean:
	rm -rf $(BUILD_FOLDER)
