package usecase

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
)

// SalesUseCase represents a collection of use cases for sales operations
type SalesUseCase struct {
	SalesRepo r.SalesRepository
}

// StoreFileContent stores a new Sales
func (uc *SalesUseCase) StoreFileContent(ctx context.Context, sfp *SalesFileUpload) (*e.Sales, error) {
	// Read the file content
	data, err := io.ReadAll(sfp.File)
	if err != nil {
		return nil, err
	}
	// Process the file content
	salesEntries, err := processFileData(ctx, data)
	if err != nil {
		return nil, err
	}

	if _, err := uc.SalesRepo.Save(ctx, salesEntries); err != nil {
		return nil, err
	}

	// armazenar

	fmt.Println(salesEntries)

	return nil, nil
}

func processFileData(ctx context.Context, data []byte) ([]e.Sales, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var entries []e.Sales

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

func parseLine(line string) (e.Sales, error) {
	var entry e.Sales

	if len(line) < 70 {
		return entry, fmt.Errorf("invalid line format")
	}

	entry.ID = uuid.New().String()

	ttype, err := strconv.ParseUint(line[:1], 10, 8)
	if err != nil {
		return entry, fmt.Errorf("error parsing code: %v", err)
	}
	entry.TType = uint8(ttype)

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])
	if err != nil {
		return entry, fmt.Errorf("error parsing date: %v", err)
	}
	entry.TDate = date

	entry.ProductDescription = strings.TrimSpace(line[26:50])

	value, err := strconv.Atoi(strings.TrimSpace(line[50:66]))
	if err != nil {
		return entry, fmt.Errorf("error parsing amount: %v", err)
	}
	entry.Amount = value

	entry.Seller = strings.TrimSpace(line[66:])

	return entry, nil
}
