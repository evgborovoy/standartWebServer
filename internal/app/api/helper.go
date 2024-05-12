package api

import (
	"net/http"

	"github.com/evgborovoy/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
)

// пытаемся отконфигурировать наш API instance (поле logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать наш маршрутизатор (поле router)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm work!"))
	})
}

// Пытаемся отконфигурировать базу данных (поле storage)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
