package users

import (
	"sodnix/apps/server/src/common"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&User{}}
}

func PublicFields() []string {
	return []string{"id", "name", "email"}
}

type User struct {
	Email       string `gorm:"uniqueIndex"`
	Password    string `gorm:"type:text"`
	DisplayName string `gorm:"type:text"`
	Username    string `gorm:"unique;not null"`
	common.UUIDTypeModel
}

func (u *User) GetID() uuid.UUID    { return u.ID }
func (u *User) GetUserName() string { return u.Username }
func (u *User) GetEmail() string    { return u.Email }
