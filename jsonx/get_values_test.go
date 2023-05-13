package jsonx

import (
	"testing"

	"github.com/mgenware/goutil/test"
)

var sampleDict map[string]interface{}

func init() {
	dict, err := UnmarshalBytesToDict([]byte(`{"pn": 123, "nn": -1, "f": 3.1415, "s": "haha", "b": true, "dict": {"a": true, "b": ["element1", "element2"]}, "arr": [1, 2, "haha"]}`))
	if err != nil {
		panic(err)
	}
	sampleDict = dict
}

func TestGetStringOrDefault(t *testing.T) {
	test.Assert(t, GetStringOrDefault(sampleDict, "s"), "haha")
	test.Assert(t, GetStringOrDefault(sampleDict, "_"), "")
}

func TestGetStringOrNil(t *testing.T) {
	test.Assert(t, *GetStringOrNil(sampleDict, "s"), "haha")
	test.Assert(t, GetStringOrNil(sampleDict, "_"), (*string)(nil))
}

func TestGetBoolOrDefault(t *testing.T) {
	test.Assert(t, GetBoolOrDefault(sampleDict, "b"), true)
	test.Assert(t, GetBoolOrDefault(sampleDict, "_"), false)
}

func TestGetBoolOrNil(t *testing.T) {
	test.Assert(t, *GetBoolOrNil(sampleDict, "b"), true)
	test.Assert(t, GetBoolOrNil(sampleDict, "_"), (*bool)(nil))
}

func TestGetIntOrDefault(t *testing.T) {
	test.Assert(t, GetIntOrDefault(sampleDict, "nn"), -1)
	test.Assert(t, GetIntOrDefault(sampleDict, "_"), 0)
}

func TestGetIntOrNil(t *testing.T) {
	test.Assert(t, *GetIntOrNil(sampleDict, "nn"), -1)
	test.Assert(t, GetIntOrNil(sampleDict, "_"), (*int)(nil))
}

func TestGetUintOrDefault(t *testing.T) {
	test.Assert(t, GetUintOrDefault(sampleDict, "pn"), uint(123))
	test.Assert(t, GetUintOrDefault(sampleDict, "_"), uint(0))
}

func TestGetUintOrNil(t *testing.T) {
	test.Assert(t, *GetUintOrNil(sampleDict, "pn"), uint(123))
	test.Assert(t, GetUintOrNil(sampleDict, "_"), (*uint)(nil))
}

func TestGetInt64OrDefault(t *testing.T) {
	test.Assert(t, GetInt64OrDefault(sampleDict, "nn"), int64(-1))
	test.Assert(t, GetInt64OrDefault(sampleDict, "_"), int64(0))
}

func TestGetInt64OrNil(t *testing.T) {
	test.Assert(t, *GetInt64OrNil(sampleDict, "nn"), int64(-1))
	test.Assert(t, GetInt64OrNil(sampleDict, "_"), (*int64)(nil))
}

func TestGetUint64OrDefault(t *testing.T) {
	test.Assert(t, GetUint64OrDefault(sampleDict, "pn"), uint64(123))
	test.Assert(t, GetUint64OrDefault(sampleDict, "_"), uint64(0))
}

func TestGetUint64OrNil(t *testing.T) {
	test.Assert(t, *GetUint64OrNil(sampleDict, "pn"), uint64(123))
	test.Assert(t, GetUint64OrNil(sampleDict, "_"), (*uint64)(nil))
}

func TestGetDictOrNil(t *testing.T) {
	d1 := GetDictOrNil(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDictOrNil(sampleDict, "arr"), map[string]interface{}(nil))
}

func TestGetDictOrEmpty(t *testing.T) {
	d1 := GetDictOrEmpty(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDictOrEmpty(sampleDict, "arr"), make(map[string]interface{}))
}

func TestGetArrayOrNil(t *testing.T) {
	d1 := GetArrayOrNil(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArrayOrNil(sampleDict, "dict"), []interface{}(nil))
}

func TestGetArrayOrEmpty(t *testing.T) {
	d1 := GetArrayOrEmpty(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArrayOrEmpty(sampleDict, "dict"), make([]interface{}, 0))
}
