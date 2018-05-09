package codegen

import (
	"sort"
	"strings"

	"github.com/vektah/gqlgen/neelance/schema"
	"golang.org/x/tools/go/loader"
)

func (cfg *Config) buildModels(types NamedTypes, prog *loader.Program) ([]Model, error) {
	var models []Model

	for _, typ := range cfg.schema.Types {
		var model Model
		switch typ := typ.(type) {
		case *schema.Object:
			obj, err := cfg.buildObject(types, typ)
			if err != nil {
				return nil, err
			}
			if obj.Root || obj.IsUserDefined {
				continue
			}
			model = cfg.obj2Model(obj)
		case *schema.InputObject:
			obj, err := buildInput(types, typ)
			if err != nil {
				return nil, err
			}
			if obj.IsUserDefined {
				continue
			}
			model = cfg.obj2Model(obj)
		case *schema.Interface, *schema.Union:
			intf := cfg.buildInterface(types, typ, prog)
			if intf.IsUserDefined {
				continue
			}
			model = int2Model(intf)
		default:
			continue
		}

		models = append(models, model)
	}

	sort.Slice(models, func(i, j int) bool {
		return strings.Compare(models[i].GQLType, models[j].GQLType) == -1
	})

	return models, nil
}

func (cfg *Config) obj2Model(obj *Object) Model {
	model := Model{
		NamedType: obj.NamedType,
		Fields:    []ModelField{},
	}

	model.GoType = ucFirst(obj.GQLType)
	model.Marshaler = &Ref{GoType: obj.GoType}

	for i := range obj.Fields {
		field := &obj.Fields[i]
		mf := ModelField{Type: field.Type,GQLName:field.GQLName}

		if mf.IsScalar {
			mf.GoVarName = ucFirst(field.GQLName)
			if mf.GoVarName == "Id" {
				mf.GoVarName = "ID"
			}
		} else if mf.IsInput {
			mf.GoVarName = ucFirst(field.GQLName)
		} else if mf.IsSlice() {
			// one to many, we don't need a prop, we need a resolver
		} else {
			mf.GoFKName = ucFirst(field.GQLName) + "ID"
			mf.GoFKType = "string"

			if obj, ok := cfg.schema.Types[field.GQLType].(*schema.Object); ok {
				for _, f := range obj.Fields {
					if strings.EqualFold(f.Name, "id") {
						if strings.Contains(f.Type.String(), "Int") {
							mf.GoFKType = "int"
						}
					}
				}
			}
		}

		model.Fields = append(model.Fields, mf)
	}

	return model
}

func int2Model(obj *Interface) Model {
	model := Model{
		NamedType: obj.NamedType,
		Fields:    []ModelField{},
	}

	model.GoType = ucFirst(obj.GQLType)
	model.Marshaler = &Ref{GoType: obj.GoType}

	return model
}
