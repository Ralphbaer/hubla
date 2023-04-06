package usecase

import (
	"testing"

	"github.com/Ralphbaer/hubla/backend/transaction/gen/mock"

	"github.com/golang/mock/gomock"
)

func setupTransactionUseCaseMocks(mockCtrl *gomock.Controller, t *testing.T) TransactionUseCase {
	return TransactionUseCase{
		FileMetadataRepo:    nil,
		SellerRepo:          nil,
		ProductRepo:         nil,
		TransactionRepo:     nil,
		FileTransactionRepo: nil,
		SellerBalanceRepo:   nil,
	}
}

func setupFileMetadataRepo(mockCtrl *gomock.Controller) *mock.MockFileMetadataRepository {
	m := mock.NewMockFileMetadataRepository(mockCtrl)
	m.
		EXPECT().
		Find(gomock.Any(), hash).
		Return(nil, nil).
		AnyTimes()
	/*m.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		Return(happyScenario, nil).
		AnyTimes()
	m.
		EXPECT().
		FindRandom(gomock.Any()).
		Return(products[6], nil).
		AnyTimes()*/

	return m
}
