package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	//app.Get("/", func(c *fiber.Ctx){
	//	c.SendString("I'm a GET request!")
	//})

	app.Post("/", func(c *fiber.Ctx){
		c.SendString("I'm a POST request!")
	})

	//app.Group()

	app.Get("/", func(c *fiber.Ctx){
		fmt.Println("1st route!")
		c.Next()
	})

	app.Get("*", func(c *fiber.Ctx){
		fmt.Println("2nd route!")
		c.Next()
	})

	app.Get("/", func(c *fiber.Ctx){
		fmt.Println("3rd route!")
		c.SendString("Hello, World!")
	})

	err := app.Listen(3300)
	if err != nil {
		panic(err)
	}
}
