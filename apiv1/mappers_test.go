package apiv1

import (
	"testing"
	"fmt"
)

func TestMapToInternal(t *testing.T) {
	od := OrderDraft{}
	odInternal := MapToInternal(od)
	fmt.Println(odInternal)
}
