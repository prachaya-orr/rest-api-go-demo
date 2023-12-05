package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "John", Age: 20},
	{ID: 2, Name: "Ohm", Age: 26},
}

func usersHandler(w http.ResponseWriter, req *http.Request) {
	u, p, ok := req.BasicAuth()
	log.Println("auth:", u, p, ok)

	if req.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

		}

		w.Write(b)
		return
	}

	if req.Method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			// err from fail io.ReadAll
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		users = append(users, u)

		fmt.Fprintf(w, "User %s has been created", u.Name)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, req)
// 		log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
// 	}
// }

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		log.Println("auth:", u, p, ok)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`can't parse the basic auth`))
			return
		}

		if u != "admin" || p != "1234" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`username or password incorrect`))
			return
		}
		fmt.Println("Auth passed.")
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux() // multiplexer
	mux.HandleFunc("/users", AuthMiddleware((usersHandler)))
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}

	srv := http.Server{
		Addr:    ":2566",
		Handler: logMux,
	}

	log.Println("Server stared at http://localhost:2566")
	log.Fatal(srv.ListenAndServe())
	log.Println("bye bye!")
}
