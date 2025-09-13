package workspaces

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/users"

	"github.com/google/uuid"
)

type WorkspaceRequestDTO struct {
	common.InformationModel
	Slug    string    `json:"slug,omitempty"`
	OwnerID uuid.UUID `json:"ownerId,omitempty"`
}

type WorkspaceResponseDTO struct {
	common.BaseResponseModel
	Slug    string       `json:"slug,omitempty"`
	OwnerID uuid.UUID    `json:"ownerId,omitempty"`
	Owner   users.User   `json:"owner,omitempty"`
	Members []users.User `json:"members,omitempty"`
}

type WorkspaceMemberRequestDTO struct {
	common.InformationModel
	WorkspaceID uuid.UUID `json:"workspaceId,omitempty"`
	UserID      uuid.UUID `json:"userId,omitempty"`
	Role        string    `json:"role,omitempty"`
}

type WorkspaceMemberResponseDTO struct {
	common.BaseResponseModel
	WorkspaceID uuid.UUID  `json:"workspaceId,omitempty"`
	UserID      uuid.UUID  `json:"userId,omitempty"`
	Role        string     `json:"role,omitempty"`
	Workspace   Workspace  `json:"workspace,omitempty"`
	User        users.User `json:"user,omitempty"`
}
