dev:
	go run .

test: 
	go test ./...

test-cover:
	go test ./... -covermode=count