package template

import (
	"bytes"
	"fmt"
	htmltemplate "html/template"
	"io"
	texttemplate "text/template"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

type Args struct {
	HTML     bool        `json:"html"`
	Template string      `json:"template"`
	Data     interface{} `json:"data"`
}

type Template interface {
	Execute(wr io.Writer, data interface{}) error
}

// function declaration
var Declaration = &rego.Function{
	Name: "template.exec",
	Decl: types.NewFunction(
		types.Args(
			types.NewObject(
				[]*types.StaticProperty{
					{
						Key:   "template",
						Value: types.NewString(),
					},
					{
						Key:   "data",
						Value: types.NewAny(),
					},
				},
				types.NewDynamicProperty(
					types.NewString(),
					types.NewAny(),
				),
			),
		),
		types.S,
	),
}

// function implementation
func Implementation(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
	var t Template
	var err error
	args := &Args{}
	buf := bytes.NewBuffer([]byte{})

	if err := ast.As(a.Value, args); err != nil {
		return nil, err
	}

	if args.Template == "" {
		return nil, fmt.Errorf("no template specified")
	}

	switch args.HTML {
	case true:
		t, err = htmltemplate.New("tmpl").Parse(args.Template)
	default:
		t, err = texttemplate.New("tmpl").Parse(args.Template)
	}

	if err != nil {
		return nil, fmt.Errorf("template parse failed: %s", err)
	}

	if err := t.Execute(buf, args.Data); err != nil {
		return nil, fmt.Errorf("template execute failed: %s", err)
	}

	out := buf.Bytes()
	return ast.StringTerm(string(out)), nil
}

// helper to register in runtime
func RegisterBuiltin() {
	rego.RegisterBuiltin1(Declaration, Implementation)
}

// helper to return rego option
func RegoFunc() func(r *rego.Rego) {
	return rego.Function1(Declaration, Implementation)
}
