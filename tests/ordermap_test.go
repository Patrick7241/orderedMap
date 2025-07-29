package tests

import (
	"encoding/json"
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
	om.Clear()

	// test ser and deser
	om.Set("你好", "va")
	om.Set("shello", "value2")
	om.Set("大", "vale3")
	om.Set("卡达克", "value4")

	for i := 0; i < 10; i++ {
		jsonData, _ := json.Marshal(om)
		t.Logf("第%d次： %v", i, string(jsonData))
	}

	//test := make(map[string]interface{})
	//err := json.Unmarshal(jsonData, &test)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("third: %s", test)

}
