package account

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account account struct
type Account struct {
	ID        uuid.UUID  `json:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	Name      string     `json:"name" name:"ユーザー名"`
	Email     string     `json:"email" name:"メールアドレス"`
	Password  string     `json:"password" name:"パスワード"`
	CreatedBy uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time `json:"deletedAt" name:"削除日"`
}
