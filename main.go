package main

import (
	"fmt"

	movies "github.com/Danex2/go-movies-api/controllers"
	"github.com/Danex2/go-movies-api/database"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()

	app.Settings.ServerHeader = "Fiber"

	app.Use(logger.New())

	var err error

	database.DBConn, err = gorm.Open("sqlite3", "movies.db")

	if err != nil {
		panic("There was an error opening the database!")
	}

	fmt.Println("Connected to databse successfully!")

	database.DBConn.AutoMigrate(&movies.Movie{})

	defer database.DBConn.Close()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("This is my golang API")
	})

	app.Post("/movies", movies.CreateMovie)
	app.Get("/movies", movies.GetMovies)
	app.Get("/movies/:id", movies.GetMovie)

	app.Listen(4000)
}
