package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking authentication")
			var flag bool = true

			if flag {
				f(w, r)
			} else {
				fmt.Println("Authentification failed")
				return
			}
		}
	}
}

func Loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}
