package imports

import (
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

const (
	eofError        = "expected 'package', found 'EOF'"
	importPathError = "invalid import path:"
)

// Parse takes a Go file path and parses its import paths.
func Parse(file string) ([]string, error) {

	// we will parse a list of packages
	pkgs := []string{}

	// parse only imports
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
	if err != nil {

		// if the error is because of an invalid import path
		// we can handle it later
		if !strings.Contains(err.Error(), importPathError) {
			return nil, err
		}
	}

	// collect only the valid parsed import paths
	for _, i := range f.Imports {
		if Valid(i.Path.Value) { // check for import path validity
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				return nil, err
			}
			pkgs = append(pkgs, path)
		}
	}

	return pkgs, nil
}
