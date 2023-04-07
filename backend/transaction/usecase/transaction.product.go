package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/google/uuid"
)

func (uc *TransactionUseCase) findOrCreateProduct(ctx context.Context, productName string, products map[string]*e.Product, sellerID string) (*e.Product, error) {
	product, err := uc.findProduct(ctx, productName, products)
	if err != nil {
		return nil, err
	}
	if product == nil {
		product, err = uc.createProduct(ctx, productName, sellerID)
		if err != nil {
			return nil, err
		}
		products[productName] = product
	}

	return product, nil
}

func (uc *TransactionUseCase) findProduct(ctx context.Context, productName string, products map[string]*e.Product) (*e.Product, error) {
	if product, found := products[productName]; found {
		return product, nil
	}
	product, err := uc.ProductRepo.FindByProductName(ctx, productName)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (uc *TransactionUseCase) createProduct(ctx context.Context, productName string, sellerID string) (*e.Product, error) {
	product := &e.Product{
		ID:        uuid.NewString(),
		Name:      productName,
		CreatorID: sellerID,
	}
	if err := uc.ProductRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
