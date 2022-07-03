package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type Users struct {
	Name  string
	Phone string
	Age   int
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var userData = []Users{
	{Name: "Adam", Phone: "099911111", Age: 23},
	{Name: "Brown", Phone: "0933222222", Age: 22},
	{Name: "Cayla", Phone: "0933333333", Age: 18},
}

func getIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "hello-world-app",
		"Content": "hello world",
	})
}
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", getIndex)
	router.GET("/albums", getAlbums)

	router.Run("localhost:80")
}
