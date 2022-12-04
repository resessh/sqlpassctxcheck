package a

import "database/sql"

func main() {
	// don't panic with static method call
	sql.Open("mysql", "user:password@/dbname")
}
