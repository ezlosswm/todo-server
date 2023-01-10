build: 
	@go build -o bin/todoServer

run: build
	@./bin/todoServer
