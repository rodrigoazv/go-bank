package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rodrigoazv/go-bank/infra/grpc/server"
	"github.com/rodrigoazv/go-bank/infra/kafka"
	"github.com/rodrigoazv/go-bank/infra/repository"
	"github.com/rodrigoazv/go-bank/usecases"
)

func main() {
	db := setupDb()
	defer db.Close()

	producer := setupKakfa()
	processTransactionUseCase := setupTransactionUseCase(db, producer)

	fmt.Println("Rodando grp")

	serveGrpc(processTransactionUseCase)

}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecases.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecases.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}
func setupKakfa() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer("host.docker.internal:9094")
	return producer
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

func serveGrpc(processTransactionUseCase usecases.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase

	grpcServer.Serve()
}
