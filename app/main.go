package main

import (
	"app/configs"
	"app/routes"
	"os"

	"app/controllers"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

}

func main() {
	router := SetupRouter()
	appPort := os.Getenv("APP_PORT")
	router.Run("0.0.0.0:" + appPort)
	// log.Fatal(router.Run(":"))
}

func SetupRouter() *gin.Engine {
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
	router.GET("api/v1/test", controllers.Test)

	// init route
	routes.InitCustomerRoute(db, router)
	routes.InitUserRoute(db, router)

	return router
}
