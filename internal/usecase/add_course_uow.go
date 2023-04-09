package usecase

import (
	"context"

	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/infra/repository"
	"github.com/jailtonjunior94/go-uow/pkg/database/uow"
	"github.com/jailtonjunior94/go-uow/pkg/logger"

	"github.com/google/uuid"
)

type AddCourseUowUseCase struct {
	logger logger.Logger
	uow    uow.UowInterface
}

func NewAddCourseUowUseCase(logger logger.Logger, uow uow.UowInterface) *AddCourseUowUseCase {
	return &AddCourseUowUseCase{
		logger: logger,
		uow:    uow,
	}
}

func (a *AddCourseUowUseCase) Execute(ctx context.Context, categoryParam *CategoryParams, courseParam *CourseParams) error {
	return a.uow.Do(ctx, func(uow uow.UowInterface) error {
		category := entity.Category{
			ID:          uuid.New().String(),
			Name:        categoryParam.Name,
			Description: categoryParam.Description,
		}

		categoryRepository := a.getCategoryRepository(ctx)
		err := categoryRepository.Insert(ctx, category)
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

		courseRepository := a.getCourseRepository(ctx)
		err = courseRepository.Insert(ctx, course)
		if err != nil {
			a.logger.Error(err)
			return err
		}
		return nil
	})
}

func (a *AddCourseUowUseCase) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		a.logger.Error(err)
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}

func (a *AddCourseUowUseCase) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		a.logger.Error(err)
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}
