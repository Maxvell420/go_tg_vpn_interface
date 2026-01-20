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
	db := Context.GetDb()

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

	runSql(sql, db)

	sql = `CREATE TABLE IF NOT EXISTS
  		referal_links (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    hash varchar(255) NOT NULL UNIQUE,
    tg_id bigint unsigned NOT NULL
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE IF NOT EXISTS
  referal_users (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    tg_id bigint unsigned NOT NULL,
    owner_tg_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE
  vpn_clients (
    id int NOT NULL,
    enabled enum('yes', 'no') NOT NULL,
    total int NOT NULL,
    remaining bigint NOT NULL,
    lastOnline int NOT NULL,
    uuid varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    inbound_id int NOT NULL,
    user_id bigint unsigned NOT NULL,
    timestamp_expire int NOT NULL
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE
  inbounds (
    id int NOT NULL,
    total bigint unsigned NOT NULL,
    calc_total bigint unsigned NOT NULL,
    protocol varchar(255) NOT NULL,
    tag varchar(255) NOT NULL
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE
  sales (
    id int unsigned NOT NULL AUTO_INCREMENT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    value double(8, 2) NOT NULL,
    product_type int unsigned NOT NULL,
    product_id int DEFAULT NULL,
    user_id bigint unsigned DEFAULT NULL,
    status int unsigned DEFAULT NULL,
    PRIMARY KEY (id)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE
  vpn_traffic_products (
    id int unsigned NOT NULL AUTO_INCREMENT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title varchar(255) NOT NULL,
    cost double(8, 2) NOT NULL,
    traffic bigint DEFAULT NULL,
    time_amount bigint unsigned DEFAULT NULL,
    PRIMARY KEY (id)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)

	sql = `CREATE TABLE
  vpn_traffic_usage (
    id int unsigned NOT NULL AUTO_INCREMENT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    client_id int unsigned NOT NULL,
    traffic bigint unsigned NOT NULL,
    PRIMARY KEY (id)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci`

	runSql(sql, db)
}

func runSql(sql string, db *sql.DB) {
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	fmt.Println("Все четко")
}
