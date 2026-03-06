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
	SQL := "select * from user_m where user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Email, &user.Nama_depan, &user.Nama_belakang, &user.Photo)
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
	row := tx.QueryRowContext(ctx, SQL, email)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select * from user_m"
	rows, err := tx.QueryContext(ctx, SQL)
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
