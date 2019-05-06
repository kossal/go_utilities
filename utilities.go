package utilities

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ParseTemplatesRecursively accepts a root path where it will
// recursively look for the extension name provided in the
// argument. Once it finds all ocurrences of the extension name
// it will call ParseFiles with the paths found. It can be used
// with template.Must(parseTemplatesRecursively).
func ParseTemplatesRecursively(path, ext string) (*template.Template, error) {
	var tmp *template.Template
	var pathFiles []string

	extRegex := regexp.MustCompile(`\.` + ext + `$`)

	// Recursively walk parent path and get list of files with positive termination
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Errorf("Error opening %v on path %v. Error: %g", info.Name(), path, err)
		}

		if extRegex.Match([]byte(info.Name())) && !info.IsDir() && err == nil {
			pathFiles = append(pathFiles, path)
		}

		return nil

	})
	if err != nil {
		return tmp, err
	}

	if len(pathFiles) == 0 {
		err := errors.New("Couldn´t find any file with termination:" + ext + "on path:" + path)
		return tmp, err
	}

	tmp, err = template.New("").ParseFiles(pathFiles...)

	return tmp, err
}
