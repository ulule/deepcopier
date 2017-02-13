package deepcopier

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/guregu/null"
	"github.com/lib/pq"
	"github.com/oleiade/reflections"
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

// DeepCopier deep copies a struct to/from a struct.
type DeepCopier struct {
	Source      interface{}
	Destination interface{}
	Tagged      interface{}
	Context     map[string]interface{}
	Reversed    bool
}

// FieldOptions contains options passed to SetField method.
type FieldOptions struct {
	SourceField      string
	DestinationField string
	WithContext      bool
	Skip             bool
}

// NewFieldOptions returns a FieldOptions instance for the given instance's field.
func NewFieldOptions(instance interface{}, field string, reversed bool) *FieldOptions {
	fieldOptions := &FieldOptions{
		SourceField:      field,
		DestinationField: field,
		WithContext:      false,
		Skip:             false,
	}

	tagOptions, _ := reflections.GetFieldTag(instance, field, TagName)

	if tagOptions == "" {
		return fieldOptions
	}

	opts := GetTagOptions(tagOptions)

	if _, ok := opts[FieldOptionName]; ok {
		fieldName := opts[FieldOptionName]

		if !reversed {
			fieldOptions.SourceField = fieldName
		} else {
			fieldOptions.DestinationField = fieldName
		}
	}

	if _, ok := opts[ContextOptionName]; ok {
		fieldOptions.WithContext = true
	}

	if _, ok := opts[SkipOptionName]; ok {
		fieldOptions.Skip = true
	}

	return fieldOptions
}

// Copy sets the source.
func Copy(source interface{}) *DeepCopier {
	return &DeepCopier{
		Source:   source,
		Reversed: false,
	}
}

// To sets the given tagged struct as destination struct.
// Source -> Destination
func (dc *DeepCopier) To(tagged interface{}) error {
	dc.Destination = tagged
	dc.Tagged = tagged

	return dc.ProcessCopy()
}

// From sets the given tagged struct as source and the current source as destination.
// Source <- Destination
func (dc *DeepCopier) From(tagged interface{}) error {
	dc.Destination = dc.Source
	dc.Source = tagged
	dc.Tagged = tagged
	dc.Reversed = true

	return dc.ProcessCopy()
}

// ProcessCopy processes copy.
func (dc *DeepCopier) ProcessCopy() error {
	var (
		fields      = []string{}
		taggedValue = reflect.ValueOf(dc.Tagged).Elem()
		taggedType  = taggedValue.Type()
	)

	for i := 0; i < taggedValue.NumField(); i++ {
		var (
			fv = taggedValue.Field(i)
			ft = taggedType.Field(i)
		)

		// Embedded struct
		if ft.Anonymous {
			f, _ := reflections.Fields(fv.Interface())
			fields = append(fields, f...)
		} else {
			fields = append(fields, ft.Name)
		}
	}

	for _, field := range fields {
		if err := dc.SetField(NewFieldOptions(dc.Tagged, field, dc.Reversed)); err != nil {
			return err
		}
	}

	return nil
}

// -----------------------------------------------------------------------------
// Options
// -----------------------------------------------------------------------------

// WithContext injects the given context into the builder instance.
func (dc *DeepCopier) WithContext(context map[string]interface{}) *DeepCopier {
	dc.Context = context
	return dc
}

// -----------------------------------------------------------------------------
// Struct tags
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// Field Setters
// -----------------------------------------------------------------------------

// SetField sets the value of the given field.
func (dc *DeepCopier) SetField(options *FieldOptions) error {
	if options.Skip {
		return nil
	}

	if dc.Reversed {
		has, _ := reflections.HasField(dc.Destination, options.DestinationField)
		if !has {
			return nil
		}
	}

	has, _ := reflections.HasField(dc.Source, options.SourceField)
	if !has {
		err := dc.HandleMethod(options)
		if err != nil {
			has, _ = reflections.HasField(dc.Destination, options.DestinationField)
			if has {
				return nil
			}
		}
		return nil
	}

	kind, _ := reflections.GetFieldKind(dc.Source, options.SourceField)
	if kind == reflect.Struct {
		if err := dc.HandleStructField(options); err != nil {
			return err
		}
		return nil
	}

	if err := dc.HandleField(options); err != nil {
		return err
	}

	return nil
}

