package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetConnector() *sql.DB {
	cfg := mysql.Config{
		User:                 "dhlee",
		Passwd:               "happy3119",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               "study",
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}
