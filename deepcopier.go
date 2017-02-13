package deepcopier

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	// TagName is struct field tag name.
	TagName = "deepcopier"
	// FieldOptionName is the from field option name for struct tag.
	FieldOptionName = "field"
	// ContextOptionName is the context option name for struct tag.
	ContextOptionName = "context"
	// SkipOptionName is the skip option name for struct tag.
	SkipOptionName = "skip"
)

// ----------------------------------------------------------------------------
// Struct tag options
// ----------------------------------------------------------------------------

// TagOptions are struct tag options.
type TagOptions map[string]string

// ----------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------

// Options are copier options.
type Options struct {
	// Context given to WithContext() method.
	Context map[string]interface{}
	// Reversed reverses struct tag checkings.
	Reversed bool
}

// ----------------------------------------------------------------------------
// Deepcopier
// ----------------------------------------------------------------------------

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

// To sets the given the destination.
func (dc *DeepCopier) To(dst interface{}) error {
	dc.dst = dst
	return cp(dc.dst, dc.src, Options{
		Context: dc.ctx,
	})
}

// From sets the given the source as destination and destination as source.
func (dc *DeepCopier) From(src interface{}) error {
	dc.dst = dc.src
	dc.src = src
	return cp(dc.dst, dc.src, Options{
		Context:  dc.ctx,
		Reversed: true,
	})
}

// cp is the brand new way to process copy.
func cp(dst interface{}, src interface{}, args ...Options) error {
	var (
		options        = Options{}
		srcValue       = reflect.Indirect(reflect.ValueOf(src))
		srcFieldNames  = getFieldNames(src)
		srcMethodNames = getMethodNames(src)
		dstValue       = reflect.Indirect(reflect.ValueOf(dst))
	)

	// Options are given
	if len(args) > 0 {
		options = args[0]
	}

	// Pointer only for receiver
	if !dstValue.CanAddr() {
		return errors.New("dst value is unaddressable")
	}

	//
	// Methods
	//

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
			withContext     = false
		)

		if _, ok := opts[ContextOptionName]; ok {
			withContext = true
		}

		var args []reflect.Value
		if withContext {
			args = []reflect.Value{reflect.ValueOf(options.Context)}
		}

		result := method.Call(args)[0]
		if result.Type().AssignableTo(dstFieldType.Type) {
			dstFieldValue.Set(result)
		}
	}

	//
	// Fields
	//

	for _, f := range srcFieldNames {
		var (
			srcFieldValue                = srcValue.FieldByName(f)
			srcFieldType, srcFieldTypeOK = srcValue.Type().FieldByName(f)
			srcFieldName                 = srcFieldType.Name
			dstFieldName                 = srcFieldName
			tagOptions                   TagOptions
		)

		if options.Reversed {
			tagOptions = getTagOptions(srcFieldType.Tag.Get(TagName))

			if v, ok := tagOptions[FieldOptionName]; ok && v != "" {
				dstFieldName = v
			}
		} else {
			if name, opts := getRelatedField(dst, srcFieldName); name != "" {
				tagOptions = opts
				dstFieldName = name
			}
		}

		// Struct tag -- deepcopier:"skip"
		if _, ok := tagOptions[SkipOptionName]; ok {
			continue
		}

		var (
			dstFieldType, dstFieldTypeOK = dstValue.Type().FieldByName(dstFieldName)
			dstFieldValue                = dstValue.FieldByName(dstFieldName)
		)

		// Ptr -> Value
		if srcFieldType.Type.Kind() == reflect.Ptr && !srcFieldValue.IsNil() && dstFieldType.Type.Kind() != reflect.Ptr {
			dstFieldValue.Set(reflect.Indirect(srcFieldValue))
			continue
		}

		if srcFieldTypeOK && dstFieldTypeOK && srcFieldType.Type.AssignableTo(dstFieldType.Type) {
			dstFieldValue.Set(srcFieldValue)
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

// getTagOptions parses deepcopier tag field and returns options.
func getTagOptions(value string) TagOptions {
	options := TagOptions{}

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

// getRelatedField returns first matching field.
func getRelatedField(instance interface{}, name string) (string, TagOptions) {
	var (
		value      = reflect.Indirect(reflect.ValueOf(instance))
		fieldName  string
		tagOptions TagOptions
	)

	for i := 0; i < value.NumField(); i++ {
		var (
			v          = value.Field(i)
			t          = value.Type().Field(i)
			tagOptions = getTagOptions(t.Tag.Get(TagName))
		)

		if t.Type.Kind() == reflect.Struct && t.Anonymous {
			if n, o := getRelatedField(v.Interface(), name); n != "" {
				return n, o
			}
		}

		if v, ok := tagOptions[FieldOptionName]; ok && v == name {
			return t.Name, tagOptions
		}

		if t.Name == name {
			return t.Name, tagOptions
		}
	}

	return fieldName, tagOptions
}

// getMethodNames returns instance's method names.
func getMethodNames(instance interface{}) []string {
	var (
		t       = reflect.TypeOf(instance)
		methods []string
	)

	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	return methods
}

// getFieldNames returns instance's field names.
func getFieldNames(instance interface{}) []string {
	var (
		v      = reflect.Indirect(reflect.ValueOf(instance))
		t      = v.Type()
		fields []string
	)

	if t.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	return fields
}
