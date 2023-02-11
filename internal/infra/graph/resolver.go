package graph

import (
	uc_ordercreate "github.com/leandrojbraga/goexpert-desafio3/internal/usecase/order/create"
	uc_orderlist "github.com/leandrojbraga/goexpert-desafio3/internal/usecase/order/list"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase uc_ordercreate.CreateOrderUseCase
	ListOrderUseCase   uc_orderlist.ListOrderUseCase
}
