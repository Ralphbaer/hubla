package entity

import "time"

// TransactionFile represents a product that can be created and sold by creators or affiliates.
type FileMetadata struct {
	ID          string
	FileSize    int
	Disposition string
	Hash        string
	BinaryData  []byte
	CreatedAt   time.Time
}
