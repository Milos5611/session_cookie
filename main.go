package main

import (
	"log"
	"net/http"
)

func CORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")

		// If it's a preflight request, stop here and send a response with the headers.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, pass the request to the next handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	// "Signin" and "Signup" are handlers that we have to implement
	http.Handle("/signin", CORSHeadersMiddleware(http.HandlerFunc(Signin)))
	http.Handle("/welcome", CORSHeadersMiddleware(http.HandlerFunc(Welcome)))
	http.Handle("/refresh", CORSHeadersMiddleware(http.HandlerFunc(Refresh)))
	http.HandleFunc("/logout", Logout)
	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
