package users

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare user handler struct
type UserHandler struct {
	userUsecase domain.UserUsecase
}

// Constructor function to create new user handler
func NewUserHandler(e *echo.Group, userUsecase domain.UserUsecase) {
	h := &UserHandler{
		userUsecase: userUsecase,
	}

	// Routes
	e.POST("/users/register", h.Register)
}

// Function to register user
func (h *UserHandler) Register(c echo.Context) error {
	// Create user struct
	user := new(SchemaRegisterUser)

	// Bind request body to user struct
	if err := c.Bind(user); err != nil {
		log.Println("[users] [handler] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println("[users] [handler] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Create user domain
	userToRegister := &domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	// Create user
	_, err := h.userUsecase.Register(userToRegister)
	if err != nil {
		log.Println("[users] [handler] error creating user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created",
	})
}
