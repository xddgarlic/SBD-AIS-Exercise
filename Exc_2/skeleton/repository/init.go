package repository

import "log"

// InitSchema creates tables and seeds drinks.
// It clears the drinks table on each run to avoid duplicates.
func (r *Repository) InitSchema() {
	schema := `
    CREATE TABLE IF NOT EXISTS drinks (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        price NUMERIC(10,2) NOT NULL,
        description TEXT
    );

    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        drink_id INTEGER NOT NULL REFERENCES drinks(id) ON DELETE CASCADE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        amount INTEGER NOT NULL
    );
    `
	_, err := r.Conn.Exec(schema)
	if err != nil {
		log.Fatalf("❌ Error creating schema: %v", err)
	}

	log.Println("✅ Tables ensured (drinks, orders)")

	// Clear the drinks table before seeding
	_, err = r.Conn.Exec(`TRUNCATE TABLE drinks RESTART IDENTITY CASCADE`)
	if err != nil {
		log.Fatalf("❌ Error clearing drinks table: %v", err)
	}

	log.Println("✅ Drinks table cleared")

	seed := `
    INSERT INTO drinks (name, price, description)
    VALUES
        ('Espresso', 2.50, 'Is it that sweet? I guess so'),
        ('Latte', 3.20, 'Milk coffee with foam'),
        ('Iced Tea', 2.00, 'Refreshing cold tea');
    `
	_, err = r.Conn.Exec(seed)
	if err != nil {
		log.Fatalf("❌ Error seeding data: %v", err)
	}

	// Seed orders
	seedOrders := `
    INSERT INTO orders (drink_id, amount, created_at)
    VALUES
        (1, 2, NOW()),
        (2, 1, NOW()),
        (3, 3, NOW());
    `
	_, err = r.Conn.Exec(seedOrders)
	if err != nil {
		log.Fatalf("❌ Error seeding orders: %v", err)
	}

	log.Println("✅ Seed data inserted (orders)")

	log.Println("✅ Seed data inserted (drinks)")
}
