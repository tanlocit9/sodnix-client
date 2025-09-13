package transactions

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/accounts"
	"sodnix/apps/server/src/modules/categories"
	"time"

	"github.com/google/uuid"
)

func Models() []any {
	return []any{&Transaction{}}
}

type Transaction struct {
	common.WorkspaceScoped
	common.UUIDTypeModel
	Amount          float64             `json:"amount,omitempty" gorm:"type:decimal(10,2);not null;default:0"`
	TransactionDate time.Time           `json:"transactionDate,omitempty" gorm:"not null"`
	CategoryId      uuid.UUID           `json:"categoryId,omitempty" gorm:"not null"`
	Category        categories.Category `json:"category,omitempty" gorm:"not null"`
	SourceId        uuid.UUID           `json:"sourceId,omitempty" gorm:"not null;foreignKey:SourceId"`
	Source          accounts.Account    `json:"source,omitempty" gorm:"not null"`
	DestinationId   uuid.UUID           `json:"destinationId,omitempty" gorm:"not null;foreignKey:DestinationId"`
	Destination     accounts.Account    `json:"destination,omitempty" gorm:"not null"`
}
