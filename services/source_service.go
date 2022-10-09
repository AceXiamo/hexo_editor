package services

import (
	uuid "github.com/satori/go.uuid"
	"hexo_editor/entity"
	"hexo_editor/utils"
	"time"
)

// SourceByPage
//
//	@Description:
//	@param page
//	@param size
//	@return interface{}
func SourceByPage(page, size int) interface{} {
	limit, offset := utils.ToLimit(page, size)
	var image []entity.Sources
	sql := utils.ImageDb.Model(&entity.Sources{})
	sql.Select("sources.id", "sources.source_name", "sources.source_url", "sources.source_type", "sources.remark", "sources.create_time")
	sql.Limit(limit).Offset(offset).Order("sources.create_time DESC")
	sql.Find(&image)
	return image
}

// SaveSource
//
//	@Description: save
//	@param source
//	@return string
func SaveSource(source *entity.Sources) string {
	source.Id = uuid.NewV4().String()
	source.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	utils.ImageDb.Create(&source)
	return "操作成功"
}

// DelSource
//
//	@Description:
//	@param id
//	@return string
func DelSource(id string) string {
	utils.ImageDb.Where("id = ?", id).Delete(&entity.Sources{})
	return "操作成功"
}

// UpdateSource
//
//	@Description:
//	@param source
//	@return string
func UpdateSource(source *entity.Sources) string {
	utils.ImageDb.Save(&source)
	return "操作成功"
}
