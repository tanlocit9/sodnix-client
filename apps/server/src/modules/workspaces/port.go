package workspaces

import (
	"sodnix/apps/server/src/common/port"
)

type WorkspaceServicePort interface {
	port.CrudServicePort[Workspace, WorkspaceRequestDTO, WorkspaceResponseDTO]
}

type WorkspaceRepositoryPort interface {
	port.CrudRepositoryPort[Workspace]
}

type WorkspaceMemberServicePort interface {
	port.CrudServicePort[WorkspaceMember, WorkspaceMemberRequestDTO, WorkspaceMemberResponseDTO]
}

type WorkspaceMemberRepositoryPort interface {
	port.CrudRepositoryPort[WorkspaceMember]
}
