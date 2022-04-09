package Consumer

type ConsumerName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ConsumerAddress struct {
	Street      string `json:"street"`
	Street2     string `json:"street2"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode	"`
	PhoneNumber string `json:"phoneNumber"`
}

type Consumer struct {
	ConsumerID string          `json:"consumer_id"` // Hash of the rest of this struct instance
	Name       ConsumerName    `json:"consumer_name"`
	Address    ConsumerAddress `json:"address"`
	PayerId    uint64
}
