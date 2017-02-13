package deepcopier

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// TagName is the deepcopier struct tag name.
	TagName = "deepcopier"
	// FieldOptionName is the from field option name for struct tag.
	FieldOptionName = "field"
	// ContextOptionName is the context option name for struct tag.
	ContextOptionName = "context"
	// SkipOptionName is the skip option name for struct tag.
	SkipOptionName = "skip"
)

type (
	// TagOptions is a map that contains extracted struct tag options.
	TagOptions map[string]string

	// Options are copier options.
	Options struct {
		// Context given to WithContext() method.
		Context map[string]interface{}
		// Reversed reverses struct tag checkings.
		Reversed bool
	}
)

// DeepCopier deep copies a struct to/from a struct.
type DeepCopier struct {
	dst interface{}
	src interface{}
	ctx map[string]interface{}
}

// Copy sets source or destination.
func Copy(src interface{}) *DeepCopier {
	return &DeepCopier{src: src}
}

// WithContext injects the given context into the builder instance.
func (dc *DeepCopier) WithContext(ctx map[string]interface{}) *DeepCopier {
	dc.ctx = ctx
	return dc
}

// To sets the destination.
func (dc *DeepCopier) To(dst interface{}) error {
	dc.dst = dst
	return cp(dc.dst, dc.src, Options{Context: dc.ctx})
}

// From sets the given the source as destination and destination as source.
func (dc *DeepCopier) From(src interface{}) error {
	dc.dst = dc.src
	dc.src = src
	return cp(dc.dst, dc.src, Options{Context: dc.ctx, Reversed: true})
}

// cp is the brand new way to process copy.
func cp(dst interface{}, src interface{}, args ...Options) error {
	var (
		options        = Options{}
		srcValue       = reflect.Indirect(reflect.ValueOf(src))
		dstValue       = reflect.Indirect(reflect.ValueOf(dst))
		srcFieldNames  = getFieldNames(src)
		srcMethodNames = getMethodNames(src)
	)

	if len(args) > 0 {
		options = args[0]
	}

	if !dstValue.CanAddr() {
		return fmt.Errorf("destination %+v is unaddressable", dstValue.Interface())
	}

	for _, m := range srcMethodNames {
		name, opts := getRelatedField(dst, m)
		if name == "" {
			continue
		}

		method := reflect.ValueOf(src).MethodByName(m)
		if !method.IsValid() {
			return fmt.Errorf("method %s is invalid", m)
		}

		var (
			dstFieldType, _ = dstValue.Type().FieldByName(name)
			dstFieldValue   = dstValue.FieldByName(name)
		)

		withContext := false
		if _, ok := opts[ContextOptionName]; ok {
			withContext = true
		}

		args := []reflect.Value{}
		if withContext {
			args = []reflect.Value{reflect.ValueOf(options.Context)}
		}

		result := method.Call(args)[0]
		if result.Type().AssignableTo(dstFieldType.Type) {
			dstFieldValue.Set(result)
		}
	}

	for _, f := range srcFieldNames {
		var (
			srcFieldValue               = srcValue.FieldByName(f)
			srcFieldType, srcFieldFound = srcValue.Type().FieldByName(f)
			srcFieldName                = srcFieldType.Name
			dstFieldName                = srcFieldName
			tagOptions                  TagOptions
		)

		if options.Reversed {
			tagOptions = getTagOptions(srcFieldType.Tag.Get(TagName))
			if v, ok := tagOptions[FieldOptionName]; ok && v != "" {
				dstFieldName = v
			}
		} else {
			if name, opts := getRelatedField(dst, srcFieldName); name != "" {
				dstFieldName, tagOptions = name, opts
			}
		}

		if _, ok := tagOptions[SkipOptionName]; ok {
			continue
		}

		var (
			dstFieldType, dstFieldFound = dstValue.Type().FieldByName(dstFieldName)
			dstFieldValue               = dstValue.FieldByName(dstFieldName)
		)

		// Ptr -> Value
		if srcFieldType.Type.Kind() == reflect.Ptr && !srcFieldValue.IsNil() && dstFieldType.Type.Kind() != reflect.Ptr {
			dstFieldValue.Set(reflect.Indirect(srcFieldValue))
			continue
		}

		if srcFieldFound && dstFieldFound && srcFieldType.Type.AssignableTo(dstFieldType.Type) {
			dstFieldValue.Set(srcFieldValue)
		}
	}

	return nil
}

// getTagOptions parses deepcopier tag field and returns options.
func getTagOptions(value string) TagOptions {
	options := TagOptions{}

	for _, opt := range strings.Split(value, ";") {
		o := strings.Split(opt, ":")

		// deepcopier:"keyword; without; value;"
		if len(o) == 1 {
			options[o[0]] = ""
		}

		// deepcopier:"key:value; anotherkey:anothervalue"
		if len(o) == 2 {
			options[strings.TrimSpace(o[0])] = strings.TrimSpace(o[1])
		}
	}

	return options
}

// getRelatedField returns first matching field.
func getRelatedField(instance interface{}, name string) (string, TagOptions) {
	var (
		value      = reflect.Indirect(reflect.ValueOf(instance))
		fieldName  string
		tagOptions TagOptions
	)

	for i := 0; i < value.NumField(); i++ {
		var (
			vField     = value.Field(i)
			tField     = value.Type().Field(i)
			tagOptions = getTagOptions(tField.Tag.Get(TagName))
		)

		if tField.Type.Kind() == reflect.Struct && tField.Anonymous {
			if n, o := getRelatedField(vField.Interface(), name); n != "" {
				return n, o
			}
		}

		if v, ok := tagOptions[FieldOptionName]; ok && v == name {
			return tField.Name, tagOptions
		}

		if tField.Name == name {
			return tField.Name, tagOptions
		}
	}

	return fieldName, tagOptions
}

// getMethodNames returns instance's method names.
func getMethodNames(instance interface{}) []string {
	var methods []string

	t := reflect.TypeOf(instance)
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	return methods
}

// getFieldNames returns instance's field names.
func getFieldNames(instance interface{}) []string {
	var (
		fields []string
		v      = reflect.Indirect(reflect.ValueOf(instance))
		t      = v.Type()
	)

	if t.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < v.NumField(); i++ {
		var (
			vField = v.Field(i)
			tField = v.Type().Field(i)
		)

		if tField.Type.Kind() == reflect.Struct && tField.Anonymous {
			fields = append(fields, getFieldNames(vField.Interface())...)
			continue
		}

		fields = append(fields, tField.Name)
	}

	return fields
}
