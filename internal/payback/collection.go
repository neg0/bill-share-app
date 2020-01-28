package payback

type Collection []Reader

// NewCollection creates a new instance of collection of PayBack object
func NewCollection() Collection {
	var instance Collection
	return instance
}

// Add will insert a new PayBack record into the collection
func (c *Collection) Add(payback Reader) Collection {
	*c = append(*c, payback)
	return *c
}
