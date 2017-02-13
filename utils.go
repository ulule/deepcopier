package deepcopier

import (
	"reflect"
	"strings"
)

// GetTagOptions parses deepcopier tag field and returns options.
func GetTagOptions(value string) map[string]string {
	options := map[string]string{}

	for _, opt := range strings.Split(value, ";") {
		o := strings.Split(opt, ":")

		// deepcopier:"keyword; without; value;"
		if len(o) == 1 {
			k := o[0]
			options[k] = ""
		}

		// deepcopier:"key:value; anotherkey:anothervalue"
		if len(o) == 2 {
			k, v := o[0], o[1]
			k = strings.TrimSpace(k)
			v = strings.TrimSpace(v)
			options[k] = v
		}
	}

	return options
}

// GetMethods returns type methods.
func GetMethods(instance interface{}) []string {
	t := reflect.TypeOf(instance)

	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	return methods
}

// InStringSlice checks if the given string is in the given slice of string.
func InStringSlice(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}

	return false
}
