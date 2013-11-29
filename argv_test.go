package conf

import (
	"testing"
)

func TestArgv_Apply(t *testing.T) {
	c := &Conf{make(map[string]interface{})}
	a := argv{[]string{"cmd", "--foo", "-f", "--bar", `"bar"`, `-baz="baz"`, "deploy"}}
	err := a.Apply(c)

	if err != nil {
		t.Fatalf("error should be nil, but it's %s", err)
	}

	if c.Size() != 4 {
		t.Errorf("the length of conf should be 3, but it's %v", c.Size())
	}

	v := c.Bool("foo")
	if !v {
		t.Errorf("the value of foo should be true, but it's %v", v)
	}

	v = c.Bool("f")
	if !v {
		t.Errorf("the value of f should be true, but it's %v", v)
	}

	vv := c.String("bar")
	if vv != "bar" {
		t.Errorf("the value of bar should be bar, but it's %v", vv)
	}

	vv = c.String("baz")
	if vv != "baz" {
		t.Errorf("the value of baz should be baz, but it's %v", vv)
	}
}
