package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/cmd/db"

	userBalanceHistories "gitlab.com/bpradana/privy-pretest/pkg/user-balance-histories"
	userBalances "gitlab.com/bpradana/privy-pretest/pkg/user-balances"
	users "gitlab.com/bpradana/privy-pretest/pkg/users"
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
	config := middleware.JWTConfig{
		Claims:      &auth.Claims{},
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:access_token",
	}

	userRepository := users.NewUserRepository(db)
	balanceRepository := userBalances.NewUserBalanceRepository(db)
	balanceHistoryRepository := userBalanceHistories.NewUserBalanceHistoryRepository(db)

	userUsecase := users.NewUserUsecase(userRepository, balanceRepository)
	balanceUsecase := userBalances.NewUserBalanceUsecase(userRepository, balanceRepository, balanceHistoryRepository)
	balanceHistoryUsecase := userBalanceHistories.NewUserBalanceHistoryUsecase(userRepository, balanceRepository, balanceHistoryRepository)

	usersGroup := e.Group("/api/v1/users")
	users.NewUserHandler(usersGroup, userUsecase)

	balancesGroup := e.Group("/api/v1/balances")
	balancesGroup.Use(middleware.JWTWithConfig(config))
	userBalances.NewUserBalanceHandler(balancesGroup, balanceUsecase)
	userBalanceHistories.NewUserBalanceHistoryHandler(balancesGroup, balanceHistoryUsecase)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
