package repository

import (
	"context"
	"database/sql"
	"errors"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Login implements [UserRepository].
func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT user_id, email, password FROM user_m where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// UpdateToken implements [UserRepository].
func (repository *UserRepositoryImpl) UpdateToken(ctx context.Context, tx *sql.Tx, userId int, token string) {
	SQL := "UPDATE user_m set token = ? where user_id = ?"
	result, err := tx.ExecContext(ctx, SQL, token, userId)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user_m(email, nama_depan, nama_belakang, photo, password, token) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Nama_depan, user.Nama_belakang, user.Photo, user.Password, user.Token)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.User_id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update user_m set email = ?, nama_depan = ?, nama_belakang = ?, photo = ? where user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Nama_depan, user.Nama_belakang, user.Photo, user.User_id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from user_m where user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.User_id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select a.user_id, a.email, a.nama_depan, a.nama_belakang, a.photo, b.balance from user_m a join userbalance_m b on a.user_id = b.user_id where a.user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Email, &user.Nama_depan, &user.Nama_belakang, &user.Photo, &user.UserBalance.Balance)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) IsEmailExist(ctx context.Context, tx *sql.Tx, email string) (bool, error) {
	SQL := "select count(*) from user_m where email = ?"
	row := tx.QueryRowContext(ctx, SQL, email)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository *UserRepositoryImpl) IsEmailExistByIdAndEmail(ctx context.Context, tx *sql.Tx, userId int, email string) (bool, error) {
	SQL := "select count(*) from user_m where user_id != ? and email = ?"
	row := tx.QueryRowContext(ctx, SQL, userId, email)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Count implements [UserRepository].
func (repository *UserRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) FROM user_m"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	total := 0

	if rows.Next() {
		err := rows.Scan(&total)
		helper.PanicIfError(err)
	}

	return total
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []domain.User {
	SQL := "SELECT user_id, email, nama_depan, nama_belakang, photo FROM user_m LIMIT ? OFFSET ?"
	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.User_id, &user.Email, &user.Nama_depan, &user.Nama_belakang, &user.Photo)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}
