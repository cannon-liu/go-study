package myinject

import (
	"fmt"
	"reflect"
)

type Injector interface {
	Applicator
	Invoker
	TypeMapper

	SetParent(Injector)
}


type Applicator interface {
	Apply(interface{}) error
}

type Invoker interface {
	Invoke(interface{}) ([]reflect.Value, error)
}

type TypeMapper interface {
	Map(interface{}) TypeMapper
	MapTo(interface{}, interface{}) TypeMapper
	Set(reflect.Type, reflect.Value) TypeMapper
	Get(reflect.Type) reflect.Value
}

type injector struct {
	values map[reflect.Type]reflect.Value
	parent Injector
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}

	return t
}

func New() Injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}


func (inj *injector) Invoke(f interface{}) ([]reflect.Value,error) {

	t := reflect.TypeOf(f)

	var in = make([]reflect.Value,t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := inj.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}
		in[i] = val
	}
	return reflect.ValueOf(f).Call(in),nil
}


func (inj *injector) Apply(val interface{}) error {
	v := reflect.ValueOf(val)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()

	for i := 0 ; i < v.NumField(); i++ {
		f := v.Field(i)
		structField := t.Field(i)
		if f.CanSet() && (structField.Tag == "inject" || structField.Tag.Get("inject") != "") {
			ft := f.Type()
			v := inj.Get(ft)
			if !v.IsValid() {
				return fmt.Errorf("Value not found for type %v", ft)
			}
			f.Set(v)
		}
	}
	return nil
}


func (inj *injector) Map(val interface{}) TypeMapper {
	inj.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
	return inj
}

func (inj *injector) MapTo(val interface{},ifacePtr interface{}) TypeMapper {
	inj.values[InterfaceOf(ifacePtr)] = reflect.ValueOf(val)
	return inj
}


func (inj *injector) Set(typ reflect.Type, val reflect.Value) TypeMapper {
	inj.values[typ] = val
	return inj
}

func (inj *injector) Get(t reflect.Type) reflect.Value {
	val := inj.values[t]
	if val.IsValid() {
		return val
	}

	if t.Kind() == reflect.Interface {
		for k,v := range inj.values {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	if !val.IsValid() && inj.parent != nil {
		val = inj.parent.Get(t)
	}

	return val

}

func (inj *injector) SetParent(parent Injector) {
	inj.parent = parent
}
