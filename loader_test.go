package conf

import (
	"os"
	"testing"
)

func TestLoader_Load(t *testing.T) {
	l := NewLoader()
	c, err := l.Env().Defaults(map[string]interface{}{"foo": "bar"}).Load()

	if err != nil {
		t.Fatalf("err should be nil, but it's %s", err)
	}

	v := c.String("foo")
	if v != "bar" {
		t.Errorf("the value of foo should be bar, but it's %v", v)
	}

	os.Setenv("foo", "baz")
	c, err = l.Env().Defaults(map[string]interface{}{"foo": "bar"}).Load()

	if err != nil {
		t.Fatalf("err should be nil, but it's %s", err)
	}

	v = c.String("foo")
	if v != "baz" {
		t.Errorf("the value of foo should be baz, but it's %v", v)
	}
}

func TestLoader_AdapterList(t *testing.T) {
	s := adapterList{make([]Adapter, 0)}
	r := s.Add(file{"path"})

	if !r {
		t.Errorf("should be able to add a new adapter")
	}

	r = s.Add(file{"path"})
	if r {
		t.Errorf("shouldn't be able to add another adapter")
	}
}
