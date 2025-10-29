package repository

import "log"

// InitSchema creates tables if they don’t exist and inserts sample data.
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

	seed := `
    INSERT INTO drinks (name, price, description)
    VALUES
        ('Espresso', 2.50, 'That's that me espresso),
        ('Latte', 3.20, 'Milk coffee with foam'),
        ('Iced Tea', 2.00, 'Refreshing cold tea')
    ON CONFLICT DO NOTHING;
    `
	_, err = r.Conn.Exec(seed)
	if err != nil {
		log.Fatalf("❌ Error seeding data: %v", err)
	}

	log.Println("✅ Seed data inserted (drinks)")
}
