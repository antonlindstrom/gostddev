all:
	@mkdir -p bin/
	@go build -o bin/go-stddev go-stddev.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

clean:
	@rm -rf bin/
