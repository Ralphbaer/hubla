package entity

type FileTransaction struct {
	ID            string `json:"id"`
	FileID        string `json:"file_id"`
	TransactionID string `json:"transaction_id"`
}
