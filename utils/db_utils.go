package utils

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hexo_editor/entity"
)

var ImageDb *gorm.DB
var ImageDbErr error

// DbInit
// @Description: 对数据库进行初始化
func DbInit() {
	log.Info("DB Init...")
	ImageInit()
}

// ImageInit
// @Description: Image
func ImageInit() {
	dbName := "image.db"
	ImageDb, ImageDbErr = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if ImageDbErr != nil {
		// 无文件，创建后重连
		GenFileToRoot(dbName)
		ImageDb, _ = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	}
	// table init
	res := ImageDb.First(&entity.Image{})
	if res.Error != nil {
		ImageDb.AutoMigrate(&entity.Image{})
	}
	res = ImageDb.First(&entity.ImageDetail{})
	if res.Error != nil {
		ImageDb.AutoMigrate(&entity.ImageDetail{})
	}
	res = ImageDb.First(&entity.ImageAuthor{})
	if res.Error != nil {
		ImageDb.AutoMigrate(&entity.ImageAuthor{})
	}
}

func ToLimit(page, size int) (limit, offset int) {
	return size, (page - 1) * size
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
