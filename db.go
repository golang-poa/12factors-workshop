package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type record struct {
	Kind  string      `json:"kind"`
	Value interface{} `json:"value"`
}

type product struct {
	Name string `json:"name"`
}

// DB is the application database file location
const DB = "/tmp/db.json"

var memoryDB []record

func init() {
	if err := loadFile(); err != nil {
		panic(err)
	}
}

func add(kind, val string) (*record, error) {

	row := record{
		Kind: kind,
		Value: product{
			Name: val,
		},
	}

	if err := checkIfRowExists(row); err != nil {
		return nil, err
	}

	memoryDB = append(memoryDB, row)

	if err := saveFile(); err != nil {
		return nil, err
	}

	return &row, nil
}

func find(kind string) ([]interface{}, error) {
	var out []interface{}

	for _, row := range memoryDB {
		if row.Kind == kind {
			out = append(out, row.Value)
		}
	}

	if len(out) == 0 {
		return nil, errors.New("not data found")
	}

	return out, nil
}

// database checks
func checkIfRowExists(newRow record) error {

	for _, row := range memoryDB {
		if row.Kind == newRow.Kind && row.Value == newRow.Value {
			return errors.New("Record already exists")
		}
	}

	return nil
}

// file manage functions
func saveFile() error {

	file, err := json.MarshalIndent(memoryDB, "", " ")

	if err != nil {
		return fmt.Errorf("could not unmarshal the memorydb: %v", err)
	}

	err = ioutil.WriteFile(conf.DB, file, 0644)

	if err != nil {
		return fmt.Errorf("could not unmarshal save the file: %v", err)
	}

	return nil
}

func loadFile() error {
	fmt.Println("carregando db file")
	file, err := ioutil.ReadFile(conf.DB)

	if err != nil {
		return fmt.Errorf("could not read the database file: %v", err)
	}

	err = json.Unmarshal(file, &memoryDB)

	if err != nil {
		return fmt.Errorf("could not unmarshal the database file contents: %v", err)
	}

	return nil
}
