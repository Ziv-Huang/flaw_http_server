package main

import (
	"flaw_http_server/domain/configs"
	log "flaw_http_server/domain/logger"
	"flaw_http_server/usecase/server"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var config configs.Configuration
	configs.ReadConfig(&config)
	log.SetLevel(config.Project.Log)
	log.Info(fmt.Sprintf("project: %s  version: %s", config.Project.Name, config.Project.Version))

	// Launch server
	wg.Add(1)
	log.Info("Launching server...")
	server.NewServer(&wg, config).LaunchServer()
	wg.Wait()

	log.Info("Server Finished!")
}
