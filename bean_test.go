package gbatis

import (
	"log"
	"testing"
	//"github.com/mallbook/gbatis/bean"
)

type student struct {
	id   string
	name string
	age  int
}

func newStudent() *student {
	return &student{}
}

// func init() {
// 	RegisterBean("student", newStudent)
// }

func TestNewBean(t *testing.T) {
	f := getBeanFactoryInstance()
	v, err := f.newBean("student")
	if err != nil {
		t.Log(err)
		return
	}

	log.Println(v.Type())
	log.Println(v)
}
