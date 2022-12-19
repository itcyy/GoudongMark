package target

type MySql struct {
	Username string "username"
	Password string "password"
	Url      string "127.0.0.1:3306"
}

func SqlValue() MySql {
	var sqlStr MySql
	return sqlStr
}
