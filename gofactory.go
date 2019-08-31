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

func (this *factory) GetObject(name string, val interface{}) error {
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

// For example I is an interface, you can call GetInterface in two way. factroy.GetInterface("xxxx", new(I)) or factory.GetInterface("xxxx", (*I)(nil))
func (this *factory) GetInterface(name string, val interface{}) (interface{}, error) {
	v, ok := this.container[name]
	if !ok {
		return nil, E_NOT_FOUND
	}

	targetType := reflect.TypeOf(val).Elem()

	if !reflect.TypeOf(v).Implements(targetType) {
		return nil, E_INVALID_ARGUMENT
	}

	return v, nil
}

func (this *factory) GetAny(name string) (interface{}, error) {
	v, ok := this.container[name]
	if !ok {
		return nil, E_NOT_FOUND
	}

	return v, nil
}
