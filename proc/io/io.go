package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lustigo/refingerprint/proc/data"
)

// ParseData reads the given Filename and parses the content to the Data Structure
func ParseData(path string) (*data.Data, error) {
	c, err := readFile(path)
	if err != nil {
		return nil, fmt.Errorf("Could not parse Data: %v", err)
	}

	d := &data.Data{}
	err = json.Unmarshal(c, &d)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal: %v", err)
	}

	return d, nil
}

// Reads the given Path
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("Could not read file: %v", err)
	}

	f.Close()
	return content, nil
}
