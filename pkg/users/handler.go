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
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/logout", h.Logout)
}

// Function to register user
func (h *UserHandler) Register(c echo.Context) error {
	// Create user struct
	user := new(SchemaRegisterUser)

	// Bind request body to user struct
	if err := c.Bind(user); err != nil {
		log.Println("[users] [handler] [Register] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println("[users] [handler] [Register] error validating request body, err: ", err.Error())
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
	registeredUser, err := h.userUsecase.Register(userToRegister)
	if err != nil {
		log.Println("[users] [handler] [Register] error creating user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created",
		"data":    registeredUser,
	})
}

// Function to login user
func (h *UserHandler) Login(c echo.Context) error {
	// Create user struct
	user := new(SchemaLoginUser)

	// Bind request body to user struct
	if err := c.Bind(user); err != nil {
		log.Println("[users] [handler] [Login] error binding request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println("[users] [handler] [Login] error validating request body, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Create user domain
	userToLogin := &domain.User{
		Email:    user.Email,
		Password: user.Password,
	}

	// Login user
	loggedInUser, token, err := h.userUsecase.Login(userToLogin, c)
	if err != nil {
		log.Println("[users] [handler] [Login] error logging in user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User logged in",
		"data":    loggedInUser,
		"token":   token,
	})
}

// Function to logout user
func (h *UserHandler) Logout(c echo.Context) error {
	// Logout user
	err := h.userUsecase.Logout(c)
	if err != nil {
		log.Println("[users] [handler] [Logout] error logging out user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User logged out",
	})
}
