package tests

import (
	"fmt"
	"github.com/Patrick7241/orderedMap"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	om := orderedMap.New()

	// test set get delete method
	om.Set("key1", "value1")
	value, e := om.Get("key1")
	if !e {
		t.Error("key1 is not exist")
	}
	if value != "value1" {
		t.Error("key1's value is not value1")
	}
	t.Logf("key1's value is %s\n", value)
	om.Delete("key1")
	value, e = om.Get("key1")
	if !e {
		t.Log("key1 is not exist after delete")
	} else {
		t.Error("key1 is exist after delete")
	}

	// test string representation
	for i := 0; i < 10; i++ {
		om.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	t.Log(om)
	om.Clear()
	t.Log(om)

	// test escape html
	dangerousValue := "<script>alert('xss')</script>"
	om.Set("dangerous", dangerousValue)
	om.SetEscapeHTML(false)
	t.Log(om)
	om.SetEscapeHTML(true)
	t.Log(om)
}
