package configs

import (
	"demerzel-events/internal/db"

	"github.com/joho/godotenv"
)

func Load() {
	// Load env variables
	_ = godotenv.Load()

	// Setup database connection
	db.SetupDB()
}
