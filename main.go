package main

import (
	"github.com/WalmirSouza/Challenge/config"
	"github.com/WalmirSouza/Challenge/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Initialize()

	if err != nil {
		logger.Errorf("Erro na inicialização do config: %v", err)
		return
	}
	router.Initialize()
}
