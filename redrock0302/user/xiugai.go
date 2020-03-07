package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"redrock0302/token"
)

func Xiugai(c *gin.Context)  {
	users    := c.PostForm("username")
	password := c.PostForm("password")
	tokens := c.GetHeader("token")
	fmt.Println(tokens)
	id, username, err := token.CheckToken(tokens)
	if err != nil{
		fmt.Println("err:",err)
		c.JSON(500,gin.H{"status":http.StatusBadRequest,"message":"token验证失败"})
		return
	}
	db, err1 := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	defer db.Close()
	var u User
	db.Where("id=?", id).First(&u)
	fmt.Println(id, u)
	if username == u.Username {
		fmt.Println("OK")
		db.Model(User{}).Where("id=?", id).Update(User{Password: password, Username: users})
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"修改成功"})
		token.Create(username, id)
		return
	}else {
		fmt.Println("err",u.Username)
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户不存在"})
		return
	}
}
