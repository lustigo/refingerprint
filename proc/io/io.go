// Package io contains files for the interaction with the filesystem
package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/lustigo/refingerprint/proc/data"
	"github.com/sbinet/go-arff"
)

// ParseData reads the given Filename and parses the content to the Data Structure
func ParseData(path string) (*data.Data, error) {
	c, err := ReadFile(path)
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

// ReadFile reads the given Path
func ReadFile(path string) ([]byte, error) {
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

// GetFileName returns the FileName without extension
func GetFileName(path string) string {
	base := filepath.Base(path)
	splitted := strings.Split(base, ".")
	if len(splitted) == 0 {
		return ""
	}
	return splitted[0]
}

// WriteCSVFile writes the given Features to the given path as a CSV File
func WriteCSVFile(path string, data []*data.ProcessedFeatures) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not open file %v: %v\n", file, err)
		return
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		fmt.Printf("Could not write to file %v: %v\n", file, err)
	}
}

// WriteARFFFile writes the given Features to the given path as a ARFF File
func WriteARFFFile(path string, d []*data.ProcessedFeatures) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not open file %v: %v\n", file, err)
		return
	}

	enc, err := arff.NewEncoder(file)
	if err != nil {
		fmt.Printf("Could not get an ARFF Encoder: %v\n", err)
		file.Close()
		return
	}

	enc.Header = data.GetARFFHeader()
	for _, feature := range d {
		if err = enc.Encode(feature); err != nil {
			fmt.Printf("Could not encode %v: %v\n", feature, err)
		}
	}

	file.Sync()
	file.Close()
}
