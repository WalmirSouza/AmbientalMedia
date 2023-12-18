package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// _, err := os.Stat((dbPath))

	// if os.IsNotExist(err) {
	// 	logger.Info("Criando banco de dados...")

	// 	err = os.MkdirAll("./db", os.ModePerm)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	file, err := os.Create((dbPath))

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file.Close()
	// }

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao abrir o Banco de dados: %v", err)
		return nil, err
	}

	// err = db.AutoMigrate((&schemas.WindForecast{}))

	// if err != nil {
	// 	logger.Errorf("Erro no Automigrate Sqlite: %v", err)
	// 	return nil, err
	// }

	return db, nil
}
