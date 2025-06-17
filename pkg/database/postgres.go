package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLGORMDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Limite para consultas lentas
			LogLevel:                  logger.Info,   // Nível de log (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,          // Ignorar erro de "registro não encontrado"
			ParameterizedQueries:      true,          // Mostrar consultas parametrizadas
			Colorful:                  true,          // Saída colorida
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados com GORM: %w", err)
	}

	log.Println("Conectado ao banco de dados PostgreSQL com GORM sucesso!")
	return db, nil
}

func InitSchemaGORM(db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("erro ao realizar AutoMigrate para o modelo %T: %w", model, err)
		}
		log.Printf("Tabela para o modelo '%T' verificada/migrada.\n", model)
	}
	return nil
}