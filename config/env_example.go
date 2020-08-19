package config

import (
	"os"
)

//Change you're own database configuration here
const (
	DB_USER     = "root"
	DB_PASSWORD = ""
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "go_api"
)

//And it will be automatically set as env
func init() {
	os.Setenv("DB_USER", DB_USER)
	os.Setenv("DB_PASSWORD", DB_PASSWORD)
	os.Setenv("DB_HOST", DB_HOST)
	os.Setenv("DB_PORT", DB_PORT)
	os.Setenv("DB_NAME", DB_NAME)
}
