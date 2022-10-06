package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Calmantara/go-fga/config/postgres"
	"github.com/gin-gonic/gin"

	engine "github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/pkg/domain/message"
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
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		// secara default map jika di return dalam
		// response API, dia akan menjadi JSON
		respMap := map[string]any{
			"code":       0,
			"message":    "server up and running",
			"start_time": startTime,
		}

		// golang memiliki json package
		// json package bisa mentranslasikan
		// map menjadi suatu struct
		// nb: struct harus memiliki tag/annotation JSON
		var respStruct message.Response

		// marshal -> mengubah json/struct/map menjadi
		// array of byte atau bisa kita translatekan menjadi string
		// dengan format JSON
		resByte, err := json.Marshal(respMap)
		if err != nil {
			panic(err)
		}
		// unmarshal -> translasikan string/[]byte dengan format JSON
		// menjadi map/struct dengan tag/annotation json
		err = json.Unmarshal(resByte, &respStruct)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, respStruct)
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
