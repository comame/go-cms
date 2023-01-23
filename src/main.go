package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var HOSTNAME = os.Getenv("HOSTNAME")

func checkEnv() {
	if HOSTNAME == "" {
		panic("Not present")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != "GET" && r.Method != "HEAD" {
		res, _ := json.Marshal(ErrorResponse{
			Message: "Method Not Allowed",
			Code:    NotFound.String(),
		})
		http.Error(w, string(res), http.StatusMethodNotAllowed)
		return
	}

	done := HandleCms(w, r)

	if done {
		return
	}

	res, _ := json.Marshal(ErrorResponse{
		Message: "Not Found",
		Code:    NotFound.String(),
	})
	http.Error(w, string(res), http.StatusNotFound)
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("%v", err)
}

func Main() {
	var cmd string = "serve"
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "serve":
		{
			println("Start serving...")
			serve()
		}
	default:
		println("commands: serve, compose")
	}
}
