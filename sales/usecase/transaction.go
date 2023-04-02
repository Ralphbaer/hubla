package usecase

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SalesUseCase represents a collection of use cases for sales operations
type SalesUseCase struct {
	SalesRepo         r.SalesRepository
	SellerRepo        r.SellerRepository
	ProductRepo       r.ProductRepository
	SellerBalanceRepo r.SellerBalanceRepository
}

// StoreFileContent stores a new Sales
func (uc *SalesUseCase) StoreFileContent(ctx context.Context, binaryData []byte) (*TransactionLine, error) {
	// Process the file content
	entries, err := processFileData(ctx, binaryData)
	if err != nil {
		return nil, err
	}
	// uc.SalesRepo.Save(ctx, entries)

	if err := uc.orchestrateAndPersist(ctx, entries); err != nil {
		return nil, err
	}
	// Get the productName

	// if _, err := uc.SalesRepo.Save(ctx, nil); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func processFileData(ctx context.Context, binaryData []byte) ([]TransactionLine, error) {
	var entries []TransactionLine

	scanner := bufio.NewScanner(bytes.NewReader(binaryData))
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			log.Printf("Error parsing line: %v", err)
			continue
		}
		entries = append(entries, entry)

		// Check if the context is cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return entries, nil
}

func parseLine(line string) (TransactionLine, error) {
	var entry TransactionLine

	if len(line) < 70 {
		return entry, fmt.Errorf("invalid line format")
	}

	entry.ID = uuid.New().String()

	ttype, err := strconv.ParseUint(line[:1], 10, 8)
	if err != nil {
		return entry, fmt.Errorf("error parsing code: %v", err)
	}
	entry.TType = e.TransactionTypeMap[uint8(ttype)]

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])
	if err != nil {
		return entry, fmt.Errorf("error parsing date: %v", err)
	}
	entry.TDate = date

	entry.ProductName = strings.TrimSpace(line[26:50])

	value, err := decimal.NewFromString(strings.TrimSpace(line[50:66]))
	if err != nil {
		return entry, fmt.Errorf("error parsing amount: %v", err)
	}
	entry.Amount = value

	entry.SellerName = strings.TrimSpace(line[66:])

	return entry, nil
}

func (uc *SalesUseCase) orchestrateAndPersist(ctx context.Context, entries []TransactionLine) error {
	for _, v := range entries {
		// Check if the seller already exists
		count, err := uc.SellerRepo.Count(ctx, v.SellerName)
		if err != nil {
			return err
		}
		if count == 0 {
			seller := &e.Seller{
				ID:         uuid.NewString(),
				Name:       v.SellerName,
				SellerType: e.TransactionTypeToSellerTypeMap[v.TType],
			}
			_, err := uc.SellerRepo.Save(ctx, seller)
			if err != nil {
				return err
			}
		}
		product := &e.Product{
			ID:   uuid.NewString(),
			Name: v.ProductName,
		}
		if seller.SellerType != e.CREATOR {
			product.CreatorID = sellerID
		}

		if err := uc.ProductRepo.Save(ctx, product); err != nil {
			fmt.Println(err)
			continue
		}
		transaction := &e.Transaction{
			ID:        uuid.NewString(),
			TType:     v.TType,
			TDate:     v.TDate,
			ProductID: product.ID,
			Amount:    v.Amount,
			SellerID:  sellerID,
		}
		t, err := uc.SalesRepo.Save(ctx, transaction)
		if err != nil {
			return err
		}
		sb := &e.SellerBalance{
			ID:       uuid.NewString(),
			SellerID: t.SellerID,
			Balance:  t.Amount,
		}

		newBalance, err := uc.SellerBalanceRepo.Upsert(ctx, sb)
		if err != nil {
			return err
		}

		fmt.Println(newBalance)
	}

	return nil
}

/*
// Every sum, is a sum to the product, every subtract, is a pay
func calculateSellerAmount(sales []e.Transaction) (psAmount []*e.ProductSeller, affiliateAmount []*e.Affiliate) {
	psAmount = []*e.ProductSeller{}

	for _, sale := range sales {
		switch sale.TType {
		case e.CREATOR_SALE, e.AFFILIATE_SALE:
			upsertCreatorAmount(&psAmount, sale)
		case e.COMMISSION_PAID:
			upsertCreatorAmount(&psAmount, sale)
		case e.COMMISSION_RECEIVED:
			addToAffiliateAmountMap(&psAmount, sale)
		}
	}

	return psAmount, affiliateAmount
}

func upsertProductSeller(psAmountList *[]*e.ProductSeller, sale e.Transaction, uType e.SellerTypeEnum) {
	for i, ps := range *psAmountList {
		if ps.Seller.Name == sale.SellerName ||
			(uType == e.AFFILIATE && ps.Product.Name == sale.ProductName) {
			// Update existing element
			(*psAmountList)[i].Seller = e.Seller{
				//	ID:       ps.Product.SellerID,
				Name:     sale.SellerName,
				UType:    uType,
				Balance:  calculateBalance(sale.Amount, ps.Seller.Balance, sale.TType),
				ParentID: ps.Seller.ParentID,
			}
			return
		}
	}

	// Insert new element
	uuidd := uuid.NewString()
	*psAmountList = append(*psAmountList, &e.ProductSeller{
		Seller: e.Seller{
			ID:      uuidd,
			Name:    sale.SellerName,
			UType:   uType,
			Balance: sale.Amount.Add(decimal.Zero),
		},
		Product: &e.Product{
			ID:   uuid.NewString(),
			Name: sale.ProductName,
			//	SellerID: uuidd,
		},
	})
}

func upsertCreatorAmount(psAmountList *[]*e.ProductSeller, sale e.Transaction) {
	upsertProductSeller(psAmountList, sale, e.CREATOR)
}

func addToAffiliateAmountMap(psAmountList *[]*e.ProductSeller, sale e.Transaction) {
	upsertProductSeller(psAmountList, sale, e.AFFILIATE)
}

func calculateBalance(amount decimal.Decimal, balance decimal.Decimal, tType e.TransactionTypeEnum) decimal.Decimal {
	if tType == e.COMMISSION_PAID {
		return balance.Sub(amount)
	}
	return balance.Add(amount)
}

/*
func addToAffiliateAmountMap(psAmountList *[]*e.ProductSeller, sale e.Sales) {
	found := false
	for i, ps := range *psAmountList {
		if ps.Seller.Name == sale.SellerName ||
			(sale.TType == e.AFFILIATE_SALE && ps.Product.Name == sale.ProductName) {
			// Update existing element
			(*psAmountList)[i].Seller = e.Seller{
				ID:       ps.Seller.ID,
				Name:     ps.Seller.Name,
				UType:    e.AFFILIATE,
				Balance:  ps.Seller.Balance,
				ParentID: ps.Seller.ParentID,
			}
			found = true
			break
		}
	}

	if !found {
		uuidd := uuid.NewString()
		*psAmountList = append(*psAmountList, &e.ProductSeller{
			Seller: e.Seller{
				ID:      uuidd,
				Name:    sale.SellerName,
				UType:   e.AFFILIATE,
				Balance: sale.Amount.Add(decimal.Zero),
			},
			Product: &e.Product{
				ID:       uuid.NewString(),
				Name:     sale.ProductName,
				SellerID: uuidd,
			},
		})
	}
}
*/
