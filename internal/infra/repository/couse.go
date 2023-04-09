package repository

import (
	"context"
	"database/sql"

	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/infra/db"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
	})
}
