all:
	@mkdir -p bin/
	@go build -o bin/gostddev gostddev.go

get-deps:
	@go get -d -v ./...

format:
	@go fmt ./...

test:
	@go test ./...

coverage:
	@go get code.google.com/p/go.tools/cmd/cover
	@go test -cover ./...

benchmark:
	@go test -bench ./...

clean:
	@rm -rf bin/
