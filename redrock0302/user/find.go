package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Users struct {
	Username string
	Id       int
}

func Find(c *gin.Context)  {
	user := c.Query("username")
	//fmt.Println(err)
	fmt.Println(user)
	db, err1 := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err1 != nil {
		fmt.Println("err:", err1)
		return
	}
	var u Users
	db.Where("username=?",user).First(&u)
	fmt.Println(u)
	c.JSON(200,gin.H{"status":http.StatusOK,"message":u})
}