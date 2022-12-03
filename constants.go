package sqlpassctxcheck

type sqlPackage string

const (
	sql  sqlPackage = "database/sql"
	sqlx sqlPackage = "github.com/jmoiron/sqlx"
)

var allSqlPackage = []sqlPackage{
	sql,
	sqlx,
}

var sqlModuleImportPathMap = map[sqlPackage]string{
	sql:  `"database/sql"`,
	sqlx: `"github.com/jmoiron/sqlx"`,
}
