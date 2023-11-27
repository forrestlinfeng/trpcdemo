
NAME:=trpcdemo
LDFLAGS :=

.PHONY: build
## build: Compile the packages.
build:
	@env GOOS=linux go build -o $(NAME)


.PHONY: lint
lint:
	@go vet -structtag=false -tags="${GO_TAGS}" ${GO_PACKAGES}


# NOTE: cgo is required by go test
.PHONY: test
test:
	CGO_ENABLED=1 go test -v -race -cover -failfast -vet=off -tags="${GO_TAGS}" ${GO_PACKAGES}

.PHONY: clean
clean:
	rm -f $(NAME)
	go clean -cache ${GO_PACKAGES}



