package db

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

func InitDB() *supabase.Client {
	_ = godotenv.Load()

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	client := supabase.CreateClient(supabaseURL, supabaseKey)

	return client
}
