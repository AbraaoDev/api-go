package establishment

import (
	"api/internal/store"

	"gorm.io/gorm"
)

type Establishment struct {
	gorm.Model
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	CorporateNumber string    `gorm:"type:varchar(20);not null" json:"corporate_number"`
	CorporateName   string    `gorm:"type:varchar(255);not null" json:"corporate_name"`
	Address         string    `gorm:"type:varchar(255);not null" json:"address"`
	City            string    `gorm:"type:varchar(100);not null" json:"city"`
	State           string    `gorm:"type:varchar(2);not null" json:"state"`
	Code       		  string    `gorm:"type:varchar(10);not null" json:"code"`
	Number          string    `gorm:"type:varchar(20);not null" json:"number"`
	Stores          []store.Store   `gorm:"foreignKey:EstablishmentID;constraint:OnDelete:RESTRICT;" json:"stores,omitempty"` 

}

