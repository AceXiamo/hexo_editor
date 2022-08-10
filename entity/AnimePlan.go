package entity

import (
	"gorm.io/gorm"
)

type Sources struct {
	gorm.Model
	Id         string  `json:"id,omitempty" gorm:"primary_key"`
	SourceUrl  string  `json:"source_url,omitempty"`
	SourceType string  `json:"source_type,omitempty"`
	Remark     *string `json:"remark,omitempty"`
	CreateTime string  `json:"create_time"`
}

type AnimePlan struct {
	gorm.Model   `json:"gorm_._model"`
	Id           string `json:"id,omitempty" gorm:"primary_key"`
	Title        string `json:"title,omitempty"`
	Remark       string `json:"remark,omitempty"`
	RecordStatus string `json:"record_status,omitempty"`
	RecordType   string `json:"record_type,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
}

type AnimePlanDate struct {
	gorm.Model `json:"gorm_._model"`
	Id         string `json:"id,omitempty" gorm:"primary_key"`
	Pid        string `json:"pid,omitempty"`
	ViewDate   string `json:"view_date,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type AnimePlanSource struct {
	gorm.Model `json:"gorm_._model"`
	Id         string `json:"id,omitempty" gorm:"primary_key"`
	Pid        string `json:"pid,omitempty"`
	SourceUrl  string `json:"source_url,omitempty"`
	SourceType string `json:"source_type,omitempty"`
	Remark     string `json:"remark,omitempty"`
	CreateTime string `json:"create_time"`
}
