package repository

import (
	"strings"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

type TransactionSellerResult struct {
	Seller e.Seller
	Status string
	Error  error
}

// SalesPostgreSQLModel is the model of entity.Sales
type SalesPostgreSQLModel struct {
	ID string `bson:"_id,omitempty"`
}

// ToEntity converts a SalesPostgreSQLModel to e.Sales
func (d *SalesPostgreSQLModel) ToEntity() *e.Transaction {
	transaction := &e.Transaction{
		ID: "a",
	}

	return transaction
}

// FromEntity converts an entity.Sales to SalesPostgreSQLModel
func (d *SalesPostgreSQLModel) FromEntity(transaction *e.Transaction) error {
	if strings.TrimSpace(transaction.ID) != "" {
	}

	return nil
}
