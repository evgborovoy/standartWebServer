package api

// Base Api server instance description
type API struct {
}

// API constructor: build base API instance

func New() *API {
	return &API{}
}

// Start http server/configure loggers, router, db connection
func (api *API) Start() error {
	return nil
}
