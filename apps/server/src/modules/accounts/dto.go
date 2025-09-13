package accounts

import (
	"sodnix/apps/server/src/common"

	"github.com/google/uuid"
)

type AccountRequestDTO struct {
	Balance        float64   `json:"balance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	InitialBalance float64   `json:"initialBalance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	TypeGroupId    uuid.UUID `json:"typeGroupId,omitempty" gorm:"not null"`
	common.InformationModel
}

type AccountResponseDTO struct {
	Balance        float64   `json:"balance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	InitialBalance float64   `json:"initialBalance,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	TypeGroupId    uuid.UUID `json:"typeGroupId,omitempty" gorm:"not null"`
	common.BaseResponseModel
}
