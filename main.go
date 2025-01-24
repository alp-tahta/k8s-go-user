package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var users = map[int]string{
	1: "Guts",
	2: "Rose",
	3: "Axl",
	4: "Zodd",
	5: "Kurosaki",
	6: "Ichigo",
	7: "Teresa",
	8: "Claire",
	9: "Derya",
}

func main() {
	port := ":8081"

	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Path Parameter
	router.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name, ok := users[i]
		if !ok {
			http.NotFound(w, r)
			return
		}

		_, err = w.Write([]byte(name))
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not write data", http.StatusInternalServerError)
			return
		}
	})

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on Port%s\n", port)
	log.Fatal(s.ListenAndServe())
}
