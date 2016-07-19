package structure

import (
	"strings"

	"github.com/flourish-ship/skeleton/errors"
)

type mapper interface {
	Put(key string, val interface{}) (bool, error)
	Pop(key string) (interface{}, error)
	GetValue(key string) (interface{}, error)
	ContainsKey(key string) bool
	Delete(key string) (bool, error)
	Sort(order string)
}

type Map struct {
	mapper map[string]interface{}
}

func (mapper *Map) Put(key string, val interface{}) (bool, error) {

	if mapper.mapper == nil {
		return false, errors.NewErrors("map can't be empty!")
	}

	if len(strings.TrimSpace(key)) {
		return false, errors.NewErrors("map key can't be empty!")
	}

	if val == nil {
		return false, errors.NewErrors("map value can't be empty!")
	}

	mapper.mapper[key] = val
	return true, nil
}

// popup the given key and delete the key
func (mapper *Map) Pop(key string) (interface{}, error) {

	var val interface{}
	if mapper.mapper == nil {
		return nil, errors.NewErrors("map can't be empty!")
	}

	if len(strings.TrimSpace(key)) {
		return nil, errors.NewErrors("map key can't be empty!")
	}

	if ok := mapper.mapper[key]; !ok {
		return nil, errors.NewErrors("there is not exist the key which given!")
	}

	val = mapper.mapper[key]
	delete(mapper.mapper, key)

	return val, nil
}

// get the value form the map which key is the given key
func (mapper *Map) GetValue(key string) (interface{}, error) {

	if mapper.mapper == nil {
		return nil, errors.NewErrors("map can't be empty!")
	}

	if len(strings.TrimSpace(key)) {
		return nil, errors.NewErrors("map key can't be empty!")
	}

	if _, ok := mapper.mapper[key]; !ok {
		return nil, errors.NewErrors("there is not exist the key which given!")
	}

	return mapper.mapper[key], nil
}

func (mapper *Map) ContainsKey(key string) (bool, error) {
	if mapper.mapper == nil {
		return nil, errors.NewErrors("map can't be empty!")
	}

	if len(strings.TrimSpace(key)) {
		return nil, errors.NewErrors("map key can't be empty!")
	}

	_, ok := mapper.mapper[key]
	return ok, nil
}

// delete key from map
func (mapper *Map) Delete(key string) (bool, error) {

	if mapper.mapper == nil {
		return false, errors.NewErrors("map can't be empty!")
	}

	if len(strings.TrimSpace(key)) {
		return false, errors.NewErrors("map key can't be empty!")
	}

	delete(mapper.mapper, key)

	return true, nil
}