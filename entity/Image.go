package entity

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Id         string `json:"id,omitempty"`
	TypeId     string `json:"type_id,omitempty"`
	Url        string `json:"url,omitempty"`
	CreateTime string `json:"create_time"`

	Detail ImageDetail `gorm:"-" json:"detail"`
	Auth   ImageAuthor `gorm:"-" json:"auth"`
}

type ImageDetail struct {
	gorm.Model
	Id           string `json:"id,omitempty"`
	ImageId      string `json:"image_id,omitempty"`
	AuthorId     string `json:"author_id,omitempty"`
	OriginalPath string `json:"original_path,omitempty"`
	PixivCode    string `json:"pixiv_code,omitempty"`
	CreateTime   string `json:"create_time"`
}

type ImageAuthor struct {
	gorm.Model
	Id           string `json:"id,omitempty"`
	Nick         string `json:"nick,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	PixivCode    string `json:"pixiv_code,omitempty"`
	PixivHomeUrl string `json:"pixiv_home_url,omitempty"`
	CreateTime   string `json:"create_time"`

	Num int64 `gorm:"-" json:"num"`
}
