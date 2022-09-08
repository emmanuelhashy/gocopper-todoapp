package app

import (
	"github.com/emmanuelhashy/todolist/pkg/todos"
	"github.com/google/wire"
)

var WireModule = wire.NewSet(
	todos.WireModule,
)
