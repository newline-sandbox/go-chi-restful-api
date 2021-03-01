run_service:
	go run .

install_deps:
	rm -f go.mod go.sum
	go mod init github.com/newline-sandbox/go-chi-restful-api
	go mod tidy

test:
	go test -v ./...

bench:
	go test -bench=. ./routes -run=^$