package grammar

import (
	"encoding/json"
)

type Function struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}
type Functions []Function

func (f Functions) ToJSONStructure() JSONStructure {
	js := JSONStructure{}
	for _, function := range f {
		//	t := function.Parameters["type"]
		//tt := t.(string)

		properties := function.Parameters["properties"]
		dat, _ := json.Marshal(properties)
		prop := map[string]interface{}{}
		json.Unmarshal(dat, &prop)
		js.OneOf = append(js.OneOf, Item{
			Type: "object",
			Properties: Properties{
				Function: FunctionName{Const: function.Name},
				Arguments: Argument{
					Type:       "object",
					Properties: prop,
				},
			},
		})
	}
	return js
}

// Select returns a list of functions containing the function with the given name
func (f Functions) Select(name string) Functions {
	var funcs Functions

	for _, f := range f {
		if f.Name == name {
			funcs = []Function{f}
			break
		}
	}

	return funcs
}
