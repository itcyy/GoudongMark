package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_project/src/main/User"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	User.GetUser()

}
