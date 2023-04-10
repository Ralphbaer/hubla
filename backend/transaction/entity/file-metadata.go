package entity

import "time"

// FileMetadata is a struct containing information about a file,
// including ID, file size, disposition, hash, binary data, and creation time.
type FileMetadata struct {
	ID          string
	FileSize    int
	Disposition string
	Hash        string
	BinaryData  []byte
	CreatedAt   time.Time
}
