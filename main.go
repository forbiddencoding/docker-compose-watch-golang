package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	slog.Info("starting server...")

	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error starting server: %s", err.Error())
		os.Exit(1)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("request received")
	_, _ = fmt.Fprintf(w, "Request time: %s", time.Now().Format(time.RFC3339))
}
