package services

import (
	"context"

	"github.com/Israel-Ferreira/codebank/dto"
	"github.com/Israel-Ferreira/codebank/infra/grpc/pb"
	"github.com/Israel-Ferreira/codebank/usecase"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TxnService struct {
	ProcessTxnUseCase usecase.UseCaseTxn
}

func (t *TxnService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	txnDto := dto.TxnDto{
		Name:            in.GetCreditCard().Name,
		Number:          in.GetCreditCard().Number,
		ExpirationMonth: uint32(in.CreditCard.GetExpirationMonth()),
		ExpirationYear:  uint32(in.GetCreditCard().ExpirationYear),
		Cvv:             in.GetCreditCard().Cvv,
		Amount:          float64(in.GetAmount()),
		Status:          in.GetStatus(),
		Store:           in.GetStore(),
		Description:     in.GetDescription(),
	}

	transaction, err := t.ProcessTxnUseCase.ProcessTxn(txnDto)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != "APPROVED" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "Transaction Rejected by Bank")
	}

	return &empty.Empty{}, nil
}

func NewTransactionService() *TxnService {
	return &TxnService{}
}
