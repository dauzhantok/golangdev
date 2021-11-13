package main

import (
	"context"
	"gomidka/internal/http"
	"gomidka/internal/store/postgres"

)

func main() {

	urlExample := "postgres://postgres:Dauzhan02@localhost:5432/bread"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}