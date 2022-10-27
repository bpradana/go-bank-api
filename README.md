# PrivyID Backend Pretest
Bintang Pradana Erlangga Putra
bintangpradana02@gmail.com

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
│   ├── auth
│   │   └── authentication.go
│   ├── db
│   │   └── db.go
│   ├── util
│   │   └── iplocation.go
└── pkg
    ├── domain
    │   ├── user.go
    │   ├── user-balance.go
    │   └── user-balance-history.go
    ├── users
    │   ├── handler.go
    │   ├── repository.go
    │   ├── schema.go
    │   └── usecase.go
    ├── user-balances
    │   ├── handler.go
    │   ├── repository.go
    │   ├── schema.go
    │   └── usecase.go
    └── user-balance-histories
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
+ _auth_ - holds the authentication and authorization
+ _util_ - holds the utility functions

### _pkg_ folder contains the following:
+ _domain_ - holds the domain models
+ _user_ - holds the user business logic
+ _user_balance_ - holds the user balance business logic
+ _user_balance_history_ - holds the user balance history business logic


## Architecture
### Domain Driven Design
In this test I decided to use Domain Driven Design (DDD) as the architecture. The business logic is separated from the infrastructure and frameworks. The result is a flexible and loosely coupled system that can scale with changing requirements, thus making it easier to maintain and extend.
The DDD architecture is divided into 4 layers:
+ Domain Layer
+ Repository Layer
+ Use Case Layer
+ Handler Layer

Each layer has its own responsibility and is independent of the other layers. The domain layer is the core of the application and contains the blueprint of its business logic. The repository layer is responsible for the data access and persistence. The use case layer is responsible for the business logic. The handler layer is responsible for the HTTP request and response.

## Running the project
### Using Docker (Recommended)
```sh
$ docker compose build
$ docker compose up
```
### Bare Metal
```sh
$ go build -o main ./cmd/app/server.go
$ ./main
```

## API Documentation
The API documentation is available as a [Postman Collection](api-docs.json) and can be imported to Postman.
### Routes
#### User
+ `POST /api/v1/users/register` - Register a new user
+ `POST /api/v1/users/login` - User login
+ `GET /api/v1/users/logout` - User logout
#### Balance
+ `GET /api/v1/balances/check` - Get user balance
+ `PATCH /api/v1/balances/deposit` - Deposit balance to user
+ `PATCH /api/v1/balances/withdraw` - Withdraw balance from user
+ `PATCH /api/v1/balances/transfer` - Transfer balance to another user
+ `GET /api/v1/balances/histories` - Get user balance history