// SetFieldValue Sets the given value to the given field.
func (dc *DeepCopier) SetFieldValue(entity interface{}, name string, value reflect.Value) error {
	kind := value.Kind()

	if kind == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
		kind = value.Kind()
	}

	// Maps
	if kind == reflect.Map {
		switch v := value.Interface().(type) {
		case map[string]interface{}, map[string]string, map[string]map[string]string, map[string]map[string]map[string]string:
			if err := reflections.SetField(entity, name, v); err != nil {
				return err
			}
			return nil
		}
	}

	// Structs
	if kind == reflect.Struct {
		switch v := value.Interface().(type) {
		case time.Time, pq.NullTime, null.String:
			if err := reflections.SetField(entity, name, v); err != nil {
				return err
			}
			return nil
		}
	}

	// Slices
	if kind == reflect.Slice {
		switch v := value.Interface().(type) {
		case []int8, []int16, []int32, []int64, []int, []uint8, []uint16, []uint32, []uint64, []uint, []float32, []float64, []string, []bool:
			if err := reflections.SetField(entity, name, v); err != nil {
				return err
			}
			return nil
		}
	}

	// Reflect
	switch kind {
	case reflect.Int8:
		if err := reflections.SetField(entity, name, int8(value.Int())); err != nil {
			return err
		}
		return nil
	case reflect.Int16:
		if err := reflections.SetField(entity, name, int16(value.Int())); err != nil {
			return err
		}
		return nil
	case reflect.Int32:
		if err := reflections.SetField(entity, name, int32(value.Int())); err != nil {
			return err
		}
		return nil
	case reflect.Int64:
		if err := reflections.SetField(entity, name, value.Int()); err != nil {
			return err
		}
		return nil
	case reflect.Int:
		if err := reflections.SetField(entity, name, int(value.Int())); err != nil {
			return err
		}
		return nil
	case reflect.Uint8:
		if err := reflections.SetField(entity, name, uint8(value.Uint())); err != nil {
			return err
		}
		return nil
	case reflect.Uint16:
		if err := reflections.SetField(entity, name, uint16(value.Uint())); err != nil {
			return err
		}
		return nil
	case reflect.Uint32:
		if err := reflections.SetField(entity, name, uint32(value.Uint())); err != nil {
			return err
		}
		return nil
	case reflect.Uint64:
		if err := reflections.SetField(entity, name, value.Uint()); err != nil {
			return err
		}
		return nil
	case reflect.Uint:
		if err := reflections.SetField(entity, name, uint(value.Uint())); err != nil {
			return err
		}
		return nil
	case reflect.Float32:
		if err := reflections.SetField(entity, name, float32(value.Float())); err != nil {
			return err
		}
		return nil
	case reflect.Float64:
		if err := reflections.SetField(entity, name, value.Float()); err != nil {
			return err
		}
		return nil
	case reflect.String:
		if err := reflections.SetField(entity, name, value.String()); err != nil {
			return err
		}
		return nil
	case reflect.Bool:
		if err := reflections.SetField(entity, name, value.Bool()); err != nil {
			return err
		}
		return nil
	}

	return nil
}

// -----------------------------------------------------------------------------
// Field Type Handlers
// -----------------------------------------------------------------------------

// HandleStructField sets the value for the given supported struct field.
func (dc *DeepCopier) HandleStructField(options *FieldOptions) error {
	f, err := reflections.GetField(dc.Source, options.SourceField)
	if err != nil {
		return err
	}

	switch v := f.(type) {
	case pq.NullTime:
		if v.Valid {
			if err := reflections.SetField(dc.Destination, options.DestinationField, &v.Time); err != nil {
				return err
			}
		}
	case time.Time:
		if err := reflections.SetField(dc.Destination, options.DestinationField, v); err != nil {
			return err
		}
	}

	return nil
}

