test:
	@go get golang.org/x/tools/cmd/cover
	@go test ./... -cover

test_local:
	@go get golang.org/x/tools/cmd/cover
	@go test ./... -coverprofile="/tmp/go-cover.tmp" $@
	@go tool cover -html="/tmp/go-cover.tmp"
	@unlink "/tmp/go-cover.tmp"