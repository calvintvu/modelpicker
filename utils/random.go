package random

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func RandomElement[T any](array []T) (T, error) {
	var zero T
	if len(array) == 0 {
		return zero, errors.New("empty slice")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(array))))
	if err != nil {
		return zero, err
	}
	return array[n.Int64()], nil
}
