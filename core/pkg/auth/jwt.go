package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	// ErrValidatingSignature error
	ErrValidatingSignature string = "ErrValidatingSignature"

	// ErrCreatingToken error
	ErrCreatingToken string = "ErrCreatingToken"

	// ErrValidatingClaims error
	ErrValidatingClaims string = "ErrValidatingClaims"
)

var userTokenSecret string = (os.Getenv("USER_TOKEN_SECRET"))
var svcTokenSecret string = (os.Getenv("SVC_TOKEN_SECRET"))

// userClaims containes the token claims
type userClaims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

// VerifyUserAccessToken varifies the user token
func VerifyUserAccessToken(tokenString string, userID string, userRole string) (valid bool, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Verify the signing method is HMAC
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s != jwt.SigningMethodHS512 {
			return nil, errors.New(ErrValidatingSignature)
		}

		return []byte(userTokenSecret), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		if claims.UserID != userID || claims.UserRole != userRole {
			return false, errors.New(ErrValidatingClaims)
		}
	}

	return true, nil
}

// CreateUserAccessToken return a new user token
func CreateUserAccessToken(userID string, userRole string) (token string, err error) {
	token, err = createUserAccessToken(userID, userRole)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(ErrCreatingToken)
	}

	return
}

func createUserAccessToken(userID string, userRole string) (token string, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	c := userClaims{
		UserID:   userID,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			Id:        id.String(),
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err = t.SignedString([]byte(userTokenSecret))

	return
}
