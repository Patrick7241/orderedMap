package tests

import (
	"fmt"
	"orderedMap"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	om := orderedMap.New()
	om.Set("key1", "value1")
	om.Set("key2", "value2")
	value, e := om.Get("key1")
	if !e {
		t.Error("key1 is not exist")
	}
	if value != "value1" {
		t.Error("key1's value is not value1")
	}
	fmt.Print("key1's value is: ", value)
	fmt.Println(om)
}
