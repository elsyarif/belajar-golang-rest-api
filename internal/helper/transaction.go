package helper

import "database/sql"

func Exec(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRoleback := tx.Rollback()
		PanicError(errorRoleback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicError(errorCommit)
	}
}
