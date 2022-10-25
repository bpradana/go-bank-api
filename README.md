# PrivyID Backend Pretest

## Test Description
1. Create a web service with RESTful API type with the database schema listed above
2. USER domain can do the login and logout process
3. USER domain can perform additional balance actions which will be recorded in the user_balance table
4. USER domain can perform other transfer actions to other USER
5. If necessary, it is allowed to add columns with any data types to the table
6. Each endpoint of the web service must put forward a security issue with the authorization method
7. Apply the Single-responsibility principle
8. Using Ruby or Golang is a plus

## Tech Stack
1. Go 1.19
2. Echo v4
3. PostgresQL 14

## Folder Structure
### Folder Tree

```
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
├── cmd
│   ├── app
│   │   └── server.go
│   ├── db
│   │   └── db.go
└── pkg
    ├── domain
    │   ├── user.go
    │   ├── user_balance.go
    │   └── user_balance_history.go
    ├── user
    │   ├── handler.go
    │   ├── repository.go
    │   └── usecase.go
    ├── user_balance
    │   ├── handler.go
    │   ├── repository.go
    │   └── usecase.go
    └── user_balance_history
        ├── handler.go
        ├── repository.go
        └── usecase.go
```

### _root_ folder contains the following:
+ _cmd_ - holds the main application and configurations
+ _pkg_ - holds the business logic

### _cmd_ folder contains the following:
+ _app_ - holds the main application
+ _db_ - holds the database connection

### _pkg_ folder contains the following:
+ _domain_ - holds the domain models
+ _user_ - holds the user business logic
+ _user_balance_ - holds the user balance business logic
+ _user_balance_history_ - holds the user balance history business logic


## Running the project
### Using Docker
```sh
$ docker compose build && docker compose up
```
### Bare Metal
```sh
$ go build -o main ./cmd/app/server.go
$ ./main
```