package fileread

import (
	"io/ioutil"
	"path/filepath"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "fs.file_read",
	Decl: types.NewFunction(types.Args(types.S), types.S),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
	var file string

	if err := ast.As(a.Value, &file); err != nil {
		return nil, err
	}

	absFile, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(absFile)
	if err != nil {
		return nil, err
	}

	return ast.StringTerm(string(content)), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin1(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function1(Declaration, Implementation)
}
