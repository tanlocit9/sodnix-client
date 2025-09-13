package categories

import (
	"sodnix/apps/server/src/common"

	"github.com/google/uuid"
)

type CategoryRequestDTO struct {
	TypeId   uuid.UUID  `json:"typeId,omitempty"`
	ParentId *uuid.UUID `json:"parentId,omitempty"` // nullable
	common.InformationModel
}

type CategoryResponseDTO struct {
	common.BaseResponseModel
	TypeId   uuid.UUID   `json:"typeId,omitempty"`
	ParentId *uuid.UUID  `json:"parentId,omitempty"` // nullable
	Children []*Category `json:"children,omitempty"`
}
