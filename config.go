package gbatis

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type configuration struct {
	XMLName xml.Name `xml:"configuration"`
	DBs     dbs      `xml:"dbs"`
}

type dbs struct {
	Default string `xml:"default,attr"`
	DBList  []db   `xml:"db"`
}

type db struct {
	ID         string     `xml:"id,attr"`
	Properties []property `xml:"property"`
}

type property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type dbInfo struct {
	Driver       string
	DataSource   string
	MaxOpenConns int
	MaxIdleConns int
}

func newDBInfo() *dbInfo {
	return &dbInfo{}
}

func (di dbInfo) check() (ok bool) {
	if di.Driver != "" && di.DataSource != "" && di.MaxOpenConns > 0 && di.MaxIdleConns > 0 && di.MaxOpenConns > di.MaxIdleConns {
		return true
	}
	return false
}

func (di *dbInfo) parse(properties []property) (err error) {
	if properties == nil {
		return fmt.Errorf("Input parameter invalid")
	}

	for _, property := range properties {

		switch property.Name {
		case "driver":
			di.Driver = property.Value
			break
		case "dataSource":
			di.DataSource = property.Value
			break
		case "maxOpenConns":
			di.MaxOpenConns, err = strconv.Atoi(property.Value)
			break
		case "maxIdleConns":
			di.MaxIdleConns, err = strconv.Atoi(property.Value)
			break
		}
	}
	return nil
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

	// log.Println(c)

	return &c, nil
}
