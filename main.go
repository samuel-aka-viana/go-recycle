package main

import (
	"go-project/config"
	"go-project/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("init config error: %v", err)
		return
	}

	// initialize router
	router.InitializeRouter()
}
