package usecase

import (
	"time"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/shopspring/decimal"
)

type TransactionLine struct {
	ID          string
	TType       e.TransactionTypeEnum
	TDate       time.Time
	ProductName string
	Amount      decimal.Decimal
	SellerName  string
}

type SellerBalance struct {
	ID        string
	SellerID  string
	Balance   string
	CreatedAt time.Time
}

type CreateFileMetadata struct {
	FileSize    string
	Disposition string
	BinaryData  []byte
}
