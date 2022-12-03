package sqlpassctxcheck

type SqlPackage string

const (
	SQL  SqlPackage = "database/sql"
	SQLX SqlPackage = "github.com/jmoiron/sqlx"
)

var allSqlPackage = []SqlPackage{
	SQL,
	SQLX,
}

var SqlModuleImportPathMap = map[SqlPackage]string{
	SQL:  `"database/sql"`,
	SQLX: `"github.com/jmoiron/sqlx"`,
}
