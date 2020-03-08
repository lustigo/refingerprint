package cmd

import (
	"errors"
	"os"
	"path/filepath"
)

// ALLOWEDEXTENSIONS defines the AllowedExtensions for the files
var ALLOWEDEXTENSIONS = []string{".txt", ".json"}

// GetFiles looks for:
// - all the files, if a directory is given
// - all the files, if multiple files are given
// - the file, if a single file is given
// returns an error if there is no valid file
// returns all the correct given files
func GetFiles(args []string) ([]string, error) {
	paths := make([]string, 0)

	for _, arg := range args {
		paths = checkArg(arg, paths)
	}

	if len(paths) == 0 {
		return nil, errors.New("Please provide a valid file as an argument or check if you are allowed to access that file")
	}
	return paths, nil
}

// checkArg checks a single argument if it is a file or a directory and if it exists
func checkArg(arg string, paths []string) []string {
	s, err := os.Stat(arg)

	if err != nil {
		return paths
	}

	if s.IsDir() {
		f, _ := os.Open(s.Name())
		contents, _ := f.Readdirnames(0)
		for _, name := range contents {
			paths = checkArg(filepath.Join(s.Name(), name), paths)
		}
		f.Close()
	} else {
		paths = checkFileAndAppend(arg, paths)
	}

	return paths
}

// checkFileAndAppend checks the given file for the correct extension and appends it to the given slice
func checkFileAndAppend(file string, paths []string) []string {
	if Contains(ALLOWEDEXTENSIONS, filepath.Ext(file)) {
		paths = append(paths, file)
	}
	return paths
}

// Contains checks if the given needle is in the haystack
func Contains(haystack []string, needle string) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}
