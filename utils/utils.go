package utils

import (
	"activeReconBot/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

var (
	JwtSecret = GenerateRandomString(128)
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomString(length int) string {
	str := StringWithCharset(length, charset)
	log.Println("JWT Secret: " + str)
	return str
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
	token, err := at.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserIDFromToken(token string) (string, error) {
	claims := &JwtClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
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
		tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) { return []byte(JwtSecret), nil })
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

func RemoveFromDomainSlice(slice []models.Domain, idx int) []models.Domain {
	return append(slice[:idx], slice[idx+1:]...)
}

func RemoveFromSubDomainSlice(slice []models.Subdomain, idx int) []models.Subdomain {
	return append(slice[:idx], slice[idx+1:]...)
}

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
