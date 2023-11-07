package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/santos95mat/go-fiber-request/user"
)

var userRequest user.UserRequest

func app() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/user", userRequest.GetMany)
	app.Post("/user", userRequest.Create)
	app.Post("/login", userRequest.Login)

	err := app.Listen(":3030")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	app()
}
