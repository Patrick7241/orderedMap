package orderedMap

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
