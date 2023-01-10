# TODO API Server 
### An easy-to-use RESTful Todo API written in Go with in memory storage.

## Usage 
```bash
# Installation 
go get 

# Makefile
make run 

# Basic 
go run *.go 

# Binary 
./bin/todoServer

```
## Getting Started 
Currently, there is one that handles GET and POST requests:  
```JSON
localhost:3000/todo 

localhost:3000/todo?activity=<todo-item>
```


## Notes
- The server runs on port 3000
- If you'd want to change the port navigate to `main.go` file and change the port number to the desired value.
- The project is still in progress, endpoints for DELETE, UPDATE and GET BY ID. 


## Contact
[Twitter](https://twitter.com/EzlosSWM)

[Github](https://github.com/EzlosSWM)

[Email](ezlosswm@gmail.com)