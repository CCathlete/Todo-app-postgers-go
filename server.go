package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", indexHandler)

	app.Post("/", postHandler)

	app.Put("/update", putHandler)

	app.Delete("/delete", deleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := app.Listen(fmt.Sprintf(":%v", port))
	log.Fatalln(err)
}

func indexHandler(c *fiber.Ctx) {
	c.SendString("Hemlo")
}

func postHandler(c *fiber.Ctx) {
	c.SendString("Hemlo")
}

func putHandler(c *fiber.Ctx) {
	c.SendString("Hemlo")
}

func deleteHandler(c *fiber.Ctx) {
	c.SendString("Hemlo")
}
