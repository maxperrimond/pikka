package pikka

import (
	"reflect"
	"strconv"
	"strings"
)

func getNextField(fields []string) (string, []string) {
	if len(fields) > 1 {
		return fields[0], fields[1:]
	} else if len(fields) == 1 {
		return fields[0], []string{}
	}

	return "", []string{}
}

func getFromField(field string, subFields []string, v interface{}) interface{} {
	if field == "" {
		return v
	}

	nextField, newSubFields := getNextField(subFields)

	valueOf := reflect.ValueOf(v)
	switch valueOf.Kind() {
	case reflect.Map:
		return getFromField(nextField, newSubFields, v.(map[string]interface{})[field])
	case reflect.Slice, reflect.Array:
		if field == "#" {
			values := make([]interface{}, valueOf.Len())

			for i := 0; i < valueOf.Len(); i++ {
				values[i] = getFromField(strconv.Itoa(i), subFields, v)
			}

			return values
		}

		i, err := strconv.Atoi(field)
		if err != nil {
			return nil
		}

		return getFromField(nextField, newSubFields, valueOf.Index(i).Interface())
	default:
		return nil
	}
}

func GetFromPath(path string, m map[string]interface{}) *Value {
	fields := strings.Split(path, ".")
	if len(fields) == 0 {
		return &Value{m}
	}

	v := getFromField(fields[0], fields[1:], m)

	return &Value{v}
}
