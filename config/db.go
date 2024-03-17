package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	server   string
	port     string
	user     string
	password string
	database string
	sslMode  string
)

func initConst() {
	server = os.Getenv("DB_HOSTNAME")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
	sslMode = os.Getenv("DB_SSLMODE")
}

func CreateCon() *sql.DB {
	initConst()
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%v", server, port, user, password, database, sslMode)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected", err.Error())
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
		order_id Serial PRIMARY KEY,
		customer_name VARCHAR(100) NOT NULL,
		ordered_at timestamptz default now(),
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
		);

		CREATE TABLE IF NOT EXISTS items (
			item_id SERIAL PRIMARY KEY,
			item_code VARCHAR(50) NOT NULL,
			description TEXT,
			quantity INT NOT NULL,
			order_id INT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			FOREIGN KEY (order_id) REFERENCES orders (order_id)
		);
	
	`) 
	if err != nil {
		panic(err)
	}
	return db
}
