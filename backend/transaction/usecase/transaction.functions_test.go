package usecase

import (
	"testing"
	"time"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/shopspring/decimal"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		name        string
		line        string
		expected    *TransactionLine
		expectedErr error
	}{
		{
			name: "valid line",
			line: "42022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500THIAGO OLIVEIRA",
			expected: &TransactionLine{
				TType:       e.COMMISSION_RECEIVED,
				TDate:       mustParseTime("2022-01-16T14:13:54-03:00"),
				ProductName: "CURSO DE BEM-ESTAR",
				Amount:      decimal.NewFromInt(4500),
				SellerName:  "THIAGO OLIVEIRA",
			},
		},
		{
			name:        "invalid transaction type",
			line:        "92021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       0000050000MARIA CANDIDA",
			expectedErr: ErrInvalidTransactionType,
		},
		{
			name:        "error parsing date",
			line:        "1202X-12-03T11:46:02-03:00DESENVOLVEDOR FULL STACK      0000050000MARIA CANDIDA",
			expectedErr: ErrInvalidDate,
		},
		{
			name:        "error parsing amount",
			line:        "12021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       000005X000MARIA CANDIDA",
			expectedErr: ErrInvalidAmount,
		},
		{
			name:        "invalid seller name",
			line:        "12021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       0000050000          ",
			expectedErr: ErrInvalidSellerName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			line, err := parseLine(tc.line)
			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, but got nil", tc.expectedErr)
				}
				if tc.expectedErr.Error() != err.Error() {
					t.Fatalf("expected error %v, but got %v", tc.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !cmpTransactionLine(tc.expected, line) {
					t.Fatalf("expected transaction line %v, but got %v", tc.expected, line)
				}
			}
		})
	}
}

func mustParseTime(timeStr string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func cmpTransactionLine(tl1, tl2 *TransactionLine) bool {
	if tl1 == nil || tl2 == nil {
		return false
	}
	return tl1.TType == tl2.TType &&
		tl1.TDate.Equal(tl2.TDate) &&
		tl1.ProductName == tl2.ProductName &&
		tl1.Amount.Equal(tl2.Amount) &&
		tl1.SellerName == tl2.SellerName
}
