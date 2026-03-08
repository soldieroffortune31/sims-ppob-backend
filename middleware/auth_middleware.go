package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/repository"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(userRepository repository.UserRepository, db *sql.DB) func(httprouter.Handle) httprouter.Handle {

	return func(next httprouter.Handle) httprouter.Handle {

		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				panic(exception.NewUnauthorizedError("Missing token"))
			}

			tokenString := strings.Split(authHeader, "Bearer ")[1]

			token, err := helper.ValidateToken(tokenString)
			if err != nil {
				panic(exception.NewUnauthorizedError("Invalid token"))
			}

			claims := token.Claims.(jwt.MapClaims)

			userId := int(claims["user_id"].(float64))

			tx, err := db.Begin()
			helper.PanicIfError(err)
			defer helper.CommitOrRollback(tx)

			user, err := userRepository.FindById(r.Context(), tx, userId)
			if err != nil {
				panic(exception.NewUnauthorizedError("User not found"))
			}

			// simpan user di context
			ctx := context.WithValue(r.Context(), "user", user)

			next(w, r.WithContext(ctx), ps)
		}
	}
}
