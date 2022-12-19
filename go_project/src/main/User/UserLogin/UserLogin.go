package UserLogin

import (
	"fmt"
	"go_project/src/main/utils"
)

func Login() {
	fmt.Println(6556454)
	var data = utils.SelectData("select * from login where username = \"root\" and password = \"123456\"")
	fmt.Printf("255:%#v\n", data)

}