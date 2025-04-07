package cacheutil

import (
	"testing"
	"reflect"
)

func TestCacheLoadStore(t *testing.T) {
	cache := &Cache[string,string]{}
	cache.Store("123","123")
	value := ""
	if _, ok := cache.Load("123") ; !ok {
		t.Errorf("values not exist %s",value)
	}
	value, _ = cache.Load("123")
	if value !=  "123" {
		t.Errorf("value not same %s",value)
	}
}
func TestCacheRange(t *testing.T){
	cache := &Cache[string,string]{}
	a := []string{"11","22","33"}
	cache.StoreAll(func(t string) string{
		return t
	},a);
	b := cache.RangeAll()
	if !reflect.DeepEqual(a,b) {
		t.Error("not same")
	}
}