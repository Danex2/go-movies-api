package main

import (
	"net/http/httptest"
	"testing"

	movies "github.com/Danex2/go-movies-api/controllers"
	"github.com/Danex2/go-movies-api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func dbSetup() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "movies_test.db")

	if err != nil {
		panic("There was an error opening the database!")
	}

	database.DBConn.AutoMigrate(&movies.Movie{})

	defer database.DBConn.Close()
}

func testStatus200(t *testing.T, app *App, url string, method string) {
	req := httptest.NewRequest(method, url, nil)

	resp, err := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode, "Status code")
}

func TestAllMovies(t *testing.T) {
	app := fiber.New()

	dbSetup()

	app.Get("/movies", movies.GetMovies)

	req := httptest.NewRequest("GET", "/movies", nil)

	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode, "Status Code")
}

func TestSingleMovie(t *testing.T) {
	app := fiber.New()

	dbSetup()

	app.Get("/movies/:id", movies.GetMovie)

	req := httptest.NewRequest("GET", "/movies/:id", nil)

	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode, "Status Code")
}

func TestCreateMovie(t *testing.T) {
	app := fiber.New()

	dbSetup()

	func dummyHandler(c *fiber.Ctx) {
		db := database.DBConn

		m := &Movie{}
	
		m.Name = "fast and furious"
		m.Year = 2001
	
		db.Create(m)
		c.JSON(m)
	}

	app.Post("/movies", dummyHandler)

	testStatus200(t, app, "/movies", "POST")
}
