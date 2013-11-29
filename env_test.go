package conf

import (
	"os"
	"testing"
)

func TestEnv_Apply(t *testing.T) {
	os.Setenv("Foo", "Bar")
	c := &Conf{make(map[string]interface{})}
	e := env{}
	err := e.Apply(c)

	if err != nil {
		t.Fatalf("error should be nil, but it's %s", err)
	}

	if c.String("Foo") != "Bar" {
		t.Errorf("the value of Foo should be Bar, but it's %v", c.String("Foo"))
	}
}
