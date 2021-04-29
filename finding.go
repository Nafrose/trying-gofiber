package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"trying-gofiber/findingpeople"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello World")
	})

	app.Get("/people", findingpeople.ShowAll)
	app.Get("/people/:input", findingpeople.FindOne)
	app.Get("/people/:job/:name", findingpeople.FindWithTwo)
	app.Post("/people", findingpeople.CreteNew)

	err := app.Listen(3400)
	if err != nil {
		panic(err)
	}
}