package workspaces

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type workspaceRepository struct {
	repository.GenericRepository[Workspace]
}

func NewWorkspaceRepository(db *gorm.DB) workspaceRepository {
	return workspaceRepository{
		GenericRepository: repository.NewGenericRepository[Workspace](db),
	}
}

var _ WorkspaceRepositoryPort = (*workspaceRepository)(nil)

type workspaceMemberRepository struct {
	repository.GenericRepository[WorkspaceMember]
}

func NewWorkspaceMemberRepository(db *gorm.DB) workspaceMemberRepository {
	return workspaceMemberRepository{
		GenericRepository: repository.NewGenericRepository[WorkspaceMember](db),
	}
}

var _ WorkspaceMemberRepositoryPort = (*workspaceMemberRepository)(nil)
