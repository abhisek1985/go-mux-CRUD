package middlewares

import (
    "os"
    "log"
	"net/http"
	"github.com/joho/godotenv"   // package used to read the .env file
)


func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			w.WriteHeader(401)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	return username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD")
}