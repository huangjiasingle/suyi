package db

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func Init(dsn string, debug bool) error {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	Engine = engine
	return engine.Ping()
}
