package a

import (
	aliasedsql "database/sql"

	aliasedsqlx "github.com/jmoiron/sqlx"
)

func CallSql(db *aliasedsql.DB) {
	db.Query("SELECT 1") // want `use \(\*aliasedsql\.DB\)\.QueryContext instead of \(\*aliasedsql\.DB\)\.Query`
}

func CallSqlx(db *aliasedsqlx.DB) {
	db.Queryx("SELECT 1") // want `use \(\*aliasedsqlx\.DB\)\.QueryxContext instead of \(\*aliasedsqlx\.DB\)\.Queryx`
}
