package main

import (
	"fmt"
	"os"

	"github.com/deputadosemfoco/api-gateway/config"
	"github.com/deputadosemfoco/api-gateway/routes"
	"github.com/dimiro1/banner"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	config.Init()

	in, _ := os.Open("banner.txt")
	defer in.Close()
	banner.Init(os.Stdout, true, false, in)

	e := routes.SetUp()
	port := os.Getenv("PORT")

	fmt.Printf("gateway service will run on port %s", port)
	fmt.Println("")

	e.Run(standard.New(":" + port))
}
