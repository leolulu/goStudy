package webserver

import (
	"fmt"
	"ginStudy/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, util.GenData())
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/", func(c *gin.Context) {
		c.BindJSON()
		c.String(http.StatusOK, "index page")
	})
	return router
}

func init() {
	fmt.Println("gin set mode")
	gin.SetMode("debug")
}
