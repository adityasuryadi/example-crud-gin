package main

import (
	"app/configs"
	"app/routes"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	appPort := os.Getenv("APP_PORT")
	router.Run("0.0.0.0:" + appPort)
	// log.Fatal(router.Run(":"))
}

func setupRouter() *gin.Engine {
	db := configs.InitDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	// init route
	routes.InitCustomerRoute(db, router)
	routes.InitUserRoute(db, router)

	return router
}
