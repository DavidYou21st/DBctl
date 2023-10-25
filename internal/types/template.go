package types

// Types defines a template for types in model.
const Types = `
type {{.upperStartCamelObject}} struct {
		{{.fields}}
}
`

// Field defines a filed template for types
const Field = `{{.name}} {{.type}}`
