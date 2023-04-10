package usecase

import (
	"time"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/shopspring/decimal"
)

// CreateTransaction is a struct representing a transaction that can be created,
// including the transaction ID, type, date, product name, amount, and seller name.
type CreateTransaction struct {
	ID          string
	TType       e.TransactionTypeEnum
	TDate       time.Time
	ProductName string
	Amount      decimal.Decimal
	SellerName  string
}

// SellerBalance is a struct representing a seller's balance,
// including the balance ID, seller ID, balance amount, and creation time.
type SellerBalance struct {
	ID        string
	SellerID  string
	Balance   string
	CreatedAt time.Time
}

// CreateFileMetadata is a struct representing file metadata that can be created,
// including the file ID, size, disposition, and binary data.
type CreateFileMetadata struct {
	ID          string
	FileSize    string
	Disposition string
	BinaryData  []byte
}
