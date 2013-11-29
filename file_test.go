package conf

import (
	"testing"
)

func TestFile_Apply(t *testing.T) {
	c := &Conf{make(map[string]interface{})}
	f := file{"./test.json"}
	err := f.Apply(c)

	if err != nil {
		t.Errorf("error should be nil, but it's %s", err)
	}

	if c.String("Foo") != "Bar" {
		t.Errorf("the value of Foo should be Bar, but it's %v", c.String("Foo"))
	}

	f = file{"./not_exist.json"}
	err = f.Apply(c)

	if err == nil {
		t.Errorf("error should not be nil since file doesn't exit")
	}
}
