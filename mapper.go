package gbatis

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type mapper struct {
	XMLName   xml.Name   `xml:"mapper"`
	Namespace string     `xml:"namespace,attr"`
	Select    []selecter `xml:"select"`
	Delete    []deleter  `xml:"delete"`
	Insert    []inserter `xml:"insert"`
	Update    []updater  `xml:"update"`
}

type selecter struct {
	ID         string `xml:"id,attr"`
	ResultType string `xml:"resultType,attr"`
	SQL        string `xml:",innerxml"`
}

type deleter struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

type inserter struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

type updater struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

func parseNamedSQL(namedSQL string) (*mapper, error) {
	file, err := os.Open(namedSQL)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	m := mapper{}
	err = xml.Unmarshal(data, &m)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	// fmt.Println(m)
	// fmt.Println(m.Namespace)
	// for index, value := range m.Select {
	// 	fmt.Printf("(%d)id=%s, sql=%s\n", index, value.ID, value.SQL)
	// }

	return &m, nil
}
