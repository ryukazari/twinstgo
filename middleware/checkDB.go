package middleware

import (
	"net/http"

	"github.com/ryukazari/twinstgo/database"
)

// CheckDB check the connection to the database
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !database.CheckConnection() {
			http.Error(w, "Lost connection to the database", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
