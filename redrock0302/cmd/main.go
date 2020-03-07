package main

import (
	"github.com/gin-gonic/gin"
	"redrock0302/user"
)

func main()  {
	app := gin.Default()

	user.SetupRouter(app)

	app.Run(":8080")
}