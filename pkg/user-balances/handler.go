package userBalances

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance handler struct
type UserBalanceHandler struct {
	balanceUsecase domain.UserBalanceUsecase
}

// Constructor function to create new balance handler
func NewUserBalanceHandler(e *echo.Group, balanceUsecase domain.UserBalanceUsecase) {
	h := &UserBalanceHandler{
		balanceUsecase: balanceUsecase,
	}

	// Routes
	e.GET("/check", h.Check)
	e.PATCH("/deposit", h.Deposit)
	e.PATCH("/withdraw", h.Withdraw)
	e.PATCH("/transfer", h.Transfer)
}

// Function to get balance by username
func (h *UserBalanceHandler) Check(c echo.Context) error {
	// Get username from jwt context
	username, err := auth.GetUsernameFromToken(c)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting username from token, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Get balance by username
	balance, err := h.balanceUsecase.Check(username)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting balance by username, err: ", err.Error())
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
func (h *UserBalanceHandler) Deposit(c echo.Context) error {
	// Get username from jwt context
	username, err := auth.GetUsernameFromToken(c)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting username from token, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Create deposit struct
	deposit := new(SchemaDepositWithdraw)
	if err := c.Bind(deposit); err != nil {
		log.Println("[balances] [handler] [Deposit] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate deposit struct
	validate := validator.New()
	if err := validate.Struct(deposit); err != nil {
		log.Println("[balances] [handler] [Deposit] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Deposit balance
	balance, err := h.balanceUsecase.Deposit(username, deposit.Amount, c)
	if err != nil {
		log.Println("[balances] [handler] [Deposit] error depositing balance, err: ", err.Error())
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

// Function to withdraw balance
func (h *UserBalanceHandler) Withdraw(c echo.Context) error {
	// Get username from jwt context
	username, err := auth.GetUsernameFromToken(c)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting username from token, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Create withdraw struct
	withdraw := new(SchemaDepositWithdraw)
	if err := c.Bind(withdraw); err != nil {
		log.Println("[balances] [handler] [Withdraw] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate withdraw struct
	validate := validator.New()
	if err := validate.Struct(withdraw); err != nil {
		log.Println("[balances] [handler] [Withdraw] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Withdraw balance
	balance, err := h.balanceUsecase.Withdraw(username, withdraw.Amount, c)
	if err != nil {
		log.Println("[balances] [handler] [Withdraw] error withdrawing balance, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return balance
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Balance withdrawn",
		"balance": balance,
	})
}

// Function to transfer balance
func (h *UserBalanceHandler) Transfer(c echo.Context) error {
	// Get username from jwt context
	username, err := auth.GetUsernameFromToken(c)
	if err != nil {
		log.Println("[balances] [handler] [Check] error getting username from token, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Create transfer struct
	transfer := new(SchemaTransfer)
	if err := c.Bind(transfer); err != nil {
		log.Println("[balances] [handler] [Transfer] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate transfer struct
	validate := validator.New()
	if err := validate.Struct(transfer); err != nil {
		log.Println("[balances] [handler] [Transfer] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Transfer balance
	balance, err := h.balanceUsecase.Transfer(username, transfer.Amount, transfer.Username, c)
	if err != nil {
		log.Println("[balances] [handler] [Transfer] error transferring balance, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return balance
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Balance transferred",
		"balance": balance,
	})
}
