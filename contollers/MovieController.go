package movies

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Name string `json:"name"`
	Year int16  `json:"year"`
}

func CreateMovie(c *fiberCtx) {
	body := c.Body()

	c.Send(body)
}
