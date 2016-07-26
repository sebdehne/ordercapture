package apiv1

import (
	"github.com/sebdehne/structmapping"
	"github.com/sebdehne/ordercapture/domain/v1"
)

var mapper = structmapping.New(structmapping.DstFieldBased, false)

func init() {
	mapper.Add(func(src Individual, dst *v1.Individual) {
		dst.FirstName = src.FirstName
		dst.Surname = src.LastName
		dst.ContactEmailAddress = src.ContactEmailAddress
	})
}

// Converts the API version to the most resent internal domain version
func MapToInternal(o OrderDraft) v1.OrderDraft {
	result := new(v1.OrderDraft)
	mapper.Map(&o, result)
	return *result
}
