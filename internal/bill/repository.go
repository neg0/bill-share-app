package bill

import (
	"bsocial/pkg/bill"

	"errors"
	"io/ioutil"
	"os"
)

type CollectionFinder interface {
	FindAll() bill.Collection
}

type Repository struct {
	collection bill.Collection
}

func NewRepository() (*Repository, error) {
	filePath := os.Getenv("FILE_PATH_BILLS")
	if filePath == "" {
		return nil, errors.New("`FILE_PATH_BILLS` hasn't been defined")
	}

	collection, err := fromFile(filePath)
	if err != nil {
		return nil, err
	}

	return &Repository{collection}, nil
}

func (r *Repository) FindAll() bill.Collection {
	return r.collection
}

func fromFile(filePath string) (bill.Collection, error) {
	jsonCollection, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return bill.NewCollection().FromJSON(jsonCollection)
}
