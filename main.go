package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mystery struct {
	ID         string  `json:"id"`
	Name       string  `json:"Name"`
	Role       string  `json:"Role"`
	Popularity float64 `json:"Popularity"`
}

var Gravity_falls = []Mystery{
	{ID: "1", Name: "Dipper Pines", Role: "Protagonist", Popularity: 56.99},
	{ID: "2", Name: "Mabel Pines", Role: "Protagonist", Popularity: 17.99},
	{ID: "3", Name: "Bill Cipher", Role: "Antagonist", Popularity: 40.99},
}

func getMysterys(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Gravity_falls)
}
func main() {
	router := gin.Default()
	router.GET("/Mysterys", getMysterys)

	router.Run("localhost:8080")
}
