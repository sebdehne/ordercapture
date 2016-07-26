package domain

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"github.com/sebdehne/ordercapture/domain/v1"
)

var validators map[string]*validator.Validate

func init() {
	validators = make(map[string]*validator.Validate)

	orderValidator := validator.New(&validator.Config{TagName: "valid"})
	atLeastOne := mapValidation{minLen:1, maxLen:1000, keyValidation:"min=1,max=100,alphanum"}
	orderValidator.RegisterStructValidation(atLeastOne.mapValidator, v1.Order{})
	validators[reflect.TypeOf(v1.Order{}).Name()] = orderValidator

	orderDraftValidator := validator.New(&validator.Config{TagName: "valid"})
	zeroOreMore := mapValidation{minLen:1, maxLen:1000, keyValidation:"min=0,max=100,alphanum"}
	orderDraftValidator.RegisterStructValidation(zeroOreMore.mapValidator, v1.OrderDraft{})
	validators[reflect.TypeOf(v1.OrderDraft{}).Name()] = orderDraftValidator
}

func Validate(o interface{}) error {
	if v, ok := validators[reflect.TypeOf(o).Name()]; !ok {
		panic("Could not find validator for " + reflect.TypeOf(o).Name())
	} else {
		return v.Struct(o)
	}
}

type mapValidation struct {
	minLen        int
	maxLen        int
	keyValidation string
}

func (mv *mapValidation) mapValidator(v *validator.Validate, structLevel *validator.StructLevel) {

	for i := 0; i < structLevel.CurrentStruct.NumField(); i++ {

		if structLevel.CurrentStruct.Field(i).Kind() != reflect.Map {
			continue
		}

		mapField := structLevel.CurrentStruct.Field(i)
		mapFieldName := structLevel.CurrentStruct.Type().Field(i).Name

		// min length validation
		if mapField.Len() < mv.minLen {
			structLevel.ReportError(mapField, mapFieldName, mapFieldName, "min")
		}

		// max length validation
		if mapField.Len() > mv.maxLen {
			structLevel.ReportError(mapField, mapFieldName, mapFieldName, "max")
		}

		// key/value validation
		for _, key := range mapField.MapKeys() {
			if err := v.Field(key.String(), "min=1,max=100,alphanum"); err != nil {
				structLevel.ReportValidationErrors(mapFieldName + "." + key.String(), err.(validator.ValidationErrors))
			}

			if err := v.Struct(mapField.MapIndex(key).Interface()); err != nil {
				structLevel.ReportValidationErrors(mapFieldName + "." + key.String() + ".", err.(validator.ValidationErrors))
			}
		}

	}
}