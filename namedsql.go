package gbatis

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Mapper struct {
	XMLName   xml.Name   `xml:"mapper"`
	Namespace string     `xml:"namespace,attr"`
	Select    []Selecter `xml:"select"`
	Delete    []Deleter  `xml:"delete"`
	Insert    []Inserter `xml:"insert"`
	Update    []Updater  `xml:"update"`
}

type Selecter struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

type Deleter struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

type Inserter struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

type Updater struct {
	ID  string `xml:"id,attr"`
	SQL string `xml:",innerxml"`
}

func parseNamedSQL(namedSQL string) error {
	file, err := os.Open(namedSQL)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	d := Mapper{}
	err = xml.Unmarshal(data, &d)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	fmt.Println(d)
	fmt.Println(d.Namespace)
	for index, value := range d.Select {
		fmt.Printf("(%d)id=%s, sql=%s\n", index, value.ID, value.SQL)
	}

	return nil
}
