package balances

import (
	"log"
	"net/http"

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
		"balance": balance,
	})
}
