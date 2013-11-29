package conf

import (
	"testing"
)

func TestConf_Bool(t *testing.T) {
	c := Conf{map[string]interface{}{
		"foo": "bar",
		"baz": true,
	}}
	v := c.Bool("foo")
	if v {
		t.Errorf("the value of foo should be false, but it's %v", v)
	}

	v = c.Bool("baz")
	if !v {
		t.Errorf("the value of baz should be true, but it's %v", v)
	}
}

func TestConf_Int(t *testing.T) {
	c := Conf{map[string]interface{}{
		"foo": "bar",
		"baz": 1,
	}}
	v := c.Int("foo")
	if v != 0 {
		t.Errorf("the value of foo should be 0, but it's %v", v)
	}

	v = c.Int("baz")
	if v != 1 {
		t.Errorf("the value of baz should be 1, but it's %v", v)
	}
}

func TestConf_String(t *testing.T) {
	c := Conf{map[string]interface{}{
		"foo": "bar",
		"baz": 1,
	}}
	v := c.String("foo")
	if v != "bar" {
		t.Errorf("the value of foo should be bar, but it's %s", v)
	}

	v = c.String("baz")
	if v != "" {
		t.Errorf("the value of baz should be empty, but it's %s", v)
	}
}
