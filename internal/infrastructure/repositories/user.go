package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/elSyarif/belajar-golang-rest-api/internal/domain"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain/model"
	"github.com/elSyarif/belajar-golang-rest-api/internal/helper"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(connection *sqlx.DB) domain.UserRepository {
	return &userRepository{DB: connection}
}

func (u *userRepository) AddUser(ctx context.Context, user *model.User) error {
	SQL := `INSERT INTO users (id, name, username, email, password) values (?, ?, ?, ?, ?)`

	tx, err := u.DB.Beginx()
	if err != nil {
		panic(err)
	}

	result, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("err mustExecContext: %s", err.Error())
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("add user error : %s", err.Error())
	}

	tx.Commit()
	return nil
}

func (u *userRepository) VerifyUserEmail(ctx context.Context, email string) (string, error) {
	SQL := "SELECT id, email, name FROM users WHERE email = ?"
	user := &model.User{}

	tx, err := u.DB.Beginx()
	helper.PanicError(err)
	defer helper.Exec(tx)

	stmt, err := tx.PreparexContext(ctx, SQL)
	helper.PanicError(err)
	defer stmt.Close()

	err = stmt.GetContext(ctx, &user, email)
	if err != nil {
		return "", fmt.Errorf("error verify email : %s", err.Error())
	}

	return user.Id, nil
}

func (u *userRepository) VerifyUserCredential(ctx context.Context, email string, password string) (string, error) {
	SQL := "SELECT id, email, password FROM users WHERE email = ? AND password = ?"
	user := &model.User{}

	tx, err := u.DB.Beginx()
	helper.PanicError(err)
	defer helper.Exec(tx)

	stmt, err := tx.PreparexContext(ctx, SQL)
	helper.PanicError(err)
	defer stmt.Close()

	err = stmt.GetContext(ctx, &user, email, password)
	if err != nil {
		return "", fmt.Errorf("error verify user credential : %s", err.Error())
	}

	return user.Id, nil
}

func (u *userRepository) VerifyUser(ctx context.Context, userId string) (*model.User, error) {
	SQL := "SELECT id, name, email FROM users WHERE id = ?"
	user := &model.User{}

	tx, err := u.DB.Beginx()
	helper.PanicError(err)
	defer helper.Exec(tx)

	stmt, err := tx.PreparexContext(ctx, SQL)
	helper.PanicError(err)
	defer stmt.Close()

	err = stmt.GetContext(ctx, &user, userId)
	if err != nil {
		return nil, fmt.Errorf("error verfy user : %s", err.Error())
	}

	return user, nil

}

func (u *userRepository) GetUser(ctx context.Context) (*[]model.User, error) {
	SQL := "SELECT * FROM users"
	user := []model.User{}

	tx, err := u.DB.Beginx()
	helper.PanicError(err)
	defer helper.Exec(tx)

	stmt, err := tx.PreparexContext(ctx, SQL)
	helper.PanicError(err)
	defer stmt.Close()

	err = stmt.SelectContext(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("eror get user : %s", err.Error())
	}

	return &user, nil
}
