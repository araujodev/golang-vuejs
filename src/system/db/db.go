package db

import (
	_ "github.com/go-sql-driver/mysql" 
	"github.com/go-xorm/xorm"
)

func Connect() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "root:123456m@tcp(localhost:3306)/gotest?charset=utf8")
}
