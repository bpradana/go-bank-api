package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.com/bpradana/privy-pretest/cmd/db"
	"gitlab.com/bpradana/privy-pretest/pkg/balances"
	"gitlab.com/bpradana/privy-pretest/pkg/users"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("[app] [server] error loading .env file, err: ", err.Error())
	}

	// Connect to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("[app] [server] error connecting to database, err: ", err.Error())
	}

	// Create echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	userRepository := users.NewUserRepository(db)
	balanceRepository := balances.NewBalanceRepository(db)

	userUsecase := users.NewUserUsecase(userRepository, balanceRepository)

	usersGroup := e.Group("/api/v1/users")
	users.NewUserHandler(usersGroup, userUsecase)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
