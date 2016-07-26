package domain

import (
	"testing"
	"github.com/sebdehne/ordercapture/domain/v1"
)

func TestValidation(t *testing.T) {

	o := genValidOrder()

	if err := Validate(o); err != nil {
		t.Fatal(err)
	}
}

func genValidOrder() v1.Order {
	i := v1.Individual{FirstName:"Peter", Surname:"Peterson", ContactEmailAddress:"email@example.org"}

	a := v1.Address{
		StreetName:"Streetname",
		StreetNumber:"Streetnumber",
		PostalCode:"1234",
		City:"City",
		Country:"NOR"}

	o := v1.NewOrder()
	o.Orderer = i
	o.TotalAmount = v1.Money{Kroner:0, Ore:0}
	o.ShipTo = v1.ShipTo{Contact:i, Address:a}
	o.InvoiceTo = v1.InvoiceTo{Contact:i, Email:"a@a.com"}
	o.Items["key1234"] = v1.OrderItem{CatalogueId:"HW1", Quantity:1, ItemPrice:v1.Money{Kroner:1, Ore:0}}
	return o
}
