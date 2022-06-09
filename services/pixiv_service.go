package services

import (
	"hexo_editor/entity"
	"hexo_editor/pixiv"
	"hexo_editor/utils"
	"sort"
)

// PixivInfo
// @Description: 根据插画路径获取详情
// @param url
func PixivInfo(url string) map[string]interface{} {
	return pixiv.GetPixivInfo(url)
}

// GetImageByte
// @Description: 根据P站插画路径取图
// @param url
// @return []byte
func GetImageByte(url string) []byte {
	return pixiv.GetImageByte(url)

}

// ImageByPage
// @Description: 分页查询已保存插画
// @param page
// @param size
func ImageByPage(page, size int,nick string) interface{} {
	limit, offset := utils.ToLimit(page, size)
	var image []entity.Image
	sql := utils.ImageDb.Model(&entity.Image{}).Select("images.id", "images.created_at", "images.updated_at", "images.deleted_at", "images.type_id", "images.url", "images.create_time")
	if len(nick) > 0 {
		sql.Joins("LEFT JOIN image_details on image_details.image_id = images.id").Joins("LEFT JOIN image_authors ON image_details.author_id = image_authors.id")
		sql.Where("image_authors.nick = ?", nick)
	}
	sql.Limit(limit).Offset(offset).Order("images.created_at DESC")
	sql.Find(&image)
	for i := 0; i < len(image); i++ {
		utils.ImageDb.Model(&entity.ImageDetail{}).Where("image_id = ?", image[i].Id).First(&image[i].Detail)
		utils.ImageDb.Model(&entity.ImageAuthor{}).Where("id = ?", image[i].Detail.AuthorId).First(&image[i].Auth)
	}
	return image
}

// SaveImage
// @Description: 保存插画
// @param images
// @param detail
// @param auth
func SaveImage(images []entity.Image, detail []entity.ImageDetail, auth entity.ImageAuthor) string {
	for i := 0; i < len(images); i++ {
		utils.ImageDb.Create(&images[i])
		utils.ImageDb.Create(&detail[i])
	}
	var total int64 = 0
	utils.ImageDb.Model(&entity.ImageAuthor{}).Where("id = ?", detail[0].AuthorId).Count(&total)
	if total < 1 {
		utils.ImageDb.Create(&auth)
	}
	return "保存成功!"
}

// ImageInfo
// @Description: 根据id查询已保存插画
// @param id
// @return interface{}
func ImageInfo(id string) interface{} {
	var image entity.Image
	utils.ImageDb.Where("id = ?", id).First(&image)
	utils.ImageDb.Where("image_id = ?", id).First(&image.Detail)
	utils.ImageDb.Where("id = ?", image.Detail.AuthorId).First(&image.Auth)
	return image
}

// DelImage
// @Description: 根据id删除插画
// @param id
// @return string
func DelImage(id string) string {
	image := new(entity.Image)
	utils.ImageDb.Where("id = ?", id).Delete(&image)
	detail := new(entity.ImageDetail)
	utils.ImageDb.Where("image_id = ?", id).Delete(&detail)
	return "删除成功!"
}

// AuthList
// @Description: 查询所有已保存作者信息，一次性返回不进行分页
// @param ctx
func AuthList() interface{} {
	var author []entity.ImageAuthor
	utils.ImageDb.Find(&author)
	for i := 0; i < len(author); i++ {
		utils.ImageDb.Model(&entity.ImageDetail{}).Where("author_id = ?", author[i].Id).Count(&author[i].Num)
	}
	sort.Slice(author, func(i, j int) bool {
		return author[i].Num > author[j].Num
	})
	return author
}
