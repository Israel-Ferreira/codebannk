package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Israel-Ferreira/codebank/domain"
	"github.com/Israel-Ferreira/codebank/infra/kafka"
	"github.com/Israel-Ferreira/codebank/infra/repository"
	"github.com/Israel-Ferreira/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()

	defer db.Close()

	repo := repository.NewTxnRepository(db)

	cc := domain.NewCreditCard()
	cc.Cvv = 255
	cc.Number = "5161 6129 1985 9781"
	cc.ExpirationMonth = 06
	cc.ExpirationYear = 2022

	if err := repo.CreateCreditCard(*cc); err != nil {
		fmt.Println(err.Error())
		log.Fatalln("Error ao criar o cartão de credito")
	}

	fmt.Println("Gravou no Banco!!!")

}

func SetupTxnUseCase(db *sql.DB) usecase.UseCaseTxn {
	txnRepositoryDB := repository.NewTxnRepository(db)


	return usecase.NewUseCaseTxn(txnRepositoryDB, kafka.NewKafkaProducer())
}





func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("Error on connecting database")
	}

	return db

}
