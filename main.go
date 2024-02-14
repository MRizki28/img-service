package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MRizki28/img-service/Config"
	"github.com/MRizki28/img-service/Database"
	"github.com/MRizki28/img-service/App/Controllers/ImageController"
)

func main() {
	r := gin.Default()
	Config.ConnectionDatabase()
	Database.Migrate()

	r.GET("/api/" , ImageController.GetAllData)
	r.POST("/api/upload" , ImageController.UploadImage)

	r.Run()
}
