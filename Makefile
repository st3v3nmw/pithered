LEVELDB_PATH = "level.db"

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: test
test:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out

.PHONY: build
build:
	go build -o pithered ./...

.PHONY: run
run: build
	./pithered --leveldb=${LEVELDB_PATH} --verbose
