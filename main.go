package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/Danex2/go-movies-api/controllers/movies"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open("sqlite3", "movies.db")
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("This is my golang API")
	})

	app.Post("/test", movies.CreateMovie)

	app.Listen(4000)
}
