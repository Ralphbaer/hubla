package repository

import (
	"strings"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SalesPostgreSQLModel is the model of entity.Sales
type SalesPostgreSQLModel struct {
	ID string `bson:"_id,omitempty"`
}

// ToEntity converts a SalesPostgreSQLModel to e.Sales
func (d *SalesPostgreSQLModel) ToEntity() *e.Sales {
	sales := &e.Sales{
		ID: "a",
	}

	return sales
}

// FromEntity converts an entity.Sales to SalesPostgreSQLModel
func (d *SalesPostgreSQLModel) FromEntity(sales *e.Sales) error {
	if strings.TrimSpace(sales.ID) != "" {
	}

	return nil
}
