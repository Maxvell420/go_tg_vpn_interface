package core

import (
	"database/sql"
	"os"

	"GO/app/core/database"
)

type Context struct {
	db *database.Mysql
}

func (c *Context) GetDb() *sql.DB {
	if c.db == nil {
		pass := os.Getenv("db_pass")
		user := os.Getenv("db_user")
		db_name := os.Getenv("db_name")
		port := os.Getenv("db_port")
		host := os.Getenv("db_host")
		db := database.Mysql{Host: host, Port: port, User: user, Dbname: db_name, Pass: pass}
		c.db = &db
	}
	return c.db.GetDb()
}
