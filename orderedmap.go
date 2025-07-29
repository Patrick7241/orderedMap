package orderedMap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

type OrderedMap struct {
	// store map's order
	keys []string
	// map's kv
	values map[string]interface{}
	// escape html or not
	escapeHTML bool
}

// New create a new ordered map
func New(escapeHTML ...bool) *OrderedMap {
	// default escape html is true
	e := true
	if len(escapeHTML) > 0 {
		e = escapeHTML[0]
	}
	return &OrderedMap{
		keys:       make([]string, 0),
		values:     make(map[string]interface{}),
		escapeHTML: e,
	}
}

func (om *OrderedMap) SetEscapeHTML(on bool) {
	om.escapeHTML = on
}

// Set a key-value pair
func (om *OrderedMap) Set(key string, value interface{}) {
	if _, ok := om.values[key]; !ok {
		// if key not exist, add key to ensure order
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

// Get a value by key
func (om *OrderedMap) Get(key string) (interface{}, bool) {
	return om.values[key], om.values[key] != nil
}

// Delete a key-value pair
func (om *OrderedMap) Delete(key string) {
	if _, ok := om.values[key]; !ok {
		return
	}
	// delete key
	for i, k := range om.keys {
		if k == key {
			om.keys = append(om.keys[:i], om.keys[i+1:]...)
			break
		}
	}
	// delete key-value pair
	delete(om.values, key)
}

// Clear the ordered map
func (om *OrderedMap) Clear() {
	om.keys = make([]string, 0)
	om.values = make(map[string]interface{})
}

// String returns a string representation of the ordered map
func (om *OrderedMap) String() string {
	var b strings.Builder
	b.WriteString("{")
	for i, key := range om.keys {
		if i > 0 {
			b.WriteString(", ")
		}
		val := om.values[key]
		if om.escapeHTML {
			if strVal, ok := val.(string); ok {
				val = html.EscapeString(strVal)
			}
		}
		b.WriteString(fmt.Sprintf("%q: %q", key, val))
	}
	b.WriteString("}")
	return b.String()
}

// MarshalJSON returns a JSON representation of the ordered map
func (om *OrderedMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(om.escapeHTML)
	for i, k := range om.keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		// add key
		if err := encoder.Encode(k); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		// add value
		if err := encoder.Encode(om.values[k]); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON parses a JSON-encoded string and stores the result in the ordered map
func (om *OrderedMap) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &om.values)
}
