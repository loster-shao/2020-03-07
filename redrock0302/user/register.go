package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type User struct {
	gorm.Model
	num      int    //`gorm:"AUTO_INCREMENT;unique;not null"`
	Username string //`gorm:"-"`  //`json:"username" bind"required"`
	Password string //`gorm:"-"`  //`json:"password" bind"required"`
}

func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//var user User
	//err := c.ShouldBindJSON(&user);
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Create(&User{Username:username,Password:password})
	c.JSON(200,gin.H{"status":http.StatusOK,"message":"注册成功"})
}

