package main

import (
	"phototutor/backend/controller"
	"phototutor/backend/models"
	"phototutor/backend/util"

	"github.com/gin-gonic/gin"
)

func setUpEnv() {

	util.SetUp()
	models.Setup()
}

func main() {
	setUpEnv()

	server := gin.Default()
	// server.Use(TlsHandler())

	// server.Static("img/", util.ImgStaticPrefix)

	// server.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })
	picRoute := server.Group("/")
	{
		controller.NewPictureController(picRoute)
		controller.NewImgController(picRoute)
	}
	// controller.NewUserController(server.Group("/users/"))
	err := server.Run("0.0.0.0:8081")
	//running in tls
	// err := server.RunTLS(":8080", "ssl/rootCA.pem", "ssl/rootCA.key")
	if err != nil {
		println(err.Error())
	}
}

// func TlsHandler() gin.HandlerFunc {
// 	secureMiddleware := secure.New(secure.Options{
// 		SSLRedirect: true,
// 		SSLHost:     "127.0.0.1:8080",
// 	})
// 	return func(c *gin.Context) {
// 		err := secureMiddleware.Process(c.Writer, c.Request)

// 		// If there was an error, do not continue.
// 		if err != nil {
// 			return
// 		}

// 		c.Next()
// 	}
// }