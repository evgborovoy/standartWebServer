package api

import (
	"net/http"

	"github.com/evgborovoy/StandardWebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base Api server instance description
type API struct {
	// UNEXPORTED FIELD
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

// API constructor: build base API instance

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		// Добавление поля для работы с хранилищем

	}
}

// Start http server/configure loggers, router, db connection
func (api *API) Start() error {
	// trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting API server at port", api.config.BindAddr)

	// Конфигурируем маршрутизатор
	api.configureRouterField()
	// Конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}
	// На этапе выполненого завершения запускаем сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
