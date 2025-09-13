package workspaces

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/users"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&Workspace{}, &WorkspaceMember{}}
}

type Workspace struct {
	common.UUIDTypeModel
	common.InformationModel
	Slug    string     `gorm:"size:100;uniqueIndex;not null"`
	OwnerID uuid.UUID  `gorm:"not null"`
	Owner   users.User `gorm:"foreignKey:OwnerID"` // Association to User

	Members []users.User `gorm:"many2many:workspace_members"` // Many-to-many
}

type WorkspaceMember struct {
	common.UUIDTypeModel
	common.InformationModel
	WorkspaceID uuid.UUID `gorm:"index;not null"`
	UserID      uuid.UUID `gorm:"index;not null"`
	Role        string    `gorm:"size:50;default:'member'"` // e.g., 'admin', 'viewer'

	Workspace Workspace
	User      users.User
}
