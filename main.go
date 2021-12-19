package main

import (
	"fmt"
	"os"
	"test-api-bs/internal/api"
)

func main() {

	router := api.New(api.SetupDB())

	APP_PORT := os.Getenv("PORT")
	PORT := fmt.Sprintf(":%s", APP_PORT)
	router.Run(PORT)
}
