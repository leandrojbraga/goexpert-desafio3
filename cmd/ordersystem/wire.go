//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/leandrojbraga/goexpert-desafio3/internal/entity"
	"github.com/leandrojbraga/goexpert-desafio3/internal/event"
	"github.com/leandrojbraga/goexpert-desafio3/internal/infra/database"
	"github.com/leandrojbraga/goexpert-desafio3/internal/infra/web"
	uc_ordercreate "github.com/leandrojbraga/goexpert-desafio3/internal/usecase/order/create"
	uc_orderlist "github.com/leandrojbraga/goexpert-desafio3/internal/usecase/order/list"
	"github.com/leandrojbraga/goexpert-desafio3/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *uc_ordercreate.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		uc_ordercreate.NewCreateOrderUseCase,
	)
	return &uc_ordercreate.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB) *uc_orderlist.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		uc_orderlist.NewListOrderUseCase,
	)
	return &uc_orderlist.ListOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
