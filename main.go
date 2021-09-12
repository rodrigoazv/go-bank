package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rodrigoazv/go-bank/domain"
	"github.com/rodrigoazv/go-bank/infra/repository"
	"github.com/rodrigoazv/go-bank/usecases"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Name = "Rodrigo"
	cc.Number = "1234"
	cc.ExpirationYear = 22
	cc.ExpirationMonth = 06
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreatedCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecases.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecases.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("err connection to database")
	}

	return db
}
