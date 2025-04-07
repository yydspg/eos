package env

import (
	"os"
	"strconv"
	"testing"

)

func TestGetString(t *testing.T) {
	k := "HELLO"
	v := "WORLD"
	os.Setenv(k,v)

	value := GetString("HELLO", "DEFAULT")
	if value != v {
		t.Error("get value error")
	}
}


func TestGetInt(t *testing.T){
	k := "HELLO"
	v := 123
	os.Setenv(k, strconv.Itoa(v))
	value,err  := GetInt(k, 456)
	if err != nil {
		t.Error("get value error")
	}
	if value != v {
		t.Error("value not same")
	}
}