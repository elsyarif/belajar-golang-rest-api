package helper

import "github.com/jmoiron/sqlx"

func Exec(tx *sqlx.Tx) {
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
