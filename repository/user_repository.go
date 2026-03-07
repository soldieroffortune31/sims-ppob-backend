package repository

import (
	"context"
	"database/sql"
	"sims-ppob/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Count(ctx context.Context, tx *sql.Tx) int
	IsEmailExist(ctx context.Context, tx *sql.Tx, email string) (bool, error)
	IsEmailExistByIdAndEmail(ctx context.Context, tx *sql.Tx, userId int, email string) (bool, error)
	FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []domain.User
	Login(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	UpdateToken(ctx context.Context, tx *sql.Tx, userId int, token string)
}
