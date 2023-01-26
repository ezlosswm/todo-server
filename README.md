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
```
## Getting Started 
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

*PUT /todo/{id}*
- Allows the user to edit the todo without marking as completed. 

```JSON
// Befoee
{
  "items": [
    {
      "id": 81,
      "activity": "go runnning",
      "completed_at": "0001-01-01T00:00:00Z",
      "completed": false
    }
  ]
}

// After
{
  "items": [
    {
      "id": 0,
      "activity": "go running",
      "completed_at": "0001-01-01T00:00:00Z",
      "completed": false
    }
  ]
}
```

*PATCH /todo/{id}*
- Marks an item as completed and logs the time it was completed. 

*DELETE /todo{id}*
- Deletes the selected item from the list of todo 


## Notes
- The server runs on port 3000.
- The project is still in progress.
- Working on dockerizing. 
- Wanting to return custom errors as json 
- Wanting to link to a database eventually, preferably PostgresSQL.
- `make live` on the Makefile needs nodemon running to use; 
```bash 
# install nodeman using npm 
# once installing nodemon, make live should be working
npm install -g nodemon

```


## Contact
Twitter: [@EzlosSWM](https://twitter.com/EzlosSWM)

Github: [@EzlosSWM](https://github.com/EzlosSWM)

Email: ezlosswm@gmail.com