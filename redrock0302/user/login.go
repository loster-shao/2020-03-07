package user

import (
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"redrock0302/token"
	"strconv"
)

//type LoginForm struct {
//	ID       int    `json:"id"       binding:"required"`
//	Username string `json:"username" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}

func Login(c *gin.Context){
	id0     := c.PostForm("user_id")
	id, err := strconv.Atoi(id0)
	if  err != nil {
		fmt.Println("err:",err)
	}
	password := c.PostForm("password")
	db, err1 := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if  err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	defer db.Close()
	var u User
	db.Where("id=?", id).First(&u)
	token := token.Create(u.Username, id)//创建token
	fmt.Println(u)
	fmt.Println(id)
	fmt.Println(password)
	fmt.Println(token)
	if password == u.Password {
		c.JSON(200,gin.H{"status:":http.StatusOK,"message":"登录成功"})
	}else {
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户名或密码错误"})
	}
}



