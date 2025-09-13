package transactions

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/modules/accounts"
	"sodnix/apps/server/src/modules/categories"
	"time"

	"github.com/google/uuid"
)

type TransactionRequestDTO struct {
	Amount          float64   `json:"amount,omitempty" example:"1"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	CategoryId      uuid.UUID `json:"categoryId,omitempty" example:"1"`
	SourceId        uuid.UUID `json:"sourceId,omitempty" gorm:"not null;foreignKey:SourceId"`
	DestinationId   uuid.UUID `json:"destinationId,omitempty" gorm:"not null;foreignKey:DestinationId"`
	common.InformationModel
}

type TransactionResponseDTO struct {
	Amount          float64             `json:"amount,omitempty" example:"100000"`
	TransactionDate time.Time           `json:"transaction_date,omitempty" example:"2025-06-08T12:34:56Z"`
	CategoryId      uuid.UUID           `json:"categoryId,omitempty" example:"6c0ce3eb-fd65-4a2c-aadc-de10ac852f81"`
	Category        categories.Category `json:"category,omitempty" gorm:"not null"`
	SourceId        uuid.UUID           `json:"sourceId,omitempty" gorm:"not null;foreignKey:SourceId" example:"2e9211a7-b41e-46bf-9788-75f69b4edfa3"`
	Source          accounts.Account    `json:"source,omitempty" gorm:"not null"`
	DestinationId   uuid.UUID           `json:"destinationId,omitempty" gorm:"not null;foreignKey:DestinationId" example:"1d65ac65-61ff-4d2b-863d-e98c2a86c9f4"`
	Destination     accounts.Account    `json:"destination,omitempty" gorm:"not null"`
	common.BaseResponseModel
}
