package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/// ---------- 🔹 Core Base Models ----------

// Base model with UUID ID and timestamps
type UUIDTypeModel struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:public.uuid_generate_v4()" json:"id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	AuditFields
}

type AuditFields struct {
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"createdAt,omitempty" example:"2025-06-08T12:34:56Z"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updatedAt,omitempty" example:"2025-06-08T14:56:12Z"`

	CreatedByID uuid.UUID `gorm:"column:created_by;type:uuid;not null" json:"createdById" example:"123e4567-e89b-12d3-a456-426614174000"`
	UpdatedByID uuid.UUID `gorm:"column:updated_by;type:uuid;not null" json:"updatedById" example:"123e4567-e89b-12d3-a456-426614174001"`

	CreatedBy *UserRef `gorm:"foreignKey:CreatedByID;references:ID" json:"createdByUser,omitempty"`
	UpdatedBy *UserRef `gorm:"foreignKey:UpdatedByID;references:ID" json:"updatedByUser,omitempty"`
}

type UserRef struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string
	Email string
}

// TableName overrides the default table name used by GORM
func (UserRef) TableName() string {
	return "users"
}

/// ---------- 🔹 Logical Base Models ----------

// Global scoped
// Common name and description fields
type InformationModel struct {
	Name        string `gorm:"type:varchar(255);not null" json:"name,omitempty" example:"Resource Name"`
	Description string `gorm:"type:text;" json:"description,omitempty" example:"Resource description"`
}

// Entity scoped to a user (e.g., personal tags, filters)
type UserScoped struct {
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	InformationModel
}

// Entity scoped to a workspace (shared across users)
type WorkspaceScoped struct {
	WorkspaceID uuid.UUID `gorm:"type:uuid;not null;index"`
	InformationModel
}

/// ---------- 🔹 API Response Base ----------

type BaseResponseModel struct {
	ID          string `json:"id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID      string `json:"userId,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	WorkspaceID string `json:"tenantId,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	InformationModel
	AuditFields
}

// BeforeCreate hook for GORM to set CreatedAt and UpdatedAt
func (a *UUIDTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	a.AuditFields.CreatedAt = time.Now()
	a.AuditFields.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook for GORM to set UpdatedAt
func (a *UUIDTypeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	a.AuditFields.UpdatedAt = time.Now()
	return
}

type Auditable interface {
	SetCreatedBy(uuid.UUID)
	SetUpdatedBy(uuid.UUID)
}

type AuditableUpdateOnly interface {
	SetUpdatedBy(uuid.UUID)
}

// SetCreatedBy sets the CreatedByID field
func (a *UUIDTypeModel) SetCreatedBy(userID uuid.UUID) {
	a.AuditFields.CreatedByID = userID
}

// SetUpdatedBy sets the UpdatedByID field
func (a *UUIDTypeModel) SetUpdatedBy(userID uuid.UUID) {
	a.AuditFields.UpdatedByID = userID
}
