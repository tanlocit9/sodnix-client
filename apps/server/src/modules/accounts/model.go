package accounts

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/types"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&Account{}}
}

type Account struct {
	common.InformationModel
	common.UUIDTypeModel
	common.UserScoped
	Balance        float64         `json:"balance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	InitialBalance float64         `json:"initialBalance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	TypeGroupId    uuid.UUID       `json:"typeGroupId,omitempty" gorm:"not null"`
	TypeGroup      types.TypeGroup `json:"typeGroup,omitempty" gorm:"not null;foreignKey:TypeGroupId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
