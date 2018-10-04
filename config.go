package gbatis

import (
	"encoding/xml"
)

type configuration struct {
	XMLName xml.Name `xml:"configuration"`
}

type environments struct {
}

type environment struct {
	ID string `xml:"id,attr"`
}
