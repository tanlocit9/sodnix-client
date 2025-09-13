package types

import (
	"sodnix/apps/server/src/common"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&Type{}, &TypeGroup{}}
}

type Type struct {
	common.InformationModel
	common.UUIDTypeModel
	TypeGroupId uuid.UUID `json:"typeGroupId,omitempty" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TypeGroup   TypeGroup `json:"typeGroup,omitempty" gorm:"not null;foreignKey:TypeGroupId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TypeGroup struct {
	common.InformationModel
	common.UUIDTypeModel
}
