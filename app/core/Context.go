package core

import (
	"database/sql"
	"os"

	"GO/app/core/database"
	"GO/app/libs/3xui"
	"GO/app/libs/telegram"
)

type Context struct {
	db              *database.Mysql
	xuiRequest      *xui.Request
	telegramRequest *telegram.Request
	secrets         *Secrets
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
		xuiHost := os.Getenv("xui_host")
		xuiHash := os.Getenv("xui_hash")
		xuiPort := os.Getenv("xui_port")
		xuiUser := os.Getenv("xui_user")
		xuiPass := os.Getenv("xui_pass")
		secrets := Secrets{BotToken: &botToken, XuiHost: &xuiHost, XuiHash: &xuiHash, XuiPort: &xuiPort, XuiUser: &xuiUser, XuiPass: &xuiPass}
		c.secrets = &secrets
	}
	return c.secrets
}

func (c *Context) Get3xuiRequest() *xui.Request {
	if c.xuiRequest == nil {
		if c.secrets == nil {
			c.GetSecrets()
		}
		c.xuiRequest = &xui.Request{Host: *c.secrets.XuiHost, Hash: *c.secrets.XuiHash, Port: *c.secrets.XuiPort, XuiUser: *c.secrets.XuiUser, XuiPass: *c.secrets.XuiPass}
	}
	return c.xuiRequest
}

func (c *Context) GetTelegramRequest() *telegram.Request {
	if c.telegramRequest == nil {
		if c.secrets == nil {
			c.GetSecrets()
		}
		c.telegramRequest = &telegram.Request{BotToken: c.secrets.BotToken}
	}
	return c.telegramRequest
}
