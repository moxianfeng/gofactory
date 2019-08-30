package gofactory

import (
	"fmt"
	"reflect"
)

type factory struct {
	container map[string]interface{}
}

var (
	E_NOT_FOUND        = fmt.Errorf("Not found")
	E_INVALID_ARGUMENT = fmt.Errorf("Invalid argument")
	Default            = NewFactory()
)

func NewFactory() *factory {
	return &factory{container: make(map[string]interface{})}
}

func (this *factory) Register(name string, val interface{}) {
	this.container[name] = val
}

func (this *factory) Get(name string, val interface{}) error {
	v, ok := this.container[name]
	if !ok {
		return E_NOT_FOUND
	}

	if reflect.TypeOf(val).Kind() != reflect.Ptr {
		return E_INVALID_ARGUMENT
	}
	outputVal := reflect.ValueOf(val).Elem()
	if outputVal.Type() != reflect.TypeOf(v) {
		return E_INVALID_ARGUMENT
	}
	outputVal.Set(reflect.ValueOf(v))
	return nil
}

func (this *factory) GetInterface(name string) (interface{}, error) {
	v, ok := this.container[name]
	if !ok {
		return nil, E_NOT_FOUND
	}

	return v, nil
}
