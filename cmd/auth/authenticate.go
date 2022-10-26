package auth

import (
	"os"

	"github.com/golang-jwt/jwt"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare claims struct
type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Function to generate token
func GenerateToken(user *domain.User) (string, error) {
	// Declare claims
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Declare token with HS256 signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
