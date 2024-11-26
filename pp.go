package pp

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/lipgloss"
	"github.com/shurcooL/go/reflectsource"
)

var keystyle = lipgloss.NewStyle().Underline(true)
var valueName = lipgloss.NewStyle().Foreground(lipgloss.Color(9)).Background(lipgloss.Color(0)).Italic(true)
var valuestyle = lipgloss.NewStyle()

var argToStringFunction reflect.Value
var argNameFormatterFunc reflect.Value

// Print handles nested structs and skips fields with tag `pp:"-"`.
func Print(v interface{}) {
	flattened := make(map[string]interface{})
	flattenStruct("", v, flattened)

	vName := reflectsource.GetParentArgExprAsString(uint32(0))
	// fmtV := fmt.Sprintf("%#v", v)
	returnMsg := ""
	// if vName == fmtV { // ? vName the same as value, just print one
	// 	returnMsg = ""
	// } else {
	// }
	returnMsg += fmt.Sprintf("[%s]", vName)

	if returnMsg != "" {
		fmt.Print(valueName.Render(returnMsg), " ")
	}

	if len(flattened) == 1 {
		for k, v := range flattened {
			if k == "" && returnMsg == "" {
				out := valuestyle.Render(fmt.Sprintf("%v", v))
				fmt.Println(out)
				return
			} else {
				fmt.Println(valuestyle.Render(fmt.Sprintf("%v", v)))
				return
			}
		}
	}

	// Format and print the flattened map
	result := ""
	for k, v := range flattened {
		result += fmt.Sprintf("%s %v | ", keystyle.Render(k), valuestyle.Render(fmt.Sprintf("%v", v)))
	}
	if len(result) != 0 {
		result = result[:len(result)-3]
	}
	fmt.Println(result)
}

// flattenStruct recursively flattens a struct or map into a single map with dot-separated keys.
func flattenStruct(prefix string, v interface{}, result map[string]interface{}) {
	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Ptr:
		if !val.IsNil() {
			flattenStruct(prefix, val.Elem().Interface(), result)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			if tag := field.Tag.Get("pp"); tag == "-" {
				continue
			}

			fieldName := field.Name
			if prefix != "" {
				fieldName = prefix + "." + fieldName
			}
			flattenStruct(fieldName, val.Field(i).Interface(), result)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			keyStr := fmt.Sprintf("%v", key)
			if prefix != "" {
				keyStr = prefix + "." + keyStr
			}
			flattenStruct(keyStr, val.MapIndex(key).Interface(), result)
		}
	default:
		result[prefix] = v
	}
}
