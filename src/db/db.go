package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	SQL_SUCCESS int = 0
	SQL_ERROR   int = 500

	SELECT_NO_RESULT int = 204

	INSERT_NO_CREATE int = 200
	INSERT_DUPLICATE int = 400

	UPDATE_NO_CHANGE  int = 200
	UPDATE_NO_COLLAPS int = 409

	DELETE_NO_CHANGE int = 204
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

func CheckErr(err error) int {
	if err != nil {
		// todo -  err 값에 따른 에러 코드 세분화  필요
		return SQL_ERROR
	} else {
		return SQL_SUCCESS
	}
}

func CheckResult(nRow int64, dbResultCode int) int {
	if nRow != int64(0) {
		return SQL_SUCCESS
	} else {
		return dbResultCode
	}
}
