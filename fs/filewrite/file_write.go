package filewrite

import (
	"io/ioutil"
	"path/filepath"

	"github.com/bhoriuchi/opa-plugin-functions/utils"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "fs.file_write",
	Decl: types.NewFunction(types.Args(types.S, types.S, types.S), types.S),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a, b, c *ast.Term) (*ast.Term, error) {
	var file string
	var content string
	var fileModeStr string

	if err := ast.As(a.Value, &file); err != nil {
		return nil, err
	}
	if err := ast.As(b.Value, &content); err != nil {
		return nil, err
	}
	if err := ast.As(c.Value, &fileModeStr); err != nil {
		return nil, err
	}

	absFile, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	fileMode, err := utils.ParseFilemode(fileModeStr)
	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(absFile, []byte(content), fileMode); err != nil {
		return nil, err
	}

	return ast.StringTerm(absFile), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin3(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function3(Declaration, Implementation)
}
