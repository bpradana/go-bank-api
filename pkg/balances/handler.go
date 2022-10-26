package balances

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance handler struct
type BalanceHandler struct {
	balanceUsecase domain.BalanceUsecase
}

// Constructor function to create new balance handler
func NewBalanceHandler(e *echo.Group, balanceUsecase domain.BalanceUsecase) {
	h := &BalanceHandler{
		balanceUsecase: balanceUsecase,
	}

	// Routes
	e.GET("/check", h.CheckBalance)
	e.PATCH("/deposit", h.Deposit)
}

// Function to get balance by username
func (h *BalanceHandler) CheckBalance(c echo.Context) error {
	// Get username from jwt context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.Claims)
	username := claims.Username

	// Get balance by username
	balance, err := h.balanceUsecase.CheckBalance(username)
	if err != nil {
		log.Println("[balances] [handler] error getting balance by username, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return balance
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"balance": balance,
	})
}

// Function to deposit balance
func (h *BalanceHandler) Deposit(c echo.Context) error {
	// Get username from jwt context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.Claims)
	username := claims.Username

	// Create deposit struct
	deposit := new(SchemaDeposit)
	if err := c.Bind(deposit); err != nil {
		log.Println("[balances] [handler] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate deposit struct
	validate := validator.New()
	if err := validate.Struct(deposit); err != nil {
		log.Println("[balances] [handler] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Deposit balance
	balance, err := h.balanceUsecase.Deposit(username, deposit.Amount)
	if err != nil {
		log.Println("[balances] [handler] error depositing balance, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return balance
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Balance deposited",
		"balance": balance,
	})
}
