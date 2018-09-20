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
			path:          "key",
			expectedValue: "foo",
			data: map[string]interface{}{
				"key": "foo",
			},
		},
		{
			path: "",
			expectedValue: map[string]interface{}{
				"key": "foo",
			},
			data: map[string]interface{}{
				"key": "foo",
			},
		},
		{
			path:          "key2",
			expectedValue: nil,
			data: map[string]interface{}{
				"key1": "foo",
			},
		},
		{
			path:          "key.foo",
			expectedValue: nil,
			data: map[string]interface{}{
				"key": "foo",
			},
		},
		{
			path:          "nested.key",
			expectedValue: "foo",
			data: map[string]interface{}{
				"nested": map[string]interface{}{
					"key": "foo",
				},
			},
		},
		{
			path:          "array.0.key",
			expectedValue: "foo",
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"key": "foo"},
				},
			},
		},
		{
			path:          "array.#.key",
			expectedValue: []interface{}{"foo", "bar"},
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"key": "foo"},
					{"key": "bar"},
				},
			},
		},
		{
			path:          "array.#.key1.#.key2",
			expectedValue: []interface{}{[]interface{}{"hello"}, []interface{}{"world"}},
			data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"key1": []map[string]interface{}{{"key2": "hello"}}},
					{"key1": []map[string]interface{}{{"key2": "world"}}},
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
