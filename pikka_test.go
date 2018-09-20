package pikka

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFromPath(t *testing.T) {
	var dataSet = []struct {
		path          string
		expectedValue interface{}
		data          map[string]interface{}
	}{
		{
			path:          "foo",
			expectedValue: "foo",
			data: map[string]interface{}{
				"foo": "foo",
			},
		},
		{
			path: "",
			expectedValue: map[string]interface{}{
				"foo": "foo",
			},
			data: map[string]interface{}{
				"foo": "foo",
			},
		},
		{
			path:          "bar",
			expectedValue: nil,
			data: map[string]interface{}{
				"foo": "foo",
			},
		},
		{
			path:          "foo.foo",
			expectedValue: nil,
			data: map[string]interface{}{
				"foo": "foo",
			},
		},
		{
			path:          "nested.foo",
			expectedValue: "foo",
			data: map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": "foo",
				},
			},
		},
		{
			path:          "array.0.foo",
			expectedValue: "foo",
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"foo": "foo"},
				},
			},
		},
		{
			path:          "array.#.foo",
			expectedValue: []interface{}{"foo", "bar"},
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"foo": "foo"},
					{"foo": "bar"},
				},
			},
		},
		{
			path:          "array.#.foo.#.bar",
			expectedValue: []interface{}{[]interface{}{"hello"}, []interface{}{"world"}},
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"foo": []map[string]interface{}{{"bar": "hello"}}},
					{"foo": []map[string]interface{}{{"bar": "world"}}},
				},
			},
		},
	}

	for _, data := range dataSet {
		t.Run(fmt.Sprintf("WithPath_'%s'", data.path), func(t *testing.T) {
			result := GetFromPath(data.path, data.data)
			assert.Equal(t, data.expectedValue, result.Interface())
		})
	}
}
