package workspaces

import (
	"sodnix/apps/server/src/common/mapper"
)

type workspaceMapper struct {
	mapper.GenericMapper[Workspace, WorkspaceRequestDTO, WorkspaceResponseDTO]
}

func NewWorkspaceMapper() *workspaceMapper {
	return &workspaceMapper{
		mapper.NewGenericMapper[Workspace, WorkspaceRequestDTO, WorkspaceResponseDTO](),
	}
}

type workspaceMemberMapper struct {
	mapper.GenericMapper[WorkspaceMember, WorkspaceMemberRequestDTO, WorkspaceMemberResponseDTO]
}

func NewWorkspaceMemberMapper() *workspaceMemberMapper {
	return &workspaceMemberMapper{
		mapper.NewGenericMapper[WorkspaceMember, WorkspaceMemberRequestDTO, WorkspaceMemberResponseDTO](),
	}
}
