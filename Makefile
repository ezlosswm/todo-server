build: 
	@go build -o bin/todoServer

run: build
	@./bin/todoServer

live: 
	@nodemon --exec go run *.go --signal SIGTERM