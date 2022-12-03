package a

import (
	"database/sql"
)

func CallSqlDBMethod(db *sql.DB) {
	db.Begin()              // want `use \(\*sql\.DB\)\.BeginTx instead of \(\*sql\.DB\)\.Begin`
	db.Exec("SELECT 1")     // want `use \(\*sql\.DB\)\.ExecContext instead of \(\*sql\.DB\)\.Exec`
	db.Ping()               // want `use \(\*sql\.DB\)\.PingContext instead of \(\*sql\.DB\)\.Ping`
	db.Prepare("SELECT 1")  // want `use \(\*sql\.DB\)\.PrepareContext instead of \(\*sql\.DB\)\.Prepare`
	db.Query("SELECT 1")    // want `use \(\*sql\.DB\)\.QueryContext instead of \(\*sql\.DB\)\.Query`
	db.QueryRow("SELECT 1") // want `use \(\*sql\.DB\)\.QueryRowContext instead of \(\*sql\.DB\)\.QueryRow`
}

func CallSqlStmtMethod(stmt *sql.Stmt) {
	stmt.Exec("SELECT 1")     // want `use \(\*sql\.Stmt\)\.ExecContext instead of \(\*sql\.Stmt\)\.Exec`
	stmt.Query("SELECT 1")    // want `use \(\*sql\.Stmt\)\.QueryContext instead of \(\*sql\.Stmt\)\.Query`
	stmt.QueryRow("SELECT 1") // want `use \(\*sql\.Stmt\)\.QueryRowContext instead of \(\*sql\.Stmt\)\.QueryRow`
}

func CallSqlTxMethod(tx *sql.Tx, stmt *sql.Stmt) {
	tx.Exec("SELECT 1")     // want `use \(\*sql\.Tx\)\.ExecContext instead of \(\*sql\.Tx\)\.Exec`
	tx.Prepare("SELECT 1")  // want `use \(\*sql\.Tx\)\.PrepareContext instead of \(\*sql\.Tx\)\.Prepare`
	tx.Query("SELECT 1")    // want `use \(\*sql\.Tx\)\.QueryContext instead of \(\*sql\.Tx\)\.Query`
	tx.QueryRow("SELECT 1") // want `use \(\*sql\.Tx\)\.QueryRowContext instead of \(\*sql\.Tx\)\.QueryRow`
	tx.Stmt(stmt)           // want `use \(\*sql\.Tx\)\.StmtContext instead of \(\*sql\.Tx\)\.Stmt`
}
