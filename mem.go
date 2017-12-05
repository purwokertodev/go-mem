package mem

import (
	"errors"
	"reflect"
	"sync"
)

type DB struct {
	sync.RWMutex
	me map[interface{}]interface{}
}

func New() *DB {
	return &DB{me: make(map[interface{}]interface{})}
}

// TODO
func (d *DB) Set(id, val interface{}) error {
	typVal := reflect.TypeOf(val)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typVal.Kind() == reflect.Ptr {
		typVal = typVal.Elem()
	}

	typeID := reflect.TypeOf(id)

	if typeID.Kind() != reflect.String && typeID.Kind() != reflect.Int {
		return errors.New("can only use string and int for ID")
	}

	if typVal.Kind() == reflect.Struct {
		d.Lock()
		d.me[id] = val
		d.Unlock()
	}

	return nil
}

func (d *DB) Get(id interface{}) interface{} {
	var result interface{}
	d.RLock()
	res, ok := d.me[id]
	if ok {
		result = res
	}
	d.RUnlock()
	return result
}

func (d *DB) Del(id interface{}) error {
	d.RLock()
	_, ok := d.me[id]
	if !ok {
		return errors.New("Data not found")
	}
	delete(d.me, id)
	d.RUnlock()
	return nil
}
