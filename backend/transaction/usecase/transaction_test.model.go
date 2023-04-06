package usecase

import "io/ioutil"

func mountInMemoryFileBytes(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}

	return data
}

var inMemoryFileMetadata = &CreateFileMetadata{
	BinaryData: mountInMemoryFileBytes("../.localenv/sales.txt"),
}
