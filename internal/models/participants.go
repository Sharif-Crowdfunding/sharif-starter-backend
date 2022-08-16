package models

import "github.com/jinzhu/gorm"

type Participant struct {
	gorm.Model
	WalletAddress   string
	PurchasedTokens int64
	ProjectTokenId  int64
	Transaction     string
	IsConfirmed     bool
}
