package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	_pkgConfig "github.com/MarkTBSS/058_Router_Check/config"
	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start()
}

type server struct {
	cfg _pkgConfig.IConfig
	app *fiber.App
}

func NewServer(cfg _pkgConfig.IConfig) IServer {
	return &server{
		cfg: cfg,
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	// Modules
	v1 := s.app.Group("v1")
	modules := InitModule(v1, s)
	modules.MonitorModule()

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
