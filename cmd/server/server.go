package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/jailtonjunior94/go-uow/configs"

	"github.com/jailtonjunior94/go-uow/internal/infra/db"
	"github.com/jailtonjunior94/go-uow/internal/infra/graph/generated"
	"github.com/jailtonjunior94/go-uow/internal/infra/graph/resolvers"
	"github.com/jailtonjunior94/go-uow/internal/infra/repository"
	"github.com/jailtonjunior94/go-uow/internal/usecase"
	migration "github.com/jailtonjunior94/go-uow/pkg/database/migrate"
	database "github.com/jailtonjunior94/go-uow/pkg/database/postgres"
	"github.com/jailtonjunior94/go-uow/pkg/database/uow"
	"github.com/jailtonjunior94/go-uow/pkg/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	dbConn, err := database.NewPostgresDatabase(config)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	migrate, err := migration.NewMigrate(logger, dbConn, config.MigratePath, config.DBName)
	if err != nil {
		panic(err)
	}

	if err = migrate.ExecuteMigration(); err != nil {
		logger.Error(err)
	}

	courseRepository := repository.NewCourseRepository(dbConn)
	categoryRepository := repository.NewCategoryRepository(dbConn)
	addCourseUseCase := usecase.NewAddCourseUseCase(logger, courseRepository, categoryRepository)

	uow := uow.NewUow(context.Background(), dbConn)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbConn)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbConn)
		repo.Queries = db.New(tx)
		return repo
	})

	addCourseUowUseCase := usecase.NewAddCourseUowUseCase(logger, uow)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		AddCourse:    addCourseUseCase,
		AddCourseUow: addCourseUowUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.GraphQLPort)
	log.Fatal(http.ListenAndServe(":"+config.GraphQLPort, nil))
}
