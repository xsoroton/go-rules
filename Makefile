run:
	go run ./


run-in-docker:
	docker build -t go-rules . && docker run -it --rm --name go-rules go-rules

# run the tests
test:
	go test ./... -v

# run the tests with coverage
test-cover:
	go test -covermode=count -coverprofile=coverage.out ./...

test-cover-stats: test-cover
	@if [ -e coverage.out ]; then \
		go tool cover -func=coverage.out; \
	fi