package main

import (
	"github.com/MRizki28/img-service/App/Controllers/ImageController"
	"github.com/MRizki28/img-service/Config"
	"github.com/MRizki28/img-service/Config/Cors"
	"github.com/MRizki28/img-service/Database"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    Config.ConnectionDatabase()
    Database.Migrate()

    api := r.Group("/api")
    api.Use(Cors.CorsMiddleware())
    api.GET("/images", ImageController.GetAllData)
    api.POST("/upload", ImageController.UploadImage)

    r.Run()
}
