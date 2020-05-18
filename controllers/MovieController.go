package movies

import (
	"log"

	"github.com/Danex2/go-movies-api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Movie struct {
	gorm.Model
	Name string `json:"name"`
	Year int16  `json:"year"`
}

func CreateMovie(c *fiber.Ctx) {

	db := database.DBConn

	m := &Movie{}

	if err := c.BodyParser(m); err != nil {
		log.Fatal(err)
	}

	db.Create(m)
	c.JSON(m)

}

func GetMovies(c *fiber.Ctx) {

	db := database.DBConn

	var movies []Movie

	db.Find(&movies)

	c.JSON(movies)
}

func GetMovie(c *fiber.Ctx) {

	db := database.DBConn

	id := c.Params("id")

	m := Movie{}

	db.First(&m, id)

	c.JSON(m)
}
