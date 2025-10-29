package repository

import (
	"database/sql"
	"fmt"
	"log"
	"ordersystem/model"
	"os"

	_ "github.com/lib/pq"
)

// Repository wraps the SQL connection.
type Repository struct {
	Conn *sql.DB
}

// Connect initializes a new database connection using environment variables.
func Connect() *Repository {
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("POSTGRES_TCP_PORT", "5432")
	user := getenv("POSTGRES_USER", "docker")
	password := getenv("POSTGRES_PASSWORD", "docker")
	dbname := getenv("POSTGRES_DB", "order")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Database not reachable: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully")

	return &Repository{Conn: db}
}

func getenv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

// GetDrinks returns all drinks from the database.
func (r *Repository) GetDrinks() ([]model.Drink, error) {
	rows, err := r.Conn.Query(`SELECT id, name, price, description FROM drinks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drinks []model.Drink
	for rows.Next() {
		var d model.Drink
		if err := rows.Scan(&d.ID, &d.Name, &d.Price, &d.Description); err != nil {
			return nil, err
		}
		drinks = append(drinks, d)
	}

	return drinks, rows.Err()
}

// GetOrders returns all orders from the database.
func (r *Repository) GetOrders() ([]model.Order, error) {
	rows, err := r.Conn.Query(`SELECT drink_id, created_at, amount FROM orders ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.DrinkID, &o.CreatedAt, &o.Amount); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, rows.Err()
}

// GetTotalledOrders returns total quantity per drink.
func (r *Repository) GetTotalledOrders() (map[uint64]uint64, error) {
	rows, err := r.Conn.Query(`SELECT drink_id, SUM(amount) FROM orders GROUP BY drink_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totals := make(map[uint64]uint64)
	for rows.Next() {
		var drinkID, sum uint64
		if err := rows.Scan(&drinkID, &sum); err != nil {
			return nil, err
		}
		totals[drinkID] = sum
	}

	return totals, rows.Err()
}

// AddOrder inserts a new order into the database.
func (r *Repository) AddOrder(order *model.Order) error {
	_, err := r.Conn.Exec(`
		INSERT INTO orders (drink_id, amount, created_at)
		VALUES ($1, $2, NOW())
	`, order.DrinkID, order.Amount)
	return err
}
