package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzuh/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err = gorm.Open()("sql", "b09s4itual7ujfpk5s5i:pscale_pw_9EjMOcHRQ8GMgsrPfCCkCHr0eP305ZGzO1nMNifxkqT@tcp(aws.connect.psdb.cloud)/blog-app?tls=true&interpolateParams=true")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
