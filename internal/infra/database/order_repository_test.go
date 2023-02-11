package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/leandrojbraga/goexpert-desafio3/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	fmt.Println("TearDownTest")
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func createOrder(suite *OrderRepositoryTestSuite) *entity.Order {
	order, err := entity.NewOrder(uuid.New().String(), 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	return order
}

func (suite *OrderRepositoryTestSuite) TestGetAllOrders() {
	repo := NewOrderRepository(suite.Db)
	orders, err := repo.GetAll()
	suite.NoError(err)
	suite.Equal(0, len(orders))

	order := createOrder(suite)

	orders, err = repo.GetAll()
	suite.NoError(err)
	suite.Equal(1, len(orders))
	suite.Equal(order.ID, orders[0].ID)
	suite.Equal(order.Price, orders[0].Price)
	suite.Equal(order.Tax, orders[0].Tax)
	suite.Equal(order.FinalPrice, orders[0].FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {

	order := createOrder(suite)

	var orderResult entity.Order
	err := suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
