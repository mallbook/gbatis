package gbatis

import (
	"testing"
)

func TestParseNamedSQL(t *testing.T) {
	mapper, err := parseNamedSQL("examples/test.xml")
	if err == nil {
		t.Log(mapper)
		t.Log("OK")
	} else {
		t.Fail()
	}
}
