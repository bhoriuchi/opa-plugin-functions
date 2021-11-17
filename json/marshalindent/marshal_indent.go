package marshalindent

import (
	"encoding/json"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "json.marshal_indent",
	Decl: types.NewFunction(types.Args(
		types.NewObject(nil, types.NewDynamicProperty(types.S, types.NewAny())),
		types.S,
		types.S,
	), types.S),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a, b, c *ast.Term) (*ast.Term, error) {
	var obj interface{}
	var prefix string
	var indent string

	if err := ast.As(a.Value, &obj); err != nil {
		return nil, err
	}
	if err := ast.As(b.Value, &prefix); err != nil {
		return nil, err
	}
	if err := ast.As(c.Value, &indent); err != nil {
		return nil, err
	}

	j, err := json.MarshalIndent(obj, prefix, indent)
	if err != nil {
		return nil, err
	}

	return ast.StringTerm(string(j)), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin3(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function3(Declaration, Implementation)
}
