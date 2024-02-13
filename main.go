package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MRizki28/img-service/Config"
	"github.com/MRizki28/img-service/Database"
)

func main() {
	r := gin.Default()
	Config.ConnectionDatabase()
	Database.Migrate()

	r.Run()
}
