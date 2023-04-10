package entity

// FileTransaction is a struct representing the relationship between
// a file and a transaction, containing the IDs of both entities.
type FileTransaction struct {
	ID            string `json:"id"`
	FileID        string `json:"file_id"`
	TransactionID string `json:"transaction_id"`
}
