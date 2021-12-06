//go:build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/multipleton/sa-3/disks"
	"github.com/multipleton/sa-3/utils"
)

func composeHTTPControllers(disksController *disks.DisksController) []utils.HTTPController {
	return []utils.HTTPController{disksController}
}

func InitializeApplication(port HttpPortNumber) (*ApiServer, error) {
	panic(wire.Build(
		wire.Struct(new(sql.DB)), // TODO: stub, replace with correct initializer
		mux.NewRouter,
		disks.Module,
		composeHTTPControllers,
		wire.Struct(new(ApiServer), "port", "router", "controllers"),
	))
}
