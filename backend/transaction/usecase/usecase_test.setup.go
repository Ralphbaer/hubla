package usecase

import (
	"testing"
	"time"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hmock"
	"github.com/shopspring/decimal"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/Ralphbaer/hubla/backend/transaction/gen/mocks"

	"github.com/golang/mock/gomock"
)

func setupTransactionUseCaseMocks(mockCtrl *gomock.Controller, t *testing.T) TransactionUseCase {
	return TransactionUseCase{
		FileMetadataRepo:    setupFileMetadataRepo(mockCtrl),
		SellerRepo:          setupSellerRepo(mockCtrl),
		ProductRepo:         nil,
		TransactionRepo:     setupTransactionRepo(mockCtrl),
		FileTransactionRepo: nil,
		SellerBalanceRepo:   nil,
	}
}

func setupFileMetadataRepo(mockCtrl *gomock.Controller) *mocks.MockFileMetadataRepository {
	m := mocks.NewMockFileMetadataRepository(mockCtrl)
	m.
		EXPECT().
		Save(gomock.Any(), hmock.FieldValueMatcher("ID", "1")).
		Return(ErrFileMetadataAlreadyExists).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	return m
}

func setupSellerRepo(mockCtrl *gomock.Controller) *mocks.MockSellerRepository {
	m := mocks.NewMockSellerRepository(mockCtrl)
	m.
		EXPECT().
		FindBySellerName(gomock.Any(), gomock.Eq("JOSE CARLOS")).
		Return(nil, nil).
		AnyTimes()
	m.
		EXPECT().
		FindBySellerName(gomock.Any(), gomock.Eq("randomName")).
		Return(nil, common.EntityNotFoundError{}).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), hmock.FieldValueMatcher("Name", "JOSE CARLOS")).
		Return(common.EntityConflictError{}).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	return m
}

func setupTransactionRepo(mockCtrl *gomock.Controller) *mocks.MockTransactionRepository {
	m := mocks.NewMockTransactionRepository(mockCtrl)
	m.
		EXPECT().
		Save(gomock.Any(), hmock.FieldValueMatcher("ID", "test-transaction-id")).
		Return(common.EntityConflictError{}).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()
	m.
		EXPECT().
		ListTransactionsByFileID(gomock.Any(), gomock.Eq("test-file-id")).
		Return([]*e.Transaction{
			validTransaction,
		}, nil).
		AnyTimes()
	m.
		EXPECT().
		ListTransactionsByFileID(gomock.Any(), gomock.Any()).
		Return(nil, nil).
		AnyTimes()

	return m
}

var validTransaction = &e.Transaction{
	ID:        "test-transaction-id",
	TType:     e.CREATOR_SALE,
	TDate:     time.Now(),
	ProductID: "test-product-id",
	Amount:    decimal.NewFromFloat(100.00),
	SellerID:  "test-seller-id",
}
