package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/repository"
	"github.com/jailtonjunior94/go-uow/pkg/logger"
)

type (
	CategoryParams struct {
		ID          string
		Name        string
		Description string
	}

	CourseParams struct {
		ID          string
		Name        string
		Description string
	}
)

type AddCourseUseCase struct {
	logger             logger.Logger
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(logger logger.Logger, courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		logger:             logger,
		CourseRepository:   courseRepository,
		CategoryRepository: categoryRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, categoryParam *CategoryParams, courseParam *CourseParams) error {
	category := entity.Category{
		ID:          uuid.New().String(),
		Name:        categoryParam.Name,
		Description: categoryParam.Description,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		a.logger.Error(err)
		return err
	}

	course := entity.Course{
		ID:          uuid.New().String(),
		Name:        courseParam.Name,
		Description: courseParam.Description,
		CategoryID:  category.ID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		a.logger.Error(err)
		return err
	}
	return nil
}
