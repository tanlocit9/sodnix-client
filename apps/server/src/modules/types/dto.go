package types

import (
	"sodnix/apps/server/src/common"

	"github.com/google/uuid"
)

type TypeRequestDTO struct {
	TypeGroupId uuid.UUID `json:"typeGroupId,omitempty" example:"076c9d54-203d-41d1-8ea5-b42cd7b727f2"`
	common.InformationModel
}

type TypeResponseDTO struct {
	TypeGroupId uuid.UUID `json:"typeGroupId,omitempty" example:"076c9d54-203d-41d1-8ea5-b42cd7b727f2"`
	common.BaseResponseModel
}

type TypeGroupRequestDTO struct {
	common.InformationModel
}

type TypeGroupResponseDTO struct {
	common.BaseResponseModel
}
