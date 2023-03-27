package usecase

import "mime/multipart"

// CreateSalesInput is the set of information that will be used to enter data through our handlers.
// We can understand it as a Command. It is used in CREATE operations.
// swagger:model SalesFileUpload
type SalesFileUpload struct {
	File multipart.File
}
