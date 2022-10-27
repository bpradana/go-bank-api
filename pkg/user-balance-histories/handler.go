package userBalanceHistories

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance history handler struct
type UserBalanceHistoryHandler struct {
	balanceHistoryUsecase domain.UserBalanceHistoryUsecase
}

// Constructor function to create new balance history handler
func NewUserBalanceHistoryHandler(e *echo.Group, balanceHistoryUsecase domain.UserBalanceHistoryUsecase) {
	h := &UserBalanceHistoryHandler{
		balanceHistoryUsecase: balanceHistoryUsecase,
	}

	// Routes
	e.GET("/histories", h.Check)
}

// Function to get balance history by username
func (h *UserBalanceHistoryHandler) Check(c echo.Context) error {
	// Get username from jwt context
	username, err := auth.GetUsernameFromToken(c)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting username from token, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Get balance history by username
	histories, err := h.balanceHistoryUsecase.Check(username)
	if err != nil {
		log.Println("[balance history] [handler] [Check] error getting balance history by username, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return balance history
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success",
		"histories": histories,
	})
}
