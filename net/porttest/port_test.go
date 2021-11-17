package porttest

import (
	"net"
	"time"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// function declaration
var Declaration = &rego.Function{
	Name: "net.port_test",
	Decl: types.NewFunction(types.Args(types.S, types.N, types.N), types.B),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a, b, c *ast.Term) (*ast.Term, error) {
	var (
		network string
		address string
		timeout string
	)

	if err := ast.As(a.Value, &network); err != nil {
		return nil, err
	}
	if err := ast.As(b.Value, &address); err != nil {
		return nil, err
	}
	if err := ast.As(c.Value, &timeout); err != nil {
		return nil, err
	}

	to, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTimeout(network, address, to)
	if err != nil {
		return ast.BooleanTerm(false), nil
	}

	defer conn.Close()
	return ast.BooleanTerm(true), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin3(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function3(Declaration, Implementation)
}
