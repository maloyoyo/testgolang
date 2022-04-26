package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	Success bool
	ErrCode error
	Value   float64
}

func main() {
	log.Print("Starting the service...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var res Result

		a, err := strconv.Atoi(r.URL.Query().Get("a"))
		if err != nil {
			res.ErrCode = err
		}
		b, err := strconv.Atoi(r.URL.Query().Get("b"))
		if err != nil {
			res.ErrCode = err
		}

		switch r.URL.Path {
		case "/add":
			res.Value = float64(a + b)
		case "/sub":
			res.Value = float64(a - b)
		case "/mul":
			res.Value = float64(a * b)
		case "/div":
			res.Value = float64(a) / float64(b)
		}
		res.Success = true

		json.NewEncoder(w).Encode(res)
	})

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8186", nil))
}
