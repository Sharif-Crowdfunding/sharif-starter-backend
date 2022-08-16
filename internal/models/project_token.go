package models

import "github.com/jinzhu/gorm"

type ProjectToken struct {
	gorm.Model
	ProjectId int64
	*ProjectTokenDistributionInfo
	*ProjectSale
}
type ProjectTokenDistributionInfo struct {
	TokenNumber int64
	TokenName   string
}

type ProjectSale struct {
	PricePerTokenByGwei int64
	MaximumTokenSale    int64
}
