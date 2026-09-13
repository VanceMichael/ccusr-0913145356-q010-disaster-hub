package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// OpenSQLite 创建可供事件日志和待投递表共享的数据库连接。
func OpenSQLite(path string) (*sql.DB, error) {
	return sql.Open("sqlite", path)
}
