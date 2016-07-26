package v2

type OrderDraft struct {
	Orderer     Individual

	Items       map[string]OrderItem
	TotalAmount Money  `valid:"required"`

	ShipTo      ShipTo
	InvoiceTo   InvoiceTo
}

type Order struct {
	Orderer     Individual `valid:"required"`

	Items       map[string]OrderItem
	TotalAmount Money      `valid:"required"`

	ShipTo      ShipTo     `valid:"required"`
	InvoiceTo   InvoiceTo  `valid:"required"`
}

func NewOrder() Order {
	o := Order{}
	o.Items = make(map[string]OrderItem)
	return o
}

type OrderItem struct {
	CatalogueId string `valid:"min=1,max=10"`
	Quantity    int    `valid:"min=1,max=1000"`
	ItemPrice   Money  `valid:"required"`
}

type Money struct {
	Kroner int `valid:"min=0,max=1000000"`
	Ore    int `valid:"min=0,max=99"`
}

type Individual struct {
	FirstName           string `valid:"min=0,max=100"`
	Surname             string `valid:"min=1,max=100"`
	ContactEmailAddress string `valid:"email,min=1,max=100"`
}

type ShipTo struct {
	Contact Individual `valid:"required"`
	Address Address    `valid:"required"`
}

type InvoiceTo struct {
	Contact Individual
	Email   string `valid:"email,min=1,max=100"`
}

type Address struct {
	StreetName   string `valid:"min=1,max=100"`
	StreetNumber string `valid:"min=0,max=100"`
	PostalCode   string `valid:"numeric,min=4,max=10"`
	City         string `valid:"min=1,max=100"`
	Country      string `valid:"alpha,len=3"`
}
