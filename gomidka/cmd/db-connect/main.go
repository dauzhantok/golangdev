package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

func main()  {
	ctx := context.Background()
	urlExample := "postgres://postgres:Dauzhan02@localhost:5432/bread"
	conn, err := pgx.Connect(ctx, urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	if err:= conn.Ping(ctx); err!= nil{
		panic(err)
	}
	log.Println("Pinged DB")
}
