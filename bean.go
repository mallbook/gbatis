package gbatis

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

type newBean func() interface{}

// RegisterBean register a bean
func RegisterBean(name string, f interface{}) {
	getBeanFactoryInstance().registerBean(name, f)
}

// NewBean new a bean object
func NewBean(name string) (reflect.Value, error) {
	return getBeanFactoryInstance().newBean(name)
}

var (
	beanFactoryInst *beanFactory
	beanFactoryOnce sync.Once
)

type beanFactory struct {
	sync.RWMutex
	newFuncs map[string]reflect.Value
}

func getBeanFactoryInstance() *beanFactory {
	beanFactoryOnce.Do(func() {
		beanFactoryInst = &beanFactory{
			newFuncs: make(map[string]reflect.Value, 0),
		}
	})
	return beanFactoryInst
}

func (f *beanFactory) newBean(name string) (reflect.Value, error) {
	f.RLock()
	defer f.RUnlock()
	v, ok := f.newFuncs[name]
	if !ok {
		return reflect.ValueOf(nil), fmt.Errorf("Not found bean named %s", name)
	}

	if v.Kind() != reflect.Func {
		return reflect.ValueOf(nil), fmt.Errorf("Need a func, but is %s, bean name=%s", v.Kind(), name)
	}

	var result reflect.Value
	results := v.Call(nil)
	if len(results) == 1 {
		result = results[0]
	}

	return result, nil
}

func (f *beanFactory) registerBean(name string, newFunc interface{}) {

	v := reflect.ValueOf(newFunc)
	if v.Kind() != reflect.Func {
		log.Printf("The second arg want to be func type, but is %s\n", v.Kind())
		return
	}

	f.Lock()
	defer f.Unlock()

	if _, ok := f.newFuncs[name]; !ok {
		f.newFuncs[name] = v
	} else {
		log.Printf("The bean with the name %s already exists", name)
	}
}
