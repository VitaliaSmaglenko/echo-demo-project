package main

import (
	"echo-demo-project/server"
	"echo-demo-project/server/routes"
	"echo-demo-project/server/validation"
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := server.NewServer()
	app.Echo.Validator = validation.NewCustomValidator(validator.New())


	routes.ConfigureRoutes(app)
	_ = app.Start(os.Getenv("PORT"))
}
