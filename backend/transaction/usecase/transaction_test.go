package usecase

import (
	"context"
	"testing"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionUseCase_StoreFileMetadata(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transactionUseCase := setupTransactionUseCaseMocks(mockCtrl, t)

	uuid := uuid.New().String()
	testCases := []struct {
		name        string
		createFile  *CreateFileMetadata
		expected    string
		expectedErr error
	}{
		{
			name: "Valid file metadata",
			createFile: &CreateFileMetadata{
				ID:          uuid,
				FileSize:    "100",
				Disposition: "test-disposition",
				BinaryData:  []byte("test-data"),
			},
			expected:    uuid,
			expectedErr: nil,
		},
		{
			name: "Entity conflict",
			createFile: &CreateFileMetadata{
				ID:          "1",
				FileSize:    "100",
				Disposition: "test-disposition",
				BinaryData:  []byte("test-data"),
			},
			expected:    "",
			expectedErr: ErrFileMetadataAlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := transactionUseCase.StoreFileMetadata(context.Background(), tc.createFile)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}

func TestTransactionUseCase_GetFileTransactions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transactionUseCase := setupTransactionUseCaseMocks(mockCtrl, t)

	testCases := []struct {
		name        string
		fileID      string
		expected    []*e.Transaction
		expectedErr error
	}{
		{
			name:        "Valid file ID",
			fileID:      "test-file-id",
			expected:    []*e.Transaction{validTransaction},
			expectedErr: nil,
		},
		{
			name:        "No content",
			fileID:      "123123",
			expected:    nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transactions, err := transactionUseCase.GetFileTransactions(context.Background(), tc.fileID)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				assert.Nil(t, transactions)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, transactions)
			}
		})
	}
}
