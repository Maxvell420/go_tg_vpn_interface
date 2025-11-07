package main

import (
	"database/sql"
	"fmt"
	"os"

	"GO/app/core"

	"github.com/joho/godotenv"
)

func main() {
	var Context core.Context
	godotenv.Load("../.env")
	sql := `CREATE TABLE IF NOT EXISTS
  users (
    id bigint unsigned NOT NULL AUTO_INCREMENT,
    tg_id bigint unsigned NOT NULL,
    user_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    created_at timestamp NULL DEFAULT NULL,
    updated_at timestamp NULL DEFAULT NULL,
    kicked enum('yes', 'no') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'no',
    is_admin enum('yes', 'no') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'no',
    PRIMARY KEY (id)
  ) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	db := Context.GetDb()
	runSql(sql, db)
}

func runSql(sql string, db *sql.DB) {
	db.Exec(sql)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
}
