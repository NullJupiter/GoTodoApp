package helper

import "crypto/rand"

// GenerateRandomByteSlice function is used to generate a truly random byte slice for signing the cookies.
func GenerateRandomByteSlice(size int) ([]byte, error) {
	byteSlice := make([]byte, size)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}
