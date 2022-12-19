package target

type MySql struct {
	Username string "root"
	Password string "chenyonyu1@"
	Url      string "127.0.0.1:3306"
}

func SqlValue() MySql {
	var sqlStr MySql
	return sqlStr
}
