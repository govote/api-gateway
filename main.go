package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/engine/standard"
	"github.com/vitorsalgado/la-democracia/gateway/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("%s : %s", "Error loading .env file", err)
	}

	router := routes.Router{}
	e := router.SetUp()
	port := os.Getenv("Gateway_Port")

	fmt.Printf("gateway service running on port %s", port)

	e.Run(standard.New(":" + port))
}
