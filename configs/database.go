package configs

import (
	"api/internal/establishment"
	"api/internal/store"
	"api/pkg/database"

	"gorm.io/gorm"
)

var DB *gorm.DB 

func InitDatabase() error {
	var err error
	DB, err = database.NewPostgreSQLGORMDB() 
	if err != nil {
		return err
	}

	// migração automática para todos os modelos
	return database.InitSchemaGORM(DB,
		&establishment.Establishment{}, 
		&store.Store{},
	)
}