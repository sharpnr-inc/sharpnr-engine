# Name of your Go app (binary)
APP_NAME := sharpnr-engine

# Go files and directories to exclude from gofmt, vet, etc.
EXCLUDE := docs

# ------------------------------
# 🏁 Build and Run
# ------------------------------
.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -o $(APP_NAME) main.go

.PHONY: clean
clean:
	go clean
	rm -f $(APP_NAME)

# ------------------------------
# ✅ Testing
# ------------------------------
.PHONY: test
test:
	go test ./...

.PHONY: test-verbose
test-verbose:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -cover ./...

# ------------------------------
# 🧹 Formatting and Linting
# ------------------------------
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

# ------------------------------
# 🔄 Tidy up
# ------------------------------
.PHONY: tidy
tidy:
	go mod tidy

# ------------------------------
# 🧪 Test a specific package
# Usage: make test-pkg pkg=organisation
# ------------------------------
.PHONY: test-pkg
test-pkg:
	go test ./$(pkg)/...
