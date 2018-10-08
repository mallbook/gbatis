package gbatis

import (
	"encoding/xml"
	"os"
	"log"
	"io/ioutil"
)

type configuration struct {
	XMLName xml.Name     `xml:"configuration"`
	Envs    environments `xml:"environments"`
}

type environments struct {
	Default string        `xml:"default,attr"`
	Envs    []environment `xml:"environment"`
}

type environment struct {
	ID         string     `xml:"id,attr"`
	DataSource dataSource `xml:"dataSource"`
}

type dataSource struct {
	Type       string     `xml:"type,attr"`
	Properties []property `xml:"property"`
}

type property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func loadConfig(configFile string) (*configuration, error) {

	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	c := configuration{}
	err = xml.Unmarshal(data, &c)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	log.Println(c)

	return &c, nil
}
