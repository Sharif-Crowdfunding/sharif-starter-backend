package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	UserEmail string
	Name      string
	Token     *ProjectToken
	*ProjectBasicInfo
	Status ProjectStatus `gorm:"default:0"`
}

type ProjectBasicInfo struct {
	Website    string
	TelegramId string
	TokenInfo  string
	Details    string
	GithubId   string
}

type ProjectStatus int

const (
	Start ProjectStatus = iota
	ReadyForSale
	InProgress
	Finished
)

//func (project *Project) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("ID", uuid.NewV4().String())
//	return nil
//}
