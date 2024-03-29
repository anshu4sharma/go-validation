package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go-validation/internal/server"
	"os"
	"strconv"
)

func main() {

	server := server.New()

	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
