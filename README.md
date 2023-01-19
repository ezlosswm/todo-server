# TODO API Server 
### An easy-to-use RESTful Todo API written in Go with in memory storage.

## Usage 
```bash
# Installation 
go get https://github.com/EzlosSWM/todo-server

cd todo-server

# Makefile
make run 

# Basic 
go run *.go 

# Binary 
./bin/todoServer

```
## Getting Started 
List of current endpoints on a browser.
```JSON
localhost:3000/todo 

localhost:3000/todo/{id}

localhost:3000/todo?activity=<todo-item>
```

### Endpoints 
*GET /todo* 
- Returns all todo items

*GET /todo/{id}*
- Returns the specified todo item 

*POST /todo*
- Creates a new todo item 


```JSON
{
    "todo": "test item"
}
```
*DELETE /todo{id}*
- Deletes the selected item from the list of todo 


## Notes
- The server runs on port 3000
- If you'd want to change the port navigate to `main.go` file and change the port number to the desired value.
- The project is still in progress, endpoints for PUT method. 
- `make live` on the Makefile needs nodemond running to use; 
```bash 
# install nodeman using npm 
npm install -g nodemon
```


## Contact
[Twitter](https://twitter.com/EzlosSWM)

[Github](https://github.com/EzlosSWM)

> Email: ezlosswm@gmail.com