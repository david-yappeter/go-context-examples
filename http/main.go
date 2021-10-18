package main

import (
	"context"
	"fmt"
	"net/http"
)

type Claims struct {
	ID int `json:"id"`
}

var contextKey = "ctx-key"

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Get UserID From Token
		// ....
		// ....

		fmt.Println("Insert Claim into Context")
		newCtx := context.WithValue(r.Context(), contextKey, &Claims{
			ID: 1, // example User ID: 1
		})

		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}

func CtxClaim(ctx context.Context) *Claims {
	raw, _ := ctx.Value(contextKey).(*Claims)
	return raw
}

func main() {

	http.Handle("/", middleware1(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// After Insert Claim into Context

		fmt.Printf("%+v\n", CtxClaim(r.Context()))
		fmt.Fprintf(w, "%+v\n", CtxClaim(r.Context()))
	})))

	http.ListenAndServe(":8080", nil)
}
