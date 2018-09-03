package gofactory

import "testing"

func Test_Main(t *testing.T) {
	Default.Register("abc", "123")
	var v string
	err := Default.Get("abc", &v)
	if nil != err {
		t.Fatal(err)
	}
	if v != "123" {
		t.Fatal("Got v expect return \"123\"")
	}

	var iv int
	err = Default.Get("abc", &iv)
	if nil == err {
		t.Fatal("Type not match, expect got error")
	}
}
