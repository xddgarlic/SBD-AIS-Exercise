package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// NewDatabaseHandler initializes the in-memory database with test data.
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Espresso", Price: 2.50, Description: "Strong and rich coffee shot"},
		{ID: 2, Name: "Cappuccino", Price: 3.20, Description: "Espresso with steamed milk and foam"},
		{ID: 3, Name: "Latte", Price: 3.50, Description: "Smooth blend of espresso and milk"},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now().Add(-2 * time.Hour), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now().Add(-1 * time.Hour), Amount: 1},
		{DrinkID: 1, CreatedAt: time.Now().Add(-30 * time.Minute), Amount: 3},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// GetTotalledOrders calculates total orders per drink.
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	return totalledOrders
}

// AddOrder adds a new order to the database.
func (db *DatabaseHandler) AddOrder(order *model.Order) {
	if order == nil {
		return
	}
	// If CreatedAt is zero, set it to now
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Now()
	}
	db.orders = append(db.orders, *order)
}
