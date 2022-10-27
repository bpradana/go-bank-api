package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare claims struct
type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Function to generate token and set cookie
func GenerateTokenAndSetCookie(user *domain.User, c echo.Context) (string, error) {
	// Set token expiration time
	expiration := time.Now().Add(time.Hour * 24)

	// Generate token
	token, err := GenerateToken(user, expiration)
	if err != nil {
		log.Println("[auth] [GenerateTokenAndSetCookie] error when generating token:", err)
		return "", err
	}

	// Set token cookie
	SetTokenCookie(token, expiration, c)

	return token, nil
}

// Function to generate token
func GenerateToken(user *domain.User, expiration time.Time) (string, error) {
	// Declare claims
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	// Declare token with HS256 signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("[auth] [GenerateToken] error when generating token:", err)
		return "", err
	}

	return signedToken, nil
}

// Function to set token cookie
func SetTokenCookie(token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

// Function to delete token cookie
func DeleteTokenCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

// Function to get username from token
func GetUsernameFromToken(c echo.Context) (string, error) {
	// Get token from cookie
	cookie, err := c.Cookie("access_token")
	if err != nil {
		log.Println("[auth] [GetUsernameFromToken] error when getting cookie:", err)
		return "", err
	}

	// Parse token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println("[auth] [GetUsernameFromToken] error when parsing token:", err)
		return "", err
	}

	if !token.Valid {
		log.Println("[auth] [GetUsernameFromToken] token is not valid")
		return "", err
	}

	return claims.Username, nil
}
