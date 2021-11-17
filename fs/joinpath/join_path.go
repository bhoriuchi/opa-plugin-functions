package joinpath

import (
	"path/filepath"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "fs.join_path",
	Decl: types.NewFunction(types.Args(types.NewArray(nil, types.S)), types.S),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
	var paths []string

	if err := ast.As(a.Value, &paths); err != nil {
		return nil, err
	}

	return ast.StringTerm(filepath.Join(paths...)), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin1(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function1(Declaration, Implementation)
}
