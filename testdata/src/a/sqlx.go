package a

import "github.com/jmoiron/sqlx"

func CallSqlxDBMethod(db *sqlx.DB) {
	db.Beginx()                    // want `use \(\*sqlx\.DB\)\.BeginTxx instead of \(\*sqlx\.DB\)\.Beginx`
	db.Get(nil, "SELECT 1")        // want `use \(\*sqlx\.DB\)\.GetContext instead of \(\*sqlx\.DB\)\.Get`
	db.MustBegin()                 // want `use \(\*sqlx\.DB\)\.MustBeginTx instead of \(\*sqlx\.DB\)\.MustBegin`
	db.MustExec("SELECT 1")        // want `use \(\*sqlx\.DB\)\.MustExecContext instead of \(\*sqlx\.DB\)\.MustExec`
	db.NamedExec("SELECT 1", nil)  // want `use \(\*sqlx\.DB\)\.NamedExecContext instead of \(\*sqlx\.DB\)\.NamedExec`
	db.NamedQuery("SELECT 1", nil) // want `use \(\*sqlx\.DB\)\.NamedQueryContext instead of \(\*sqlx\.DB\)\.NamedQuery`
	db.PrepareNamed("SELECT 1")    // want `use \(\*sqlx\.DB\)\.PrepareNamedContext instead of \(\*sqlx\.DB\)\.PrepareNamed`
	db.Preparex("SELECT 1")        // want `use \(\*sqlx\.DB\)\.PreparexContext instead of \(\*sqlx\.DB\)\.Preparex`
	db.QueryRowx("SELECT 1")       // want `use \(\*sqlx\.DB\)\.QueryRowxContext instead of \(\*sqlx\.DB\)\.QueryRowx`
	db.Queryx("SELECT 1")          // want `use \(\*sqlx\.DB\)\.QueryxContext instead of \(\*sqlx\.DB\)\.Queryx`
	db.Select(nil, "SELECT 1")     // want `use \(\*sqlx\.DB\)\.SelectContext instead of \(\*sqlx\.DB\)\.Select`
}

func CallSqlxNamedStmtMethod(stmt *sqlx.NamedStmt) {
	stmt.Exec(nil)        // want `use \(\*sqlx\.NamedStmt\)\.ExecContext instead of \(\*sqlx\.NamedStmt\)\.Exec`
	stmt.Get(nil, nil)    // want `use \(\*sqlx\.NamedStmt\)\.GetContext instead of \(\*sqlx\.NamedStmt\)\.Get`
	stmt.MustExec(nil)    // want `use \(\*sqlx\.NamedStmt\)\.MustExecContext instead of \(\*sqlx\.NamedStmt\)\.MustExec`
	stmt.Query(nil)       // want `use \(\*sqlx\.NamedStmt\)\.QueryContext instead of \(\*sqlx\.NamedStmt\)\.Query`
	stmt.QueryRow(nil)    // want `use \(\*sqlx\.NamedStmt\)\.QueryRowContext instead of \(\*sqlx\.NamedStmt\)\.QueryRow`
	stmt.QueryRowx(nil)   // want `use \(\*sqlx\.NamedStmt\)\.QueryRowxContext instead of \(\*sqlx\.NamedStmt\)\.QueryRowx`
	stmt.Queryx(nil)      // want `use \(\*sqlx\.NamedStmt\)\.QueryxContext instead of \(\*sqlx\.NamedStmt\)\.Queryx`
	stmt.Select(nil, nil) // want `use \(\*sqlx\.NamedStmt\)\.SelectContext instead of \(\*sqlx\.NamedStmt\)\.Select`
}

func CallSqlxStmtMethod(stmt *sqlx.Stmt) {
	stmt.Get(nil)    // want `use \(\*sqlx\.Stmt\)\.GetContext instead of \(\*sqlx\.Stmt\)\.Get`
	stmt.MustExec()  // want `use \(\*sqlx\.Stmt\)\.MustExecContext instead of \(\*sqlx\.Stmt\)\.MustExec`
	stmt.QueryRowx() // want `use \(\*sqlx\.Stmt\)\.QueryRowxContext instead of \(\*sqlx\.Stmt\)\.QueryRowx`
	stmt.Queryx()    // want `use \(\*sqlx\.Stmt\)\.QueryxContext instead of \(\*sqlx\.Stmt\)\.Queryx`
	stmt.Select(nil) // want `use \(\*sqlx\.Stmt\)\.SelectContext instead of \(\*sqlx\.Stmt\)\.Select`
}

func CallSqlxTxMethod(tx *sqlx.Tx, nstmt *sqlx.NamedStmt) {
	tx.Get(nil, "SELECT 1")       // want `use \(\*sqlx\.Tx\)\.GetContext instead of \(\*sqlx\.Tx\)\.Get`
	tx.MustExec("SELECT 1")       // want `use \(\*sqlx\.Tx\)\.MustExecContext instead of \(\*sqlx\.Tx\)\.MustExec`
	tx.NamedExec("SELECT 1", nil) // want `use \(\*sqlx\.Tx\)\.NamedExecContext instead of \(\*sqlx\.Tx\)\.NamedExec`
	tx.NamedStmt(nstmt)           // want `use \(\*sqlx\.Tx\)\.NamedStmtContext instead of \(\*sqlx\.Tx\)\.NamedStmt`
	tx.PrepareNamed("SELECT 1")   // want `use \(\*sqlx\.Tx\)\.PrepareNamedContext instead of \(\*sqlx\.Tx\)\.PrepareNamed`
	tx.Preparex("SELECT 1")       // want `use \(\*sqlx\.Tx\)\.PreparexContext instead of \(\*sqlx\.Tx\)\.Preparex`
	tx.QueryRowx("SELECT 1")      // want `use \(\*sqlx\.Tx\)\.QueryRowxContext instead of \(\*sqlx\.Tx\)\.QueryRowx`
	tx.Queryx("SELECT 1")         // want `use \(\*sqlx\.Tx\)\.QueryxContext instead of \(\*sqlx\.Tx\)\.Queryx`
	tx.Select(nil, "SELECT 1")    // want `use \(\*sqlx\.Tx\)\.SelectContext instead of \(\*sqlx\.Tx\)\.Select`
	tx.Stmtx(nil)                 // want `use \(\*sqlx\.Tx\)\.StmtxContext instead of \(\*sqlx\.Tx\)\.Stmtx`
}
