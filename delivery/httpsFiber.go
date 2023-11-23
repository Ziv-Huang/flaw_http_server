package delivery

import (
	"flaw_http_server/domain/configs"
	"flaw_http_server/usecase/api"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type httpFiber struct {
	Wg     *sync.WaitGroup
	Config configs.Configuration
}

func NewHttpFiber(wg *sync.WaitGroup, config configs.Configuration) *httpFiber {
	return &httpFiber{Wg: wg, Config: config}
}

func (s *httpFiber) Start() {

	go func() {
		defer s.Wg.Done()

		// Fiber instance
		app := fiber.New()

		// Fiber operations
		app.Get("/", api.Hello)
		app.Get("/heartbeat", api.Heartbeat)
		app.Get("/deviceinfo", api.GetDeviceInfo)
		app.Get("/result/:id", api.GetResultByID)
		app.Post("/job", api.CreateAJob)
		app.Delete("/job/:id", api.DeleteAJob)

		// Start server
		app.Listen(":" + s.Config.Http.Port)
	}()

}
