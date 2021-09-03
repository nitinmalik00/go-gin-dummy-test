package go_gin_dummy_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// because gin.Default() return pointer(*) to Engine struct type...
var router *gin.Engine

func main () {
	// for setting the router
	router = gin.Default()

	//loading all templates at the server start...so at every request they are not read again...
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html", gin.H{"title":"Home Page"},
			)
	})
	//start the app
	//it will attach router to http.Server and start listening and serving the http requests...
	router.Run()
}



