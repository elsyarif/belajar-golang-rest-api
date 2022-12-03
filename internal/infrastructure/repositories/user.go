package repositories

import (
	"context"
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
	SQL := "INSERT INTO users (id, name, username, email, password) values (?, ?, ?, ?, ?)"

	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.Exec(tx)

	_, err = tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Username, user.Email, user.Password)
	helper.PanicError(err)

	return nil
}

func (u *userRepository) VerifyUserEmail(ctx context.Context, email string) (string, error) {
	SQL := "SELECT id, email, FROM users WHERE email = ?"
	user := &model.User{}

	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.Exec(tx)

	err = tx.QueryRowContext(ctx, SQL, email).Scan(user.Id, user.Email)
	helper.PanicError(err)

	return user.Id, nil
}

func (u *userRepository) VerifyUserCredential(ctx context.Context, email string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) VerifyUser(ctx context.Context, userId string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetUser(ctx context.Context, userId string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
