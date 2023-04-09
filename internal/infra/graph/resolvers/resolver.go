package resolvers

import "github.com/jailtonjunior94/go-uow/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AddCourse    *usecase.AddCourseUseCase
	AddCourseUow *usecase.AddCourseUowUseCase
}
