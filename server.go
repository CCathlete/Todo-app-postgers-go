package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	connStr := "postgresql://postgres:postgres@127.0.0.1:5432/todoapp?sslmode=disable"

	engine := html.New("./views", ".html")
	configuration := fiber.Config{
		Views: engine,
	}
	app := fiber.New(configuration)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("./public", "./public")
	err = app.Listen(fmt.Sprintf(":%v", port))
	log.Fatalln(err)
}

func indexHandler(c *fiber.Ctx, db *sql.DB) (err error) {
	var res string
	var todos = []string{}
	rows, err := db.Query("SELECT * FROM todos")
	defer func() {
		err = rows.Close()
	}()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured.")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

type todo struct {
	Item string
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	oldItem := c.Query("olditem")
	newItem := c.Query("newitem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newItem, oldItem)
	return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}
