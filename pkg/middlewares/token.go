package middlewares

// import (
// 	"context"
// 	"net/http"
// )

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		tokenStr := r.Header.Get("Authorization")
// 		if tokenStr == "" {
// 			http.Error(w, "Missing token", http.StatusUnauthorized)
// 			return
// 		}

// 		claims, err := VerifyToken(tokenStr)
// 		if err != nil {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// store user ID in context
// 		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
