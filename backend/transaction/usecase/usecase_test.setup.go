package usecase

import (
	"testing"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hmock"

	"github.com/Ralphbaer/hubla/backend/transaction/gen/mocks"

	"github.com/golang/mock/gomock"
)

func setupTransactionUseCaseMocks(mockCtrl *gomock.Controller, t *testing.T) TransactionUseCase {
	return TransactionUseCase{
		FileMetadataRepo:    setupFileMetadataRepo(mockCtrl),
		SellerRepo:          nil,
		ProductRepo:         nil,
		TransactionRepo:     nil,
		FileTransactionRepo: nil,
		SellerBalanceRepo:   nil,
	}
}

func setupFileMetadataRepo(mockCtrl *gomock.Controller) *mocks.MockFileMetadataRepository {
	m := mocks.NewMockFileMetadataRepository(mockCtrl)
	m.
		EXPECT().
		FindByHash(gomock.Any(), gomock.Eq("71a1ae20f8bb23ccbc15a161364c238fe7a6360a07a26dfb2818584692c77403")).
		Return(nil, nil).
		AnyTimes()
	m.
		EXPECT().
		FindByHash(gomock.Any(), gomock.Eq("randomHash")).
		Return(nil, common.EntityNotFoundError{}).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), hmock.FieldValueMatcher("Hash", "71a1ae20f8bb23ccbc15a161364c238fe7a6360a07a26dfb2818584692c77403")).
		Return(nil, ErrFileMetadataAlreadyExists).
		AnyTimes()
	m.
		EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	return m
}
