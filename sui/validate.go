// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

const ValidSuiAddressLength = 66

var validate *defaultValidator

type defaultValidator struct {
	validate *validator.Validate
	once     sync.Once
}

func init() {
	validate = &defaultValidator{
		validate: validator.New(),
	}
	validate.registerCheckFn("checkAddress", checkAddress)
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return v.handleErr(err)
		}
	}
	return nil
}

func (v *defaultValidator) handleErr(err error) error {
	validationErrs := err.(validator.ValidationErrors)
	for _, e := range validationErrs {
		if e.Tag() == "lte" {
			err = fmt.Errorf("%v, field `%s` must be less than or equal to %s", err, e.Field(), e.Param())
		} else if e.Tag() == "gte" {
			err = fmt.Errorf("%v, field `%s` must be greater than or equal to %s", err, e.Field(), e.Param())
		}
	}
	return err
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func (v *defaultValidator) registerCheckFn(tag string, fn validator.Func) {
	if err := v.validate.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("register validation %s failed", tag)
	}
}

func checkAddress(fl validator.FieldLevel) bool {
	return isHex(fl.Field().String()) && len(fl.Field().String()) == ValidSuiAddressLength
}

func isHex(value string) bool {
	pattern := "^(0x|0X)?[a-fA-F0-9]+$"
	rex := regexp.MustCompile(pattern)
	ok := rex.Match([]byte(value))
	if ok {
		length := len(value)
		if length%2 != 0 {
			return false
		}
	}
	return true
}
