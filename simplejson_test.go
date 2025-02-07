package simplejson

import (
	"testing"
)

func TestBase(t *testing.T) {
	js, err := NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"string_array_null": ["abc", null, "efg"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true,
			"sub_obj": {"a": 1}
		}
	}`))

	if js == nil {
		t.Fatal("got nil")
	}

	if err != nil {
		t.Fatalf("got err: %#v", err)
	}

	if s, _ := js.Get("test").Get("string").String(); s != "simplejson" {
		t.Errorf("got %#v", s)
	}

	if b, _ := js.Get("test").Get("bool").Bool(); b != true {
		t.Errorf("got %#v", b)
	}
}

func TestSet(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	err = js.Set("baz", "bing")
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	s, err := js.GetPath("baz").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "bing" {
		t.Errorf("got %#v", s)
	}
}

func TestSetPath(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	err = js.SetPath([]string{"foo", "bar"}, "baz")
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	s, err := js.GetPath("foo", "bar").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "baz" {
		t.Errorf("got %#v", s)
	}
}
