package Database

import(
	"github.com/MRizki28/img-service/Config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type tb_image struct {
	Id        int64  `gorm:"primaryKey" json:"id`
	Name_File string `gorm:"type:varchar(300)" json:"name_file"`
}

func (tb_image) TableName() string {
	return "tb_image"
}

func Migrate() {
	Config.DB.AutoMigrate(&tb_image{})
}
