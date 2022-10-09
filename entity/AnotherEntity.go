package entity

import (
	"gorm.io/gorm"
)

// Sources 保存相册中存放的资源
type Sources struct {
	gorm.Model
	Id        string `json:"id,omitempty" gorm:"primary_key"`
	SourceUrl string `json:"source_url,omitempty"`
	/* 类型（image: 图片， video: 视频，voice: 音频）*/
	SourceType string  `json:"source_type,omitempty"`
	SourceName string  `json:"source_name,omitempty"`
	Remark     *string `json:"remark,omitempty"`
	CreateTime string  `json:"create_time"`
}

// AnimePlan 追番数据
type AnimePlan struct {
	gorm.Model `json:"gorm_._model"`
	Id         string `json:"id,omitempty" gorm:"primary_key"`
	Title      string `json:"title,omitempty"`
	Remark     string `json:"remark,omitempty"`
	/* 记录状态（execute: 进行中，complete: 已完成） */
	RecordStatus string `json:"record_status,omitempty"`
	/* 记录类型（anime: 追番，another: ...） */
	RecordType string `json:"record_type,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

// AnimePlanDate 记录追番的周目
type AnimePlanDate struct {
	gorm.Model `json:"gorm_._model"`
	Id         string `json:"id,omitempty" gorm:"primary_key"`
	Pid        string `json:"pid,omitempty"`
	ViewDate   string `json:"view_date,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

// AnimePlanSource 关联的资源，比如说动漫的封面图之类的
type AnimePlanSource struct {
	gorm.Model `json:"gorm_._model"`
	Id         string `json:"id,omitempty" gorm:"primary_key"`
	Pid        string `json:"pid,omitempty"`
	SourceUrl  string `json:"source_url,omitempty"`
	/* 类型（image: 图片， video: 视频，voice: 音频） */
	SourceType string `json:"source_type,omitempty"`
	Remark     string `json:"remark,omitempty"`
	CreateTime string `json:"create_time"`
}
