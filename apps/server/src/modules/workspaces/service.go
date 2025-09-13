package workspaces

import (
	"sodnix/apps/server/src/common/service"
)

type workspaceService struct {
	service.GenericService[Workspace, WorkspaceRequestDTO, WorkspaceResponseDTO]
}

func NewWorkspaceService(repo workspaceRepository, mapper *workspaceMapper) workspaceService {
	return workspaceService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ WorkspaceServicePort = (*workspaceService)(nil)

type workspaceMemberService struct {
	service.GenericService[WorkspaceMember, WorkspaceMemberRequestDTO, WorkspaceMemberResponseDTO]
}

func NewWorkspaceMemberService(repo workspaceMemberRepository, mapper *workspaceMemberMapper) workspaceMemberService {
	return workspaceMemberService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ WorkspaceMemberServicePort = (*workspaceMemberService)(nil)
