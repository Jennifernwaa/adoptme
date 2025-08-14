package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

var Supabase *supa.Client
var DB *pgx.Conn

func ConnectToDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on environment variables")
	}

	// Get the database URL
	dbURL := os.Getenv("DATABASE_URL") + "?sslmode=require"
	if dbURL == "" {
		panic("DATABASE_URL is not set")
	}

	// Create connection pool
	pool, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	DB = pool
	fmt.Println("✅ Connected to Supabase Postgres (Session Pooler)")

}
