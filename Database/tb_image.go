package Database


import (
    "github.com/MRizki28/img-service/Config"
)


type TbImage struct {
    Id        int64  `gorm:"primaryKey" json:"id"`
    Name_File string `gorm:"type:varchar(300)" json:"name_file"`
}


func (TbImage) TableName() string {
	return "tb_image"
}

func Migrate() {
	Config.DB.AutoMigrate(&TbImage{})
}
