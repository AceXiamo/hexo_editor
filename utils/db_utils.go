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

	// db
	dbName := "image.db"
	ImageDb, ImageDbErr = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if ImageDbErr != nil {
		// 无文件，创建后重连
		GenFileToRoot(dbName)
		ImageDb, _ = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	}

	// table init
	ImageInit()
	AnimePlanInit()
}

// ImageInit
// @Description: Image
func ImageInit() {
	res := ImageDb.First(&entity.Image{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.Image{})
		if err != nil {
			return
		}
	}
	res = ImageDb.First(&entity.ImageDetail{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.ImageDetail{})
		if err != nil {
			return
		}
	}
	res = ImageDb.First(&entity.ImageAuthor{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.ImageAuthor{})
		if err != nil {
			return
		}
	}
}

func AnimePlanInit() {
	res := ImageDb.First(&entity.AnimePlan{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.AnimePlan{})
		if err != nil {
			return
		}
	}
	res = ImageDb.First(&entity.AnimePlanDate{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.AnimePlanDate{})
		if err != nil {
			return
		}
	}
	res = ImageDb.First(&entity.AnimePlanSource{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.AnimePlanSource{})
		if err != nil {
			return
		}
	}
	res = ImageDb.First(&entity.Sources{})
	if res.Error != nil {
		err := ImageDb.AutoMigrate(&entity.Sources{})
		if err != nil {
			return
		}
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
