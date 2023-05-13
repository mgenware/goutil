package jsonx

// GetStringOrDefault retrieves a string value for the given key in the dictionary, or a default value if it does not exist.
func GetStringOrDefault(dict map[string]interface{}, key string) string {
	val, _ := dict[key].(string)
	return val
}

// GetStringOrNil retrieves a string value for the given key in the dictionary, or nil if it does not exist.
func GetStringOrNil(dict map[string]interface{}, key string) *string {
	val, ok := dict[key].(string)
	if ok {
		return &val
	}
	return nil
}

// GetFloat64OrDefault retrieves a float64 value for the given key in the dictionary, or a default value if it does not exist.
func GetFloat64OrDefault(dict map[string]interface{}, key string) float64 {
	val, _ := dict[key].(float64)
	return val
}

// GetFloat64OrNil retrieves a float64 value for the given key in the dictionary, or nil if it does not exist.
func GetFloat64OrNil(dict map[string]interface{}, key string) *float64 {
	val, ok := dict[key].(float64)
	if ok {
		return &val
	}
	return nil
}

// GetIntOrDefault retrieves an int value for the given key in the dictionary, or a default value if it does not exist.
func GetIntOrDefault(dict map[string]interface{}, key string) int {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return int(val)
}

// GetIntOrNil retrieves an int value for the given key in the dictionary, or nil if it does not exist.
func GetIntOrNil(dict map[string]interface{}, key string) *int {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrNil(dict, key)
	if val == nil {
		return nil
	}
	intVal := int(*val)
	return &intVal
}

// GetUintOrDefault retrieves a uint value for the given key in the dictionary, or a default value if it does not exist.
func GetUintOrDefault(dict map[string]interface{}, key string) uint {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return uint(val)
}

// GetUintOrNil retrieves a uint value for the given key in the dictionary, or nil if it does not exist.
func GetUintOrNil(dict map[string]interface{}, key string) *uint {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrNil(dict, key)
	if val == nil {
		return nil
	}
	uintVal := uint(*val)
	return &uintVal
}

// GetInt64OrDefault retrieves an int64 value for the given key in the dictionary, or a default value if it does not exist.
func GetInt64OrDefault(dict map[string]interface{}, key string) int64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return int64(val)
}

// GetInt64OrNil retrieves an int64 value for the given key in the dictionary, or a default value if it does not exist.
func GetInt64OrNil(dict map[string]interface{}, key string) *int64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrNil(dict, key)
	if val == nil {
		return nil
	}
	int64Val := int64(*val)
	return &int64Val
}

// GetUint64OrDefault retrieves a uint64 value for the given key in the dictionary, or a default value if it does not exist.
func GetUint64OrDefault(dict map[string]interface{}, key string) uint64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return uint64(val)
}

// GetUint64 retrieves a uint64 value for the given key in the dictionary, or nil if it does not exist.
func GetUint64OrNil(dict map[string]interface{}, key string) *uint64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrNil(dict, key)
	if val == nil {
		return nil
	}
	uint64Val := uint64(*val)
	return &uint64Val
}

// GetBoolOrDefault retrieves a bool value for the given key in the dictionary, or a default value if it does not exist.
func GetBoolOrDefault(dict map[string]interface{}, key string) bool {
	val, _ := dict[key].(bool)
	return val
}

// GetBoolOrNil retrieves a bool value for the given key in the dictionary, or a default value if it does not exist.
func GetBoolOrNil(dict map[string]interface{}, key string) *bool {
	val, ok := dict[key].(bool)
	if ok {
		return &val
	}
	return nil
}

// GetDictOrNil retrieves a dictionary value for the given key in the dictionary, or nil if it does not exist.
func GetDictOrNil(dict map[string]interface{}, key string) map[string]interface{} {
	val, _ := dict[key].(map[string]interface{})
	return val
}

// GetDictOrEmpty retrieves a dictionary value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetDictOrEmpty(dict map[string]interface{}, key string) map[string]interface{} {
	val := GetDictOrNil(dict, key)
	if val == nil {
		return make(map[string]interface{})
	}
	return val
}

// GetArray retrieves an array value for the given key in the dictionary, or nil if it does not exist.
func GetArrayOrNil(dict map[string]interface{}, key string) []interface{} {
	val, _ := dict[key].([]interface{})
	return val
}

// GetArrayOrEmpty retrieves the value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetArrayOrEmpty(dict map[string]interface{}, key string) []interface{} {
	val := GetArrayOrNil(dict, key)
	if val == nil {
		return make([]interface{}, 0)
	}
	return val
}
