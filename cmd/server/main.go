package main

import (
	"log"
	"net/http"
	"os"

	"github.com/travacry/prometey-app/internal/web"
)

func main() {
	addr := ":8080"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, web.NewMux()))
}
