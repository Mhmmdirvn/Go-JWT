package middlewares

import (
	"Combine-Gorm-Mux-Jwt/controllers"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func MiddlewareJWTAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizatinHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizatinHeader, "Bearer") {
			http.Error(w, "Invalid Token", http.StatusBadRequest)
			return
		}

		tokeString := strings.Replace(authorizatinHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokeString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing Method Invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing Method Invalid")
			}

			return controllers.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ctx := context.WithValue(context.Background(), "userInfo", claims)
		// r = r.WithContext(ctx)
		next(w, r)
	})
}