package bean

import (
	// "encoding/json"
	"log"
	"testing"
)

func TestMall(t *testing.T) {
	m := NewMall()
	// fmt.Println(t.(type))
	//v := reflect.Value(m)
	//fmt.Println(v)
	m.ID = "1"
	m.Name = "KKMall"
	m.Avatar = "1.gif"
	m.Story = "HI KKMall"

	log.Print(m)
	// json.Marshal(m)
	// var s string
	/*if s, err := json.Marshal(m); err == nil {
		log.Println(string(s))
	} else {
		log.Println(err)
	}*/
	// log.Println(s)
}
