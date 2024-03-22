package usecase

import (
	"context"

	"github.com/jailtonjunior94/go-uow/internal/entity"
	"github.com/jailtonjunior94/go-uow/internal/infra/repository"
	"github.com/jailtonjunior94/go-uow/pkg/database/uow"
	"github.com/jailtonjunior94/go-uow/pkg/logger"
	"github.com/jailtonjunior94/go-uow/pkg/observability"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/metric"
)

type AddCourseUowUseCase struct {
	logger        logger.Logger
	uow           uow.UowInterface
	observability observability.Observability
	metrics       metrics
}

type metrics struct {
	meter    metric.Meter
	counter  metric.Int64Counter
	duration metric.Float64Histogram
}

func NewAddCourseUowUseCase(
	logger logger.Logger,
	uow uow.UowInterface,
	observability observability.Observability,
) *AddCourseUowUseCase {

	meter := observability.MeterProvider().Meter("tramporteiro")

	counter, err := meter.Int64Counter("number.registered.courses", metric.WithDescription("number of registered courses"))
	if err != nil {
		logger.Error(err)
	}

	duration, err := meter.Float64Histogram("course.registration.time", metric.WithDescription("course registration time"))
	if err != nil {
		logger.Error(err)
	}

	metrics := metrics{
		meter:    meter,
		counter:  counter,
		duration: duration,
	}

	return &AddCourseUowUseCase{
		logger:        logger,
		uow:           uow,
		observability: observability,
		metrics:       metrics,
	}
}

func (a *AddCourseUowUseCase) Execute(ctx context.Context, categoryParam *CategoryParams, courseParam *CourseParams) error {
	return a.uow.Do(ctx, func(uow uow.UowInterface) error {
		ctx, span := a.observability.Tracer().Start(ctx, "add_couse_unit_of_work.execute")
		defer span.End()
		a.metrics.counter.Add(ctx, 1)

		courseRepository := a.getCourseRepository(ctx)
		categoryRepository := a.getCategoryRepository(ctx)

		category := entity.Category{
			ID:          uuid.New().String(),
			Name:        categoryParam.Name,
			Description: categoryParam.Description,
		}

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
