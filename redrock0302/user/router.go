package user

import (
	"github.com/gin-gonic/gin"
)

// router.go something

func SetupRouter(app *gin.Engine) {
	app.POST("/register", Register)
	app.POST("/login", Login)
	app.POST("/xiugai", Xiugai)
	app.GET("/find", Find)
	//没啥用
	// app.GET("/test", middleware.User,func(c *gin.Context) {
	//	fmt.Println(c.Get("uid"))
	// })
}
