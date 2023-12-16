package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Initialize() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Erro ao Inicializar SQLite: %v", err)
	}

	return nil
}

func GetSQLiteDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
