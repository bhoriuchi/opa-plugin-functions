package mkdirall

import (
	"os"
	"path/filepath"

	"github.com/bhoriuchi/opa-plugin-functions/utils"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "fs.mkdirall",
	Decl: types.NewFunction(types.Args(types.S, types.S), types.S),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a, b *ast.Term) (*ast.Term, error) {
	var dir string
	var fileModeStr string

	if err := ast.As(a.Value, &dir); err != nil {
		return nil, err
	}
	if err := ast.As(b.Value, &fileModeStr); err != nil {
		return nil, err
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	fileMode, err := utils.ParseFilemode(fileModeStr)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(absDir, fileMode); err != nil {
		return nil, err
	}

	return ast.StringTerm(absDir), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin2(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function2(Declaration, Implementation)
}
