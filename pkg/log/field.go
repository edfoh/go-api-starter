package log

type field struct {
	key   string
	value interface{}
}

func (f field) ToSlice() []interface{} {
	return []interface{}{f.key, f.value}
}

type fields []field

func (fs fields) ToSlice() []interface{} {
	var result []interface{}
	for _, f := range fs {
		result = append(result, f.ToSlice()...)
	}
	return result
}

func Field(key string, value interface{}) field {
	return field{key, value}
}
