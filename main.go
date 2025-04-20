package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CAMELNINJA/apiguard/config"
	"github.com/CAMELNINJA/apiguard/routes"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := routes.SetupRouter(cfg)

	fmt.Println("🚀 Listening on :8080")
	http.ListenAndServe(":8080", r)
}
