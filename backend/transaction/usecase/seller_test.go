package usecase

// func TestSellerUseCase_GetSellerBalanceByID(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	sellerUseCase := setupSellerUseCaseMocks(mockCtrl, t)

// 	testCases := []struct {
// 		name        string
// 		sellerID    string
// 		expected    *e.SellerBalanceView
// 		expectedErr error
// 	}{
// 		{
// 			name:     "Valid seller ID",
// 			sellerID: "test-seller-id",
// 			expected: &e.SellerBalanceView{
// 				SellerID:        "test-seller-id",
// 				AvailableAmount: decimal.NewFromFloat(100.00),
// 				HeldAmount:      decimal.NewFromFloat(50.00),
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:        "Invalid seller ID",
// 			sellerID:    "invalid-seller-id",
// 			expected:    nil,
// 			expectedErr: common.EntityNotFoundError{},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			balance, err := sellerUseCase.GetSellerBalanceByID(context.Background(), tc.sellerID)
// 			if tc.expectedErr != nil {
// 				assert.ErrorIs(t, err, tc.expectedErr)
// 				assert.Nil(t, balance)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tc.expected, balance)
// 			}
// 		})
// 	}
// }
