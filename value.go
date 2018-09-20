package pikka

type Value struct {
	value interface{}
}

func (value *Value) Interface() interface{} {
	return value.value
}

func (value *Value) String() string {
	if value.value == nil {
		return ""
	}

	str, _ := value.value.(string)
	return str
}

func (value *Value) Integer() int {
	if value.value == nil {
		return 0
	}

	i, _ := value.value.(int)
	return i
}

func (value *Value) Float() float64 {
	if value.value == nil {
		return 0.
	}

	f, _ := value.value.(float64)
	return f
}

func (value *Value) Bool() bool {
	if value.value == nil {
		return false
	}

	b, _ := value.value.(bool)
	return b
}

func (value *Value) Map() map[string]interface{} {
	if value.value == nil {
		return nil
	}

	m, _ := value.value.(map[string]interface{})
	return m
}

func (value *Value) Slice() []interface{} {
	if value.value == nil {
		return nil
	}

	s, _ := value.value.([]interface{})
	return s
}

func (value *Value) StringSlice() []string {
	slice := value.Slice()
	if slice == nil {
		return nil
	}

	strings := make([]string, len(slice))
	for i, item := range slice {
		strings[i] = (&Value{item}).String()
	}
	return strings
}

func (value *Value) IntegerSlice() []int {
	slice := value.Slice()
	if slice == nil {
		return nil
	}

	integers := make([]int, len(slice))
	for i, item := range slice {
		integers[i] = (&Value{item}).Integer()
	}
	return integers
}

func (value *Value) FloatSlice() []float64 {
	slice := value.Slice()
	if slice == nil {
		return nil
	}

	floats := make([]float64, len(slice))
	for i, item := range slice {
		floats[i] = (&Value{item}).Float()
	}
	return floats
}

func (value *Value) BoolSlice() []bool {
	slice := value.Slice()
	if slice == nil {
		return nil
	}

	bools := make([]bool, len(slice))
	for i, item := range slice {
		bools[i] = (&Value{item}).Bool()
	}
	return bools
}

func (value *Value) MapSlice() []map[string]interface{} {
	slice := value.Slice()
	if slice == nil {
		return nil
	}

	maps := make([]map[string]interface{}, len(slice))
	for i, item := range slice {
		maps[i] = (&Value{item}).Map()
	}
	return maps
}
