package persistence

import (
	"testing"
	"github.com/sebdehne/ordercapture/domain/v1"
	"github.com/stretchr/testify/assert"
)

func TestPersistenceService(t *testing.T) {
	assert := assert.New(t)

	c := NewInMemClient()
	p := New(c)

	o := genValidOrder()
	ok, err := p.StoreOrder("first", o, 0)
	assert.Nil(err)
	assert.True(ok)

	// storing it again against version 0 should not succeed
	ok, err = p.StoreOrder("first", o, 0)
	assert.Nil(err)
	assert.False(ok)

	// get it back from the store
	back, version, err := p.GetOrder("first")
	assert.Nil(err)
	assert.NotNil(back)
	assert.Equal(int64(1), version)

	// check if order we received
	assert.Equal("1234", back.ShipTo.Address.PostalCode)

	// delete the record now, but with the wrong version
	ok, err = p.DeleteOrder("first", 2)
	assert.Nil(err)
	assert.False(ok)

	// now with the correct version
	ok, err = p.DeleteOrder("first", 1)
	assert.Nil(err)
	assert.True(ok)

	// deleting it again should not succeed
	ok, err = p.DeleteOrder("first", 1)
	assert.Nil(err)
	assert.False(ok)
}

func genValidOrder() v1.OrderDraft {
	i := v1.Individual{FirstName:"Peter", Surname:"Peterson", ContactEmailAddress:"email@example.org"}

	a := v1.Address{
		StreetName:"Streetname",
		StreetNumber:"Streetnumber",
		PostalCode:"1234",
		City:"City",
		Country:"NOR"}

	o := v1.OrderDraft{}
	o.Items = make(map[string]v1.OrderItem)
	o.Orderer = i
	o.TotalAmount = v1.Money{Kroner:0, Ore:0}
	o.ShipTo = v1.ShipTo{Contact:i, Address:a}
	o.InvoiceTo = v1.InvoiceTo{Contact:i, Email:"a@a.com"}
	o.Items["key1234"] = v1.OrderItem{CatalogueId:"HW1", Quantity:1, ItemPrice:v1.Money{Kroner:1, Ore:0}}
	return o
}
