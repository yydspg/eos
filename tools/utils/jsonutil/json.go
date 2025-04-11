package jsonutil

import (
	"encoding/json"
	"github.com/eos/tools/errs"
)


type Json interface {
	// Interface returns the underlying data
	Interface() any
	// Encode returns its marshaled data as `[]byte`
	Encode() ([]byte, error)
	// EncodePretty returns its marshaled data as `[]byte` with indentation
	EncodePretty() ([]byte, error)
	// Implements the json.Marshaler interface.
	MarshalJSON() ([]byte, error)
	// Set modifies `Json` map by `key` and `value`
	// Useful for changing single key/value in a `Json` object easily.
	Set(key string, val any)
	// SetPath modifies `Json`, recursively checking/creating map keys for the supplied path,
	// and then finally writing in the value
	SetPath(branch []string, val any)
	// Del modifies `Json` map by deleting `key` if it is present.
	Del(key string)
	// Get returns a pointer to a new `Json` object
	// for `key` in its `map` representation
	//
	// useful for chaining operations (to traverse a nested JSON):
	//    js.Get("top_level").Get("dict").Get("value").Int()
	Get(key string) Json
	// GetPath searches for the item as specified by the branch
	// without the need to deep dive using Get()'s.
	//
	//   js.GetPath("top_level", "dict")
	GetPath(branch ...string) Json
	// CheckGet returns a pointer to a new `Json` object and
	// a `bool` identifying success or failure
	//
	// useful for chained operations when success is important:
	//    if data, ok := js.Get("top_level").CheckGet("inner"); ok {
	//        log.Println(data)
	//    }
	CheckGet(key string) (Json, bool)
	// Map type asserts to `map`
	Map() (map[string]any, error)
	// Array type asserts to an `array`
	Array() ([]any, error)
	// Bool type asserts to `bool`
	Bool() (bool, error)
	// String type asserts to `string`
	String() (string, error)
	// Bytes type asserts to `[]byte`
	Bytes() ([]byte, error)
	// StringArray type asserts to an `array` of `string`
	StringArray() ([]string, error)

	// Implements the json.Unmarshaler interface.
	UnmarshalJSON(p []byte) error
	// Float64 coerces into a float64
	Float64() (float64, error)
	// Int coerces into an int
	Int() (int, error)
	// Int64 coerces into an int64
	Int64() (int64, error)
	// Uint64 coerces into an uint64
	Uint64() (uint64, error)
}

func JsonMarshal(v any) ([]byte, error) {
	m, err := json.Marshal(v)
	return m, errs.Wrap(err)
}

func JsonUnmarshal(b []byte, v any) error {
	return errs.Wrap(json.Unmarshal(b, v))
}

func StructToJsonString(param any) string {
	dataType, _ := JsonMarshal(param)
	dataString := string(dataType)
	return dataString
}

// The incoming parameter must be a pointer
func JsonStringToStruct(s string, args any) error {
	err := json.Unmarshal([]byte(s), args)
	return err
}
