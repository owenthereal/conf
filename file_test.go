package conf

import (
	"testing"
)

func TestFile_Apply(t *testing.T) {
	c := &Conf{make(map[string]interface{})}
	f := file{"./examples/config.json"}
	err := f.Apply(c)

	if err != nil {
		t.Errorf("error should be nil, but it's %s", err)
	}

	if c.String("DATABASE") != "postgres" {
		t.Errorf("the value of DATABASE should be postgres, but it's %v", c.String("DATABASE"))
	}

	f = file{"./not_exist.json"}
	err = f.Apply(c)

	if err == nil {
		t.Errorf("error should not be nil since file doesn't exit")
	}
}
