package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	connStr := "postgresql://postgres:postgres@127.0.0.1/5432/todoapp?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) {
		postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) {
		putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) {
		deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(fmt.Sprintf(":%v", port))
	log.Fatalln(err)
}

func indexHandler(c *fiber.Ctx, db *sql.DB) {
	c.SendString("Hemlo")
}

func postHandler(c *fiber.Ctx, db *sql.DB) {
	c.SendString("Hemlo")
}

func putHandler(c *fiber.Ctx, db *sql.DB) {
	c.SendString("Hemlo")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) {
	c.SendString("Hemlo")
}
