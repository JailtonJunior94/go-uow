.PHONY: migrate
migrate:
	@migrate create -ext sql -dir database/migrations -format unix $(NAME)

test:
	go test -coverprofile coverage.out -failfast ./...
	go tool cover -func coverage.out | grep total

cover:
	go tool cover -html=coverage.outmak

gqlgen:
	@cd internal/infra && go run github.com/99designs/gqlgen

.PHONY: mockery
mock:
	@mockery --dir=internal/repository --name=CourseRepositoryInterface --filename=course_repository_mock.go --output=internal/repository/mocks --outpkg=repositoriesMock
	@mockery --dir=internal/repository --name=CategoryRepositoryInterface --filename=category_repository_mock.go --output=internal/repository/mocks --outpkg=repositoriesMock
