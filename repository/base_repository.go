package repository

import (
	"context"
	"database/sql"
	"sims-ppob/helper"
)

type BaseRepository struct {
	DB *sql.DB
}

func (r *BaseRepository) QueryWithPagination(
	ctx context.Context,
	exec helper.DBExecutor,
	baseQuery string,
	countQuery string,
	where string,
	args []interface{},
	limit int,
	offset int,
) (*sql.Rows, int) {

	// total count
	var total int
	err := exec.QueryRowContext(ctx, countQuery+where, args...).Scan(&total)
	helper.PanicIfError(err)

	// data query
	query := baseQuery + where + " LIMIT ? OFFSET ?"
	argsWithPagination := append(args, limit, offset)

	rows, err := exec.QueryContext(ctx, query, argsWithPagination...)
	helper.PanicIfError(err)

	return rows, total
}
