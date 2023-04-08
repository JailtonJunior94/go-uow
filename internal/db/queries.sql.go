// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)
`

type CreateCategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name, arg.Description)
	return err
}

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, price, category_id) VALUES ($1, $2, $3, $4, $5)
`

type CreateCourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       string
	CategoryID  string
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.CategoryID,
	)
	return err
}