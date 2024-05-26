package servers

import (
	"github.com/MarkTBSS/058_Router_Check/modules/middlewares/middlewaresHandlers"
	"github.com/MarkTBSS/058_Router_Check/modules/middlewares/middlewaresRepositories"
	"github.com/MarkTBSS/058_Router_Check/modules/middlewares/middlewaresUsecases"
	_pkgModulesMonitorMonitorHandlers "github.com/MarkTBSS/058_Router_Check/modules/monitor/monitorHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	handler := middlewaresHandlers.MiddlewaresHandler(usecase, s.cfg)
	return middlewaresHandlers.MiddlewaresHandler(handler, s.cfg)
}

func (m *moduleFactory) MonitorModule() {
	handler := _pkgModulesMonitorMonitorHandlers.MonitorHandler(m.s.cfg)
	m.r.Get("/", handler.HealthCheck)
}
