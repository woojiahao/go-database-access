package setup

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"woojiahao.com/gda/internal/utility"
)

type customer struct {
	id      string
	name    string
	allergy sql.NullString
}

func Setup() {
	db, err := sql.Open("pgx", utility.ConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database cannot be reached because %s", err)
	}

	createQuery := `
	CREATE TABLE IF NOT EXISTS customer (
	    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	    name TEXT NOT NULL,
	    allergy TEXT
	);

	CREATE TABLE IF NOT EXISTS "order" (
	    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	    food TEXT NOT NULL,
	    quantity INTEGER NOT NULL,
	    timestamp TIMESTAMP NOT NULL DEFAULT now(),
	    customer_id UUID NOT NULL,
	    FOREIGN KEY(customer_id) REFERENCES customer(id)
	);
	`
	_, err = db.ExecContext(context.TODO(), createQuery)
	if err != nil {
		log.Fatalf("Cannot create tables because %s", err)
	}

	var customers []customer
	createCustomerQuery := `
	INSERT INTO customer(name, allergy) 
	VALUES
	    ('John Doe', null), 
	    ('Mary Anne', 'Cheese'), 
	    ('Jason Borne', null) 
	RETURNING *;
	`
	rows, err := db.QueryContext(context.TODO(), createCustomerQuery)
	for rows.Next() {
		var c customer
		rows.Scan(&c.id, &c.name, &c.allergy)
		customers = append(customers, c)
	}

	var johnDoeId, maryAnneId, JasonBorneId string
	for _, c := range customers {
		switch c.name {
		case "John Doe":
			johnDoeId = c.id
		case "Mary Anne":
			maryAnneId = c.id
		case "Jason Borne":
			JasonBorneId = c.id
		}
	}

	johnDoeOrderQuery := `
	INSERT INTO "order"(food, quantity, customer_id) 
	VALUES
	    ('Pie', 2, $1), 
	    ('Soup of the Day', 1, $1), 
	    ('Pudding', 2, $1);
	`
	_, err = db.ExecContext(context.TODO(), johnDoeOrderQuery, johnDoeId)
	if err != nil {
		log.Fatalf("Failed to insert John Doe's order because %s", err)
	}

	maryAnneOrderQuery := `
	INSERT INTO "order"(food, quantity, customer_id)
	VALUES
		('Fish and Chips', 1, $1),
		('Soup of the Day', 1, $1);
	`
	_, err = db.ExecContext(context.TODO(), maryAnneOrderQuery, maryAnneId)
	if err != nil {
		log.Fatalf("Failed to insert Mary Anne's order because %s", err)
	}

	jasonBorneOrderQuery := `
	INSERT INTO "order"(food, quantity, customer_id)
	VALUES
		('Pie', 3, $1),
		('Pudding', 3, $1);
	`
	_, err = db.ExecContext(context.TODO(), jasonBorneOrderQuery, JasonBorneId)
	if err != nil {
		log.Fatalf("Failed to insert Jason Borne's order because %s", err)
	}
}
