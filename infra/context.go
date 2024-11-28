package infra

import (
	"lumoshive-be-chap38-39/config"
	"lumoshive-be-chap38-39/controller"
	"lumoshive-be-chap38-39/database"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.MainController
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	// instance controller
	Ctl := controller.NewMainController(db)

	return &ServiceContext{Cfg: config, Ctl: Ctl}, nil
}
