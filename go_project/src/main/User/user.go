package User

import (
	_ "crypto/des"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go_project/src/main/utils"
	_ "go_project/src/main/utils"
	"net/http"
)

func GetUser() {
	fmt.Println(6556454)
	var data = utils.SelectData("select * from login")
	fmt.Printf("255:%#v\n", data)
	r := gin.Default()
	r.Use(Cors())
	r.GET("/user", func(c *gin.Context) {

	})
	r.POST("/user/login", func(c *gin.Context) {
		var username = c.Query("username")
		var password = c.Query("password")
		for i := 0; i < len(data); i++ {
			if username == data[i].UserName && password == data[i].UserPassword {
				c.JSON(http.StatusOK, gin.H{
					"message": "OK",
				})
				break
			} else if i == len(data)-1 {
				c.JSON(http.StatusOK, gin.H{
					"message": "No",
				})
			}

		}

	})
	r.POST("/user/username", func(c *gin.Context) {
		var username = c.Query("username")
		for i := 0; i < len(data); i++ {
			if username == data[i].UserName {
				c.JSON(http.StatusOK, gin.H{
					"message": "OK",
				})
				break
			} else if i == len(data)-1 {
				c.JSON(http.StatusOK, gin.H{
					"message": "No",
				})
			}

		}

	})
	r.POST("/user/userinset", func(c *gin.Context) {
		var username = c.Query("username")
		var password = c.Query("password")
		var email = c.Query("email")
		var boo = utils.InsertData("insert into login(username,password,email) value(?,?,?)", username, password, email)
		if (boo) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "No",
			})
		}

	})

	r.Run(":9090")
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
