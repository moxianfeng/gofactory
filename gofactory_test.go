package gofactory

import "testing"

type T interface {
	Do(what string) string
}

type S struct {
}

func (s *S) Do(what string) string {
	return what
}

func Test_Main(t *testing.T) {
	Default.Register("abc", "123")
	var v string
	err := Default.GetObject("abc", &v)
	if nil != err {
		t.Fatal(err)
	}
	if v != "123" {
		t.Fatal("Got v expect return \"123\"")
	}

	var iv int
	err = Default.GetObject("abc", &iv)
	if nil == err {
		t.Fatal("Type not match, expect got error")
	}
}

func Test_Interface(t *testing.T) {
	Default.Register("interface", &S{})

	v, e := Default.GetInterface("abc", new(T))
	if nil == e {
		t.Fatal("Type not match, expect got error")
	}

	v, e = Default.GetInterface("interface", (*T)(nil))
	if nil != e {
		t.Fatal(e)
	}
	ret := v.(T).Do("abcd")
	t.Log(ret)
}
