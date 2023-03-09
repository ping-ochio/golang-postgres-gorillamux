package users

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var my_secret_key = []byte(os.Getenv("Secret_key"))

// Generating an encrypted password to save in database
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "Nothing here", err

	}
	Password := string(bytes)
	return Password, nil
}

// Check if the password is valid,
// "password" is from database,
// "providedPassword" is from login user,
// so first we need fetch password from database for this "user_name" or "email"
func CheckPassword(password, providedPassword string) (*JWTResponse, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return nil, err

	}
	return getTokens(), nil
}

func getTokens() *JWTResponse {
	token := singToken()
	refresh_token := singRefreshToken()

	return &JWTResponse{
		Token:   token,
		Refresh: refresh_token,
	}
}

func singToken() string {
	// Generating token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(my_secret_key)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return tokenString

}

func singRefreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "baz",
		"nbf":     time.Now().Add(time.Hour * 24).Unix(),
		"refresh": true,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(my_secret_key)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return tokenString

}

func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			berarerToken := r.Header.Get("Cookie")
			r.Cookies()
			if validateToken(berarerToken) {

				next.ServeHTTP(w, r)
				return
			}
			http.Error(w, "invalid token", http.StatusUnauthorized)
		})

}

func validateToken(s string) bool {
	s1 := strings.Split(s, "Refresh=")
	token := strings.Replace(s1[0], "bearer=", "", 1)

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return my_secret_key, nil
	})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return t.Valid
}

/*

// we need to know when the acces token expire

func validateRefreshToken(s string) bool {
	s1 := strings.Split(s, "Refresh=")
	refreshToken := strings.Replace(s1[1], "bearer=", "", 1)

	t, err := jwt.Parse(refreshToken , func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return my_secret_key, nil
	})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return t.Valid
}*/
