//go:build wireinject
// +build wireinject

package main

import (
	"github.com/emmanuelhashy/todolist/pkg/app"
	"github.com/emmanuelhashy/todolist/web"
	"github.com/emmanuelhashy/todolist/web/build"
	"github.com/gocopper/copper/csql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gocopper/copper"
	"github.com/gocopper/copper/chttp"
	"github.com/google/wire"
)

func InitServer(*copper.App) (*chttp.Server, error) {
	panic(
		wire.Build(WireModule),
	)
}

var WireModule = wire.NewSet(
	csql.WireModule,

	wire.InterfaceValue(new(chttp.HTMLDir), web.HTMLDir),
	wire.InterfaceValue(new(chttp.StaticDir), build.StaticDir),
	web.HTMLRenderFuncs,

	copper.WireModule,
	chttp.WireModule,
	wire.Struct(new(app.NewHTTPHandlerParams), "*"),
	app.NewHTTPHandler,
	app.WireModule,
	app.NewRouter,
	wire.Struct(new(app.NewRouterParams), "*"),
)
