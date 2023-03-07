.DEFAULT_GOAL = build
BUILD_FOLDER  = $(shell pwd)/build
FLAGS_LINUX   = GOOS=linux GOARCH=amd64 CGO_ENABLED=1

.PHONY: check
check:
	@echo "[lint] Running go vet"
	go vet ./...
	@echo "[lint] Running staticheck on codebase"
	@staticcheck ./...

.PHONY: build
build:
	@echo "[builder] Building tips executable"
	go get
	$(FLAGS_LINUX) go build -o $(BUILD_FOLDER)/tips
	@echo "[builder] Done!"

.PHONY: fmt
fmt:
	@echo "[gofmt] Formatting code"
	gofmt -s -w .

.PHONY: clean
clean:
	rm -rf $(BUILD_FOLDER)
