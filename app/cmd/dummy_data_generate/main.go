package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/joho/godotenv"

	"my_gql_server/infrastructures/dummy"
)

func main() {

	err := godotenv.Load("../.env")

	var PSQL_DBNAME string
	var PSQL_USER string
	var PSQL_PASS string

	if err != nil {
		log.Printf("読み込み出来ませんでした: %v", err)
	}

	PSQL_DBNAME = os.Getenv("PSQL_DBNAME")
	PSQL_USER = os.Getenv("PSQL_USER")
	PSQL_PASS = os.Getenv("PSQL_PASS")

	connectDB, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable", PSQL_DBNAME, PSQL_USER, PSQL_PASS))

	if err != nil {
		log.Fatalf("error : %s", err)
	}

	boil.SetDB(connectDB)
	boil.DebugMode = true

	ctx := context.Background()

	dummy.CreateDummyUsers(ctx, connectDB)
	dummy.UpsertDummyStatuses(ctx, connectDB)

	dummy.CreateDummyArticles(ctx, connectDB)

	dummy.UpsertDummyCategories(ctx, connectDB)

	dummy.AnnualItems(ctx, connectDB)

	dummy.CreateTodos(ctx, connectDB)

}
