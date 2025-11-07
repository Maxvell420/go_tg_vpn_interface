package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Host   string
	Port   string
	User   string
	Dbname string
	Pass   string
	db     *sql.DB
}

func (q *Mysql) GetDb() *sql.DB {
	if q.db == nil {
		db, err := sql.Open("mysql", q.User+":"+q.Pass+"@tcp"+"("+q.Host+":"+q.Port+")"+"/"+q.Dbname)
		if err != nil {
			// Добавить обработку
			fmt.Println(err)
			os.Exit(1)
		}
		q.db = db
	}
	return q.db
}
