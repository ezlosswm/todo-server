# TODO API Server 
### An easy-to-use RESTful Todo API written in Go and PostgresSQL.

## Usage 
1. Clone this repo 
`git clone https://github.com/EzlosSWM/todo-server.git`

2. Navigate to the directory 
`cd todo-server`

3. Download dependancies 
`go mod download && go mod verify`

4. Run
    - (make)   
      - `make live`
    - (Go)   
      - `go run *.go`


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

*DELETE /todo/{id}*
- Deletes the selected item from the list of todo 

*PUT /todo/{id}*
- Marks an item as completed

## Database (Postgres)
The database in use is PostgresSQL. The API will not run if you don't have a `.env` in this root directory. You can use the template below: 
1. Create **.env** file. `touch .env`

2. Add environment variables to **.env** file
```
HOSTADDR=localhost
USER_NAME=
DB_NAME=
PASSWORD=
```
*Note*: for development puposes, `HOSTADDR` is the localhost and should not be changed.

## Notes
- The server runs on port 3000.
- The project is still in progress.
- Working on dockerizing. 
- Wanting to return custom errors as json 
- `make live` on the Makefile needs nodemon running to use; 
  - Install nodemon using npm 
`npm install -g nodemon`

*Once installed, `make live` should work normally.*


## Contact
Twitter: [@EzlosSWM](https://twitter.com/EzlosSWM)

Github: [@EzlosSWM](https://github.com/EzlosSWM)

[Socials](https://ezlos-redirect.vercel.app/)

Email: ezlosswm@gmail.com