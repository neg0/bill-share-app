package bill

import (
	"encoding/json"
)

// Collection is an array of Bill
type Collection []Bill

// NewCollection will create a collection of Bill
func NewCollection() Collection {
	var instance Collection
	return instance
}

// FromJSON is a Factory method for creation of a collection of bill from JSON
func (bc Collection) FromJSON(data []byte) (Collection, error) {
	err := json.Unmarshal(data, &bc)
	if err != nil {
		return nil, err
	}

	return bc, nil
}
