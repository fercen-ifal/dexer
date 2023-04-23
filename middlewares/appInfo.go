package middlewares

import "net/http"

func AppInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-App-Name", "dexer")
		w.Header().Set("X-App-Type", "microservice")
		w.Header().Set("X-App-Domain", "dexer.fercen.ifal.edu.br")

		next.ServeHTTP(w, r)
	})
}
