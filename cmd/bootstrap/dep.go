package bootstrap

import (
	"sync"

	"github.com/edfoh/go-api-starter/internal/adding"
	"github.com/edfoh/go-api-starter/internal/storage/memory"
	"github.com/edfoh/go-api-starter/pkg/log"
	"github.com/julienschmidt/httprouter"
)

type Dependencies struct {
	routerInit sync.Once
	router     *httprouter.Router

	serviceAdderInit sync.Once
	serviceAdder     adding.Service

	storageMemory     *memory.Storage
	storageMemoryInit sync.Once

	loggerInit sync.Once
	logger     log.Logger
}

func (d *Dependencies) Router() *httprouter.Router {
	d.routerInit.Do(func() {
		d.router = httprouter.New()
	})

	return d.router
}

func (d *Dependencies) ServiceAdder(r adding.Repository, logger log.Logger) adding.Service {
	d.serviceAdderInit.Do(func() {
		d.serviceAdder = adding.NewService(r, logger)
	})

	return d.serviceAdder
}

func (d *Dependencies) StorageMemory(logger log.Logger) *memory.Storage {
	d.storageMemoryInit.Do(func() {
		d.storageMemory = memory.NewStorage(logger)
	})

	return d.storageMemory
}

func (d *Dependencies) Logger(level log.Level) log.Logger {
	d.loggerInit.Do(func() {
		d.logger = log.Must(log.New(level))
	})

	return d.logger
}
