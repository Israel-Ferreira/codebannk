package server

import (
	"log"
	"net"

	"github.com/Israel-Ferreira/codebank/infra/grpc/pb"
	txnService "github.com/Israel-Ferreira/codebank/infra/grpc/services"
	"github.com/Israel-Ferreira/codebank/usecase"
	"google.golang.org/grpc"
)


type GRPCServer struct {
	ProcessTxnUseCase usecase.UseCaseTxn
}


func NewGRPCServer() GRPCServer{
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalln("could not listen tcp port")
	}

	
	transactionService := txnService.NewTransactionService()
	transactionService.ProcessTxnUseCase = g.ProcessTxnUseCase
	
	gServer  := grpc.NewServer();
	pb.RegisterPaymentServiceServer(gServer, transactionService)


}