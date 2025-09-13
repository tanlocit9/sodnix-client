package categories

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/types"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&Category{}}
}

type Category struct {
	common.InformationModel
	common.UUIDTypeModel
	TypeId   uuid.UUID   `gorm:"not null" json:"typeId,omitempty"`
	Type     types.Type  `gorm:"not null;foreignKey:TypeId" json:"type,omitempty"`
	ParentId *uuid.UUID  `gorm:"column:parent_id" json:"parentId,omitempty"` // nullable
	Parent   *Category   `gorm:"foreignKey:ParentId"`                        // self-referencing
	Children []*Category `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"children,omitempty"`
}
