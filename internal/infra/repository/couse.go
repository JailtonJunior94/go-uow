package repository

import (
	"context"
	"database/sql"

	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/infra/db"
	"github.com/jailtonjunior94/go-uow/pkg/observability"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB            *sql.DB
	Queries       *db.Queries
	observability observability.Observability
}

func NewCourseRepository(dtb *sql.DB, observability observability.Observability) *CourseRepository {
	return &CourseRepository{
		DB:            dtb,
		Queries:       db.New(dtb),
		observability: observability,
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	ctx, span := r.observability.Tracer().Start(ctx, "course_repository.insert")
	defer span.End()

	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
	})
}
