package repository

import (
	"context"
	"database/sql"

	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/infra/db"
	"github.com/jailtonjunior94/go-uow/pkg/observability"
	"go.opentelemetry.io/otel/attribute"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB            *sql.DB
	Queries       *db.Queries
	observability observability.Observability
}

func NewCategoryRepository(dtb *sql.DB, observability observability.Observability) *CategoryRepository {
	return &CategoryRepository{
		DB:            dtb,
		Queries:       db.New(dtb),
		observability: observability,
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	ctx, span := r.observability.Tracer().Start(ctx, "category_repository.insert")
	defer span.End()

	span.SetAttributes(
		attribute.String("category.id", category.ID),
		attribute.String("category.name", category.Name),
		attribute.String("category.description", category.Description),
	)

	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	})
}
