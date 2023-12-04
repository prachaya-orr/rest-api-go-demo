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

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
		next.ServeHTTP(w, req)
	}
}

func main() {
	http.HandleFunc("/users", logMiddleware(usersHandler))
	http.HandleFunc("/health", logMiddleware(healthHandler))

	log.Println("Server stared at http://localhost:2566")
	log.Fatal(http.ListenAndServe(":2566", nil))
	log.Println("bye bye!")
}
