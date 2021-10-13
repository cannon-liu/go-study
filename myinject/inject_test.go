package myinject

import (
	"fmt"
	"reflect"
	"testing"
)

type SpecailString interface {

}

type TestStruct struct {

	Dep1 string `inject:"t" json:"-"`
	Dep2 SpecailString `inject`
	Dep3 string
}

type Greeter struct {
	Name string
}


func (g *Greeter) String() string {
	return "Hello,My Name is" + g.Name
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{},b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_InjectorInvoke(t *testing.T) {
	injector := New()
	expect(t,injector == nil,false)

	dep := "some dependecy"
	injector.Map(dep)
	dep2 := "another dep"
	injector.MapTo(dep2,(*SpecailString)(nil))
	dep3 := make(chan *SpecailString)
	dep4 := make(chan *SpecailString)
	typRecv := reflect.ChanOf(reflect.RecvDir,reflect.TypeOf(dep3).Elem())
	typSend := reflect.ChanOf(reflect.SendDir,reflect.TypeOf(dep4).Elem())
	injector.Set(typRecv,reflect.ValueOf(dep3))
	injector.Set(typSend,reflect.ValueOf(dep4))

	_,err := injector.Invoke(func(d1 string,d2 SpecailString,d3 <-chan *SpecailString, d4 chan<- *SpecailString) {
		expect(t,d1,dep)
		expect(t,d2,dep2)
		expect(t,reflect.TypeOf(d3).Elem(),reflect.TypeOf(dep3).Elem())
		expect(t,reflect.TypeOf(d4).Elem(),reflect.TypeOf(dep4).Elem())
		expect(t,reflect.TypeOf(d3).ChanDir(),reflect.RecvDir)
		expect(t,reflect.TypeOf(d4).ChanDir(),reflect.SendDir)
	})

	expect(t,err,nil)
}

func Test_InjectorInvokeReturnValues(t *testing.T) {
	injector := New()
	expect(t,injector == nil,false)

	dep := "some dependency"
	injector.Map(dep)
	dep2 := "another dep"
	injector.MapTo(dep2,(*SpecailString)(nil))

	result,err := injector.Invoke(func(d1 string,d2 SpecailString) string {
		expect(t,d1,dep)
		expect(t,d2,dep2)
		return "Hello world"
	})

	expect(t,result[0].String(),"Hello world")
	expect(t,err,nil)
}

func Test_InjectorApply(t *testing.T) {
	injector := New()

	injector.Map("a dep").MapTo("another dep", (*SpecailString)(nil))

	s := TestStruct{}
	err := injector.Apply(&s)

	expect(t,err,nil)

	expect(t,s.Dep1,"a dep")
	expect(t,s.Dep2,"another dep")
	expect(t,s.Dep3,"")
}


func Test_InterfaceOf(t *testing.T) {
	iType := InterfaceOf((*SpecailString)(nil))
	expect(t,iType.Kind(),reflect.Interface)

	iType = InterfaceOf((**SpecailString)(nil))
	expect(t,iType.Kind(),reflect.Interface)

	defer func() {
		rec := recover()
		refute(t,rec,nil)
	}()

	iType = InterfaceOf((*testing.T)(nil))
}

func Test_InjectorSet(t *testing.T) {
	injector := New()
	typ := reflect.TypeOf("string")

	typSend := reflect.ChanOf(reflect.SendDir,typ)
	typRec := reflect.ChanOf(reflect.RecvDir,typ)

	chanRecv := reflect.MakeChan(reflect.ChanOf(reflect.BothDir,typ),0)
	chanSend := reflect.MakeChan(reflect.ChanOf(reflect.BothDir,typ),0)

	injector.Set(typSend,chanSend)
	injector.Set(typRec,chanRecv)

	expect(t,injector.Get(typSend).IsValid(),true)
	expect(t,injector.Get(typRec).IsValid(),true)

	expect(t,injector.Get(chanSend.Type()).IsValid(),false)
}

func Test_InjectorGet(t *testing.T) {
	injector := New()

	injector.Map("some dependency")

	expect(t,injector.Get(reflect.TypeOf("string")).IsValid(),true)
	expect(t,injector.Get(reflect.TypeOf(11)).IsValid(),false)
}

func Test_InjectorSetParent(t *testing.T) {
	injector := New()
	injector.MapTo("another dep",(*SpecailString)(nil))

	injector2 := New()
	injector2.SetParent(injector)

	expect(t,injector2.Get(InterfaceOf((*SpecailString)(nil))).IsValid(),true)
}

func TestInjectImplementors(t *testing.T) {
	injector := New()
	g := &Greeter{"Tom"}
	injector.Map(g)

	expect(t,injector.Get(InterfaceOf((*fmt.Stringer)(nil))).IsValid(),true)
}

