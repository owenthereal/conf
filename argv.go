package conf

import (
	"strings"
)

type argv struct {
	args []string
}

func (a argv) Apply(conf *Conf) error {
	for i, arg := range a.args {
		argIsFlag := strings.HasPrefix(arg, "-")
		if !argIsFlag {
			continue
		}

		var nextArg string
		if i+1 < len(a.args) {
			nextArg = a.args[i+1]
		}

		argHasValue := strings.Contains(arg, "=")
		nextArgIsFlag := strings.HasPrefix(nextArg, "-")

		var value interface{}
		if argHasValue {
			split := strings.SplitN(arg, "=", 2)
			arg = split[0]
			value = trimQuotes(split[1])
		} else if !nextArgIsFlag && nextArg != "" {
			value = trimQuotes(nextArg)
		} else if nextArgIsFlag {
			value = true
		}

		if value != nil {
			arg = strings.TrimLeft(arg, "-")
			conf.Set(arg, value)
		}
	}

	return nil
}

// trim leading and trailing " and '
func trimQuotes(s string) string {
	s = strings.Trim(s, `"`)
	s = strings.Trim(s, `'`)

	return s
}
