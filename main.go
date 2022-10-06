package main

import (
	"github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/config/postgres"

	userrepo "github.com/Calmantara/go-fga/pkg/repository/user"
	userhandler "github.com/Calmantara/go-fga/pkg/server/http/handler/user"
	userrouter "github.com/Calmantara/go-fga/pkg/server/http/router/v1"
	userusecase "github.com/Calmantara/go-fga/pkg/usecase/user"
)

func main() {
	// generate postgres config and connect to postgres
	// this postgres client, will be used in repository layer
	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "postgresAdmin",
		DatabaseName: "postgres",
	})

	// gin engine
	ginEngine := gin.NewGinHttp(gin.Config{
		Port: ":8080",
	})

	// generate user repository
	userRepo := userrepo.NewUserRepo(postgresCln)
	// initiate use case
	userUsecase := userusecase.NewUserUsecase(userRepo)
	// initiate handler
	useHandler := userhandler.NewUserHandler(userUsecase)
	// initiate router
	userrouter.NewUserRouter(ginEngine, useHandler).Routers()

	// running the service
	ginEngine.Serve()
}
