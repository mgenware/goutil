package jsonx

// GetString retrieves a string value for the given key in the dictionary, or a default value if it does not exist.
func GetString(dict map[string]interface{}, key string) string {
	val, _ := dict[key].(string)
	return val
}

// GetFloat64 retrieves a float64 value for the given key in the dictionary, or a default value if it does not exist.
func GetFloat64(dict map[string]interface{}, key string) float64 {
	val, _ := dict[key].(float64)
	return val
}

// GetInt retrieves an int value for the given key in the dictionary, or a default value if it does not exist.
func GetInt(dict map[string]interface{}, key string) int {
	// All numbers parsed from JSON are float64
	val := GetFloat64(dict, key)
	return int(val)
}

// GetUint retrieves a uint value for the given key in the dictionary, or a default value if it does not exist.
func GetUint(dict map[string]interface{}, key string) uint {
	// All numbers parsed from JSON are float64
	val := GetFloat64(dict, key)
	return uint(val)
}

// GetInt64 retrieves an int64 value for the given key in the dictionary, or a default value if it does not exist.
func GetInt64(dict map[string]interface{}, key string) int64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64(dict, key)
	return int64(val)
}

// GetUint64 retrieves a uint64 value for the given key in the dictionary, or a default value if it does not exist.
func GetUint64(dict map[string]interface{}, key string) uint64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64(dict, key)
	return uint64(val)
}

// GetBool retrieves a bool value for the given key in the dictionary, or a default value if it does not exist.
func GetBool(dict map[string]interface{}, key string) bool {
	val, _ := dict[key].(bool)
	return val
}

// GetDict retrieves a dictionary value for the given key in the dictionary, or nil if it does not exist.
func GetDict(dict map[string]interface{}, key string) map[string]interface{} {
	val, _ := dict[key].(map[string]interface{})
	return val
}

// GetDictOrEmpty retrieves a dictionary value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetDictOrEmpty(dict map[string]interface{}, key string) map[string]interface{} {
	val := GetDict(dict, key)
	if val == nil {
		return make(map[string]interface{})
	}
	return val
}

// GetArray retrieves an array value for the given key in the dictionary, or nil if it does not exist.
func GetArray(dict map[string]interface{}, key string) []interface{} {
	val, _ := dict[key].([]interface{})
	return val
}

// GetArrayOrEmpty retrieves the value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetArrayOrEmpty(dict map[string]interface{}, key string) []interface{} {
	val := GetArray(dict, key)
	if val == nil {
		return make([]interface{}, 0)
	}
	return val
}
