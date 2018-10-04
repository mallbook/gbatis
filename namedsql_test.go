package gbatis

import (
	"testing"
)

func TestParseNamedSQL(t *testing.T) {
	err := parseNamedSQL("examples/test.xml")
	if err == nil {
		t.Log("OK")
	} else {
		t.Log("error")
		t.Fail()
	}
}
