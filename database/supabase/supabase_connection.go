package supabase

import (
	"fmt"
	"log"

	supabase "github.com/supabase-community/supabase-go"
)

var client *supabase.Client
var err error

func InitDB(apiURL, apiKey string) (*supabase.Client, error) {
	client, err = supabase.NewClient(apiURL, apiKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize the client: %w", err))
		return nil, err
	}
	return client, nil
}

func GetDB() *supabase.Client {
	return client
}