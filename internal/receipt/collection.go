package receipt

// Collection is an array of Receipt (Bill)
type Collection []API

// NewCollection will create a collection of Bill
func NewCollection() Collection {
	var instance Collection
	return instance
}

// Add will insert new receipt (Bill) inside the collection
func (c *Collection) Add(receipt API) Collection {
	*c = append(*c, receipt)
	return *c
}

// ParticipantList will return an array containing participants name
func (c Collection) ParticipantList() []string {
	if len(c) < 1 {
		return []string{}
	}
	return c[0].SpendersName()
}
