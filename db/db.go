package db

import (
	"github.com/nedpals/supabase-go"
)

const (
	supabaseURL = "https://olfhzqabncwfpyjqnxix.supabase.co"
	supabaseKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im9sZmh6cWFibmN3ZnB5anFueGl4Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTMzOTc3OTksImV4cCI6MjAyODk3Mzc5OX0.LfsNF4e_SPdIDYk8xjbaLx5lFhlpVQ3Rfd75ICH6lqQ"
)

func InitDB() *supabase.Client {
	client := supabase.CreateClient(supabaseURL, supabaseKey)

	return client
}
