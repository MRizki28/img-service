package ImageController

import (
	"net/http"

	"github.com/MRizki28/img-service/Config"
	"github.com/MRizki28/img-service/Database"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"github.com/google/uuid"
)

func GetAllData(c *gin.Context) {
	var TbImage []Database.TbImage
	err := Config.DB.Find(&TbImage).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Success get All Data",
			"data":    TbImage,
		})
	}
}


func UploadImage(c *gin.Context) {
	file, err := c.FormFile("name_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error disini kh",
			"error":   err.Error(),
		})
		return
	}

	randomStr := uuid.New().String()

	fileExt := filepath.Ext(file.Filename)

	newFilename :=  randomStr + fileExt

	filePath := filepath.Join("uploads", newFilename)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error disini kh",
			"error":   err.Error(),
		})
		return
	}

	image := Database.TbImage{Name_File: newFilename}
	result := Config.DB.Create(&image)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error save to db",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success upload file",
		"data":    newFilename,
	})
}
