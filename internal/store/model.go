// internal/store/model.go
package store

import (
	"api/internal/establishment"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	EstablishmentID uint   `gorm:"not null" json:"establishment_id"`
	StoreNumber     string         `gorm:"type:varchar(50);not null" json:"store_number"`
	Name            string         `gorm:"type:varchar(255);not null" json:"name"`
	CorporateName   string         `gorm:"type:varchar(255);not null" json:"corporate_name"`
	Address         string         `gorm:"type:varchar(255);not null" json:"address"`
	City            string         `gorm:"type:varchar(100);not null" json:"city"`
	State           string    `gorm:"type:varchar(2);not null" json:"state"`
	Code       		  string    `gorm:"type:varchar(10);not null" json:"code"`
	Number          string    `gorm:"type:varchar(20);not null" json:"number"`
	Establishment   establishment.Establishment `gorm:"foreignKey:EstablishmentID" json:"-"` 

}


type EstablishmentWithStores struct {
	establishment.Establishment 
	Stores []Store `json:"stores"`
}