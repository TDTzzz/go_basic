package day3_save_query

import "database/sql"

type Engine struct {
	db *sql.DB
	dialect
}
