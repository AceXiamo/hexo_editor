package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/entity"
	"hexo_editor/global"
	service "hexo_editor/services"
	"hexo_editor/utils"
	"strconv"
)

type ImgJson struct {
	Image       []entity.Image       `json:"image,omitempty"`
	ImageDetail []entity.ImageDetail `json:"image_detail,omitempty"`
	ImageAuthor entity.ImageAuthor   `json:"image_author"`
}

type ImgJson1 struct {
	Image       interface{} `json:"image,omitempty"`
	ImageDetail interface{} `json:"image_detail,omitempty"`
	ImageAuthor ImageAuthor `json:"image_author"`
}

type Image struct {
	Id         string `json:"id,omitempty"`
	TypeId     string `json:"type_id,omitempty"`
	Url        string `json:"url,omitempty"`
	CreateTime string `json:"create_time"`
}

type ImageDetail struct {
	Id           string `json:"id,omitempty"`
	ImageId      string `json:"image_id,omitempty"`
	AuthorId     string `json:"author_id,omitempty"`
	OriginalPath string `json:"original_path,omitempty"`
	PixivCode    string `json:"pixiv_code,omitempty"`
	CreateTime   string `json:"create_time"`
}

type ImageAuthor struct {
	Id           string `json:"id,omitempty"`
	Nick         string `json:"nick,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	PixivCode    string `json:"pixiv_code,omitempty"`
	PixivHomeUrl string `json:"pixiv_home_url,omitempty"`
	CreateTime   string `json:"create_time"`
}

// PixivInfo
// @Description: 根据插画路径获取详情
// @param ctx
func PixivInfo(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	global.Result(ctx, 200, "success", service.PixivInfo(ctx.URLParam("url")))
}

// GetImageByte
// @Description: 根据P站插画路径取图
// @param ctx
func GetImageByte(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	// 写入图片 字节数组
	_, err := ctx.Write(service.GetImageByte(ctx.URLParam("url")))
	if err != nil {
		return
	}
}

// ImageByPage
// @Description: 分页查询已保存插画
// @param ctx
func ImageByPage(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	page, _ := strconv.Atoi(ctx.URLParam("page"))
	size, _ := strconv.Atoi(ctx.URLParam("size"))
	nick := ctx.URLParam("nick")
	global.Result(ctx, 200, "success", service.ImageByPage(page, size, nick))
}

// SaveImage
// @Description: 保存插画
// @param ctx
func SaveImage(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json ImgJson
	err := ctx.ReadJSON(&json)
	if err != nil {
		return
	}
	global.Result(ctx, 200, "success", service.SaveImage(json.Image, json.ImageDetail, json.ImageAuthor))
}

// ImageInfo
// @Description: 根据id查询已保存插画
// @param ctx
func ImageInfo(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	global.Result(ctx, 200, "success", service.ImageInfo(ctx.URLParam("id")))
}

// DelImage
// @Description: 根据id删除插画
// @param ctx
func DelImage(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	global.Result(ctx, 200, "success", service.DelImage(ctx.URLParam("id")))
}

// AuthList
// @Description: 查询所有已保存作者信息，一次性返回不进行分页
// @param ctx
func AuthList(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	global.Result(ctx, 200, "success", service.AuthList())
}

// PixivNum
//  @Description: 获取收集插画的数量
//  @param ctx
func PixivNum(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	global.Result(ctx, 200, "success", service.PixivNum())
}
