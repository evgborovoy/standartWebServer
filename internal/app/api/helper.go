package api

import (
	"net/http"

	"github.com/evgborovoy/StandardWebServer/internal/app/middleware"
	"github.com/evgborovoy/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	// before JWT
	// a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	// after JWT
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods("GET")

	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/users/register", a.PostUserRegister).Methods("POST")
	a.router.HandleFunc(prefix+"/users/auth", a.PostUserAuth).Methods("POST")

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
