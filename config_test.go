package gbatis

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig("etc/conf/gbatis.xml")
	if err != nil {
		t.Fail()
	} else {
		t.Log(c)
		t.Log("ok")
	}
}
