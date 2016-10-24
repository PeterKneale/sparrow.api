package apidsl

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
)

// Metadata is a set of key/value pairs that can be assigned to an object. Each value consists of a
// slice of strings so that multiple invocation of the Metadata function on the same target using
// the same key builds up the slice. Metadata may be set on attributes, media types, actions,
// responses, resources and API definitions.
//
// While keys can have any value the following names are handled explicitly by goagen when set on
// attributes.
//
// `struct:field:name`: overrides the Go struct field name generated by default by goagen.
// Applicable to attributes only.
//
//        Metadata("struct:field:name", "MyName")
//
// `struct:tag:xxx`: sets the struct field tag xxx on generated Go structs.  Overrides tags that
// goagen would otherwise set.  If the metadata value is a slice then the strings are joined with
// the space character as separator.
// Applicable to attributes only.
//
//        Metadata("struct:tag:json", "myName,omitempty")
//        Metadata("struct:tag:xml", "myName,attr")
//
// `swagger:tag:xxx`: sets the Swagger object field tag xxx.
// Applicable to resources and actions.
//
//        Metadata("swagger:tag:Backend")
//        Metadata("swagger:tag:Backend:desc", "Quick description of what 'Backend' is")
//        Metadata("swagger:tag:Backend:url", "http://example.com")
//        Metadata("swagger:tag:Backend:url:desc", "See more docs here")
//
// `swagger:summary`: sets the Swagger operation summary field.
// Applicable to actions.
//
//        Metadata("swagger:summary", "Short summary of what action does")
//
// The special key names listed above may be used as follows:
//
//        var Account = Type("Account", func() {
//                Attribute("service", String, "Name of service", func() {
//                        // Override default name to avoid clash with built-in 'Service' field.
//                        Metadata("struct:field:name", "ServiceName")
//                })
//        })
//
func Metadata(name string, value ...string) {
	switch def := dslengine.CurrentDefinition().(type) {
	case design.ContainerDefinition:
		att := def.Attribute()
		if att.Metadata == nil {
			att.Metadata = make(map[string][]string)
		}
		att.Metadata[name] = append(att.Metadata[name], value...)
	case *design.AttributeDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.MediaTypeDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ActionDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.FileServerDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ResourceDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ResponseDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.APIDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	default:
		dslengine.IncompatibleDSL()
	}
}