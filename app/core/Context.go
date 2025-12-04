package core

import (
	"database/sql"
	"os"

	"GO/app/core/database"
)

type Context struct {
	db      *database.Mysql
	secrets *Secrets
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

func (c *Context) GetSecrets() *Secrets {
	if c.secrets == nil {
		botToken := os.Getenv("bot_token")
		secrets := Secrets{BotToken: &botToken}
		c.secrets = &secrets
		xuiHost := os.Getenv("xui_host")
		xuiHash := os.Getenv("xui_hash")
		xuiPort := os.Getenv("xui_port")
		secrets.XuiHost = &xuiHost
		secrets.XuiHash = &xuiHash
		secrets.XuiPort = &xuiPort
	}
	return c.secrets
}
