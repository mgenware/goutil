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

func TestGetString(t *testing.T) {
	test.Assert(t, GetString(sampleDict, "s"), "haha")
	test.Assert(t, GetString(sampleDict, "_"), "")
}

func TestGetBool(t *testing.T) {
	test.Assert(t, GetBool(sampleDict, "b"), true)
	test.Assert(t, GetBool(sampleDict, "_"), false)
}

func TestGetInt(t *testing.T) {
	test.Assert(t, GetInt(sampleDict, "nn"), -1)
	test.Assert(t, GetInt(sampleDict, "_"), 0)
}

func TestGetUint(t *testing.T) {
	test.Assert(t, GetUint(sampleDict, "pn"), uint(123))
	test.Assert(t, GetUint(sampleDict, "_"), uint(0))
}

func TestGetInt64(t *testing.T) {
	test.Assert(t, GetInt64(sampleDict, "nn"), int64(-1))
	test.Assert(t, GetInt64(sampleDict, "_"), int64(0))
}

func TestGetUint64(t *testing.T) {
	test.Assert(t, GetUint64(sampleDict, "pn"), uint64(123))
	test.Assert(t, GetUint64(sampleDict, "_"), uint64(0))
}

func TestGetDict(t *testing.T) {
	d1 := GetDict(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDict(sampleDict, "arr"), map[string]interface{}(nil))
}

func TestGetDictOrEmpty(t *testing.T) {
	d1 := GetDictOrEmpty(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDictOrEmpty(sampleDict, "arr"), make(map[string]interface{}))
}

func TestGetArray(t *testing.T) {
	d1 := GetArray(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArray(sampleDict, "dict"), []interface{}(nil))
}

func TestGetArrayOrEmpty(t *testing.T) {
	d1 := GetArrayOrEmpty(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArrayOrEmpty(sampleDict, "dict"), make([]interface{}, 0))
}
