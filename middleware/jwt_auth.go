package middleware

import (
	handler "SQLite_JWT/handler"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("abcdefghijklmnopq")

type Claims struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	jwt.StandardClaims
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do somethings
		fmt.Printf("Check token by middleware")
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			handler.ResponseErr(w, http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Bearer jwt_token
		if len(splitted) != 2 {
			handler.ResponseErr(w, http.StatusForbidden)
			return
		}

		tokenPart := splitted[1]
		tk := &Claims{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			fmt.Println(err)
			handler.ResponseErr(w, http.StatusInternalServerError)
			return
		}

		if token.Valid {
			handler.ResponseOk(w, token.Claims)
		}

		next.ServeHTTP(w, r)
	})
}
