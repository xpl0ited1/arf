package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	JWT_SECRET = "SECRET"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type JwtClaims struct {
	UserID   string `json:"user"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateTokenForUser(userId string, username string) (string, error) {

	expirationTime := time.Now().Add((24 * time.Hour) * 7).Unix()
	claims := &JwtClaims{
		UserID:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime,
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserIDFromToken(token string) (string, error) {
	claims := &JwtClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return "", err
	}
	if !tkn.Valid {
		return "", err
	}
	return claims.UserID, nil
}

func CheckIfAuthorized(tokenUserId string, resourceUserId string) bool {
	if tokenUserId == resourceUserId {
		return true
	} else {
		return false
	}
}

func GetUsernameFromToken(jwtToken string) string {
	claims := &JwtClaims{}
	if jwtToken != "" {
		tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) { return []byte(JWT_SECRET), nil })
		if err != nil {
			log.Printf("%s\n",
				err.Error(),
			)
			return ""
		}

		if !tkn.Valid {
			log.Printf("%s\n",
				"Invalid Token",
			)
			return ""
		} else {
			return claims.Username
		}

	} else {
		return ""
	}
}

func ShodanRecon() {
	//TODO: This will fill the primary ports and technologies with their cve's
}
