package middlewares

import (
	"github.com/mbswe/selma"
	"log"
	"net/http"
)

// CustomMiddleware is a custom middleware
func CustomMiddleware(app *selma.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform custom operations here
			log.Default().Println("Executing custom middleware")

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