// HandleField sets value for the given field.
func (dc *DeepCopier) HandleField(options *FieldOptions) error {
	v, err := reflections.GetField(dc.Source, options.SourceField)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(v)
	if err := dc.SetFieldValue(dc.Destination, options.DestinationField, value); err != nil {
		return err
	}

	return nil
}

// HandleMethod tries to call method on model and sets result in resource field.
func (dc *DeepCopier) HandleMethod(options *FieldOptions) error {
	if dc.Reversed {
		return nil
	}

	method := reflect.ValueOf(dc.Source).MethodByName(options.SourceField)
	if !method.IsValid() {
		return fmt.Errorf("Method %s does not exist", options.SourceField)
	}

	var results []reflect.Value
	if options.WithContext {
		results = method.Call([]reflect.Value{reflect.ValueOf(dc.Context)})
	} else {
		results = method.Call([]reflect.Value{})
	}

	if err := dc.SetFieldValue(dc.Destination, options.DestinationField, results[0]); err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------------
// Refacto
// -----------------------------------------------------------------------------

// Options are copier options.
type Options struct {
	// Context given to WithContext() method.
	Context map[string]interface{}
}

func getMethods(t reflect.Type) []string {
	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}
	return methods
}

// InStringSlice checks if the given string is in the given slice of string
func InStringSlice(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}
	return false
}

// Copier is the brand new way to process copy.
func Copier(dst interface{}, src interface{}, args ...Options) error {
	var (
		options    = Options{}
		srcValue   = reflect.Indirect(reflect.ValueOf(src))
		srcType    = srcValue.Type()
		srcMethods = getMethods(srcType)
		dstValue   = reflect.Indirect(reflect.ValueOf(dst))
	)

	// Pointer only for receiver
	if !dstValue.CanAddr() {
		return errors.New("dst value is unaddressable")
	}

	if len(args) > 0 {
		options = args[0]
	}

	for i := 0; i < srcValue.NumField(); i++ {
		var (
			srcFieldValue = srcValue.Field(i)
			srcFieldType  = srcValue.Type().Field(i)
			srcFieldName  = srcFieldType.Name
		)

		if !srcFieldValue.IsValid() {
			continue
		}

		for ii := 0; ii < dstValue.NumField(); ii++ {
			var (
				dstFieldValue = dstValue.Field(ii)
				dstFieldType  = dstValue.Type().Field(ii)
				dstFieldName  = dstFieldType.Name
				dstFieldTag   = dstFieldType.Tag.Get(TagName)
				// Options
				fieldName   = dstFieldName
				withContext = false
			)

			tagOptions := GetTagOptions(dstFieldTag)

			// If skip option is set, bypass copy.
			if v, ok := tagOptions[SkipOptionName]; ok && v != "" {
				continue
			}

			// Get real source field / method name src struct tag.
			if v, ok := tagOptions[FieldOptionName]; ok && v != "" {
				fieldName = v
			}

			// Give context as function argument?
			if v, ok := tagOptions[ContextOptionName]; ok && v != "" {
				withContext = true
			}

			if InStringSlice(srcMethods, fieldName) {
				method := reflect.ValueOf(src).MethodByName(fieldName)

				if !method.IsValid() {
					return fmt.Errorf("method %v in source is not valid", fieldName)
				}

				var results []reflect.Value

				if withContext {
					results = method.Call([]reflect.Value{reflect.ValueOf(options.Context)})
				} else {
					results = method.Call([]reflect.Value{})
				}

				dstFieldValue.Set(results[0])

				continue
			}

			if fieldName != srcFieldName {
				continue
			}

			// Ptr -> Value
			if srcFieldType.Type.Kind() == reflect.Ptr && !srcFieldValue.IsNil() && dstFieldType.Type.Kind() != reflect.Ptr {
				dstFieldValue.Set(reflect.Indirect(srcFieldValue))
				continue
			}

			if srcFieldType.Type.AssignableTo(dstFieldType.Type) {
				dstFieldValue.Set(srcFieldValue)
			}
		}
	}

	return nil
}
