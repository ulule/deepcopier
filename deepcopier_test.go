package deepcopier

import (
	"fmt"
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func TestCopyTo_Struct(t *testing.T) {
	var (
		now        = time.Now()
		entity     = NewEntity(now)
		entityCopy = &EntityCopy{}
		expected   = NewEntityCopy(now)
	)

	assert.Nil(t, Copy(entity).WithContext(map[string]interface{}{"version": "1"}).To(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.Title, entityCopy.Title},
		2:  {expected.Date, entityCopy.Date},
		3:  {expected.Float32, entityCopy.Float32},
		4:  {expected.Float64, entityCopy.Float64},
		5:  {expected.Int, entityCopy.Int},
		6:  {expected.Int8, entityCopy.Int8},
		7:  {expected.Int16, entityCopy.Int16},
		8:  {expected.Int32, entityCopy.Int32},
		9:  {expected.Int64, entityCopy.Int64},
		10: {expected.UInt, entityCopy.UInt},
		11: {expected.UInt8, entityCopy.UInt8},
		12: {expected.UInt16, entityCopy.UInt16},
		13: {expected.UInt32, entityCopy.UInt32},
		14: {expected.UInt64, entityCopy.UInt64},
		15: {expected.StringSlice, entityCopy.StringSlice},
		16: {expected.IntSlice, entityCopy.IntSlice},
		17: {expected.IntMethod, entityCopy.IntMethod},
		18: {expected.Int8Method, entityCopy.Int8Method},
		19: {expected.Int16Method, entityCopy.Int16Method},
		20: {expected.Int32Method, entityCopy.Int32Method},
		21: {expected.Int64Method, entityCopy.Int64Method},
		22: {expected.UIntMethod, entityCopy.UIntMethod},
		23: {expected.UInt8Method, entityCopy.UInt8Method},
		24: {expected.UInt16Method, entityCopy.UInt16Method},
		25: {expected.UInt32Method, entityCopy.UInt32Method},
		26: {expected.UInt64Method, entityCopy.UInt64Method},
		27: {expected.MethodWithContext, entityCopy.MethodWithContext},
		28: {expected.SuperMethod, entityCopy.SuperMethod},
		29: {expected.StringSlice, entityCopy.StringSlice},
		30: {expected.IntSlice, entityCopy.IntSlice},
	} {
		assertEqual(t, i, tt.in, tt.out)
	}
}

func TestCopyTo_AnonymousStruct(t *testing.T) {
	var (
		now        = time.Now()
		entity     = NewEntity(now)
		entityCopy = &EntityCopyExtended{}
		expected   = NewEntityCopyExtended(now)
	)

	assert.Nil(t, Copy(entity).WithContext(map[string]interface{}{"version": "1"}).To(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.Title, entityCopy.Title},
		2:  {expected.Date, entityCopy.Date},
		3:  {expected.Float32, entityCopy.Float32},
		4:  {expected.Float64, entityCopy.Float64},
		5:  {expected.Int, entityCopy.Int},
		6:  {expected.Int8, entityCopy.Int8},
		7:  {expected.Int16, entityCopy.Int16},
		8:  {expected.Int32, entityCopy.Int32},
		9:  {expected.Int64, entityCopy.Int64},
		10: {expected.UInt, entityCopy.UInt},
		11: {expected.UInt8, entityCopy.UInt8},
		12: {expected.UInt16, entityCopy.UInt16},
		13: {expected.UInt32, entityCopy.UInt32},
		14: {expected.UInt64, entityCopy.UInt64},
		15: {expected.StringSlice, entityCopy.StringSlice},
		16: {expected.IntSlice, entityCopy.IntSlice},
		17: {expected.IntMethod, entityCopy.IntMethod},
		18: {expected.Int8Method, entityCopy.Int8Method},
		19: {expected.Int16Method, entityCopy.Int16Method},
		20: {expected.Int32Method, entityCopy.Int32Method},
		21: {expected.Int64Method, entityCopy.Int64Method},
		22: {expected.UIntMethod, entityCopy.UIntMethod},
		23: {expected.UInt8Method, entityCopy.UInt8Method},
		24: {expected.UInt16Method, entityCopy.UInt16Method},
		25: {expected.UInt32Method, entityCopy.UInt32Method},
		26: {expected.UInt64Method, entityCopy.UInt64Method},
		27: {expected.MethodWithContext, entityCopy.MethodWithContext},
		28: {expected.SuperMethod, entityCopy.SuperMethod},
		29: {expected.StringSlice, entityCopy.StringSlice},
		30: {expected.IntSlice, entityCopy.IntSlice},
	} {
		assertEqual(t, i, tt.in, tt.out)
	}
}

func TestCopyFrom_Struct(t *testing.T) {
	var (
		now        = time.Now()
		entity     = &Entity{}
		entityCopy = NewEntityCopy(now)
		expected   = NewEntity(now)
	)

	assert.Nil(t, Copy(entity).From(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.Name, entity.Name},
		2:  {expected.Date, entity.Date},
		3:  {expected.AFloat32, entity.AFloat32},
		4:  {expected.AFloat64, entity.AFloat64},
		5:  {expected.AnInt, entity.AnInt},
		6:  {expected.AnInt8, entity.AnInt8},
		7:  {expected.AnInt16, entity.AnInt16},
		8:  {expected.AnInt32, entity.AnInt32},
		9:  {expected.AnInt64, entity.AnInt64},
		10: {expected.AnUInt, entity.AnUInt},
		11: {expected.AnUInt8, entity.AnUInt8},
		12: {expected.AnUInt16, entity.AnUInt16},
		13: {expected.AnUInt32, entity.AnUInt32},
		14: {expected.AnUInt64, entity.AnUInt64},
		15: {expected.AStringSlice, entity.AStringSlice},
		16: {expected.AnIntSlice, entity.AnIntSlice},
	} {
		assertEqual(t, i, tt.in, tt.out)
	}
}

func TestCopyFrom_AnonymousStruct(t *testing.T) {
	var (
		now        = time.Now()
		entity     = &Entity{}
		entityCopy = NewEntityCopyExtended(now)
		expected   = NewEntity(now)
	)

	assert.Nil(t, Copy(entity).From(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.Name, entity.Name},
		2:  {expected.Date, entity.Date},
		3:  {expected.AFloat32, entity.AFloat32},
		4:  {expected.AFloat64, entity.AFloat64},
		5:  {expected.AnInt, entity.AnInt},
		6:  {expected.AnInt8, entity.AnInt8},
		7:  {expected.AnInt16, entity.AnInt16},
		8:  {expected.AnInt32, entity.AnInt32},
		9:  {expected.AnInt64, entity.AnInt64},
		10: {expected.AnUInt, entity.AnUInt},
		11: {expected.AnUInt8, entity.AnUInt8},
		12: {expected.AnUInt16, entity.AnUInt16},
		13: {expected.AnUInt32, entity.AnUInt32},
		14: {expected.AnUInt64, entity.AnUInt64},
		15: {expected.AStringSlice, entity.AStringSlice},
		16: {expected.AnIntSlice, entity.AnIntSlice},
	} {
		assertEqual(t, i, tt.in, tt.out)
	}
}

// -----------------------------------------------------------------------------
// Fixtures
// -----------------------------------------------------------------------------

type Entity struct {
	Name         string
	Date         time.Time
	AFloat32     float32
	AFloat64     float64
	AnInt        int
	AnInt8       int8
	AnInt16      int16
	AnInt32      int32
	AnInt64      int64
	AnUInt       uint
	AnUInt8      uint8
	AnUInt16     uint16
	AnUInt32     uint32
	AnUInt64     uint64
	AStringSlice []string
	AnIntSlice   []int
	ANullString  null.String
	APointer     string
}

func NewEntity(now time.Time) *Entity {
	return &Entity{
		Name:         "Chuck Norris",
		Date:         now,
		AFloat32:     float32(10.0),
		AFloat64:     float64(10.0),
		AnInt:        int(10),
		AnInt8:       int8(10),
		AnInt16:      int16(10),
		AnInt32:      int32(10),
		AnInt64:      int64(10),
		AnUInt:       uint(10),
		AnUInt8:      uint8(10),
		AnUInt16:     uint16(10),
		AnUInt32:     uint32(10),
		AnUInt64:     uint64(10),
		AStringSlice: []string{"Chuck", "Norris"},
		AnIntSlice:   []int{0, 8, 15},
		ANullString:  null.StringFrom("I'm null"),
	}
}

func (e *Entity) Float32Method() float32                            { return float32(10.0) }
func (e *Entity) Float64Method() float64                            { return float64(10.0) }
func (e *Entity) IntMethod() int                                    { return int(10) }
func (e *Entity) Int8Method() int8                                  { return int8(10) }
func (e *Entity) Int16Method() int16                                { return int16(10) }
func (e *Entity) Int32Method() int32                                { return int32(10) }
func (e *Entity) Int64Method() int64                                { return int64(10) }
func (e *Entity) UIntMethod() uint                                  { return uint(10) }
func (e *Entity) UInt8Method() uint8                                { return uint8(10) }
func (e *Entity) UInt16Method() uint16                              { return uint16(10) }
func (e *Entity) UInt32Method() uint32                              { return uint32(10) }
func (e *Entity) UInt64Method() uint64                              { return uint64(10) }
func (e *Entity) MethodWithDifferentName() string                   { return "hello" }
func (e *Entity) MethodWithContext(c map[string]interface{}) string { return c["version"].(string) }

type EntityCopy struct {
	Date              time.Time   `json:"date"`
	Title             string      `json:"name" deepcopier:"field:Name"`
	Float32           float32     `json:"a_float32" deepcopier:"field:AFloat32"`
	Float64           float64     `json:"a_float64" deepcopier:"field:AFloat64"`
	Int               int         `json:"an_int" deepcopier:"field:AnInt"`
	Int8              int8        `json:"an_int8" deepcopier:"field:AnInt8"`
	Int16             int16       `json:"an_int16" deepcopier:"field:AnInt16"`
	Int32             int32       `json:"an_int32" deepcopier:"field:AnInt32"`
	Int64             int64       `json:"an_int64" deepcopier:"field:AnInt64"`
	UInt              uint        `json:"an_uint" deepcopier:"field:AnUInt"`
	UInt8             uint8       `json:"an_uint8" deepcopier:"field:AnUInt8"`
	UInt16            uint16      `json:"an_uint16" deepcopier:"field:AnUInt16"`
	UInt32            uint32      `json:"an_uint32" deepcopier:"field:AnUInt32"`
	UInt64            uint64      `json:"an_uint64" deepcopier:"field:AnUInt64"`
	NullString        null.String `json:"a_null_string" deepcopier:"field:ANullString"`
	StringSlice       []string    `json:"a_string_slice" deepcopier:"field:AStringSlice"`
	IntSlice          []int       `json:"an_int_slice" deepcopier:"field:AnIntSlice"`
	IntMethod         int         `json:"int_method"`
	Int8Method        int8        `json:"int8_method"`
	Int16Method       int16       `json:"int16_method"`
	Int32Method       int32       `json:"int32_method"`
	Int64Method       int64       `json:"int64_method"`
	UIntMethod        uint        `json:"uint_method"`
	UInt8Method       uint8       `json:"uint8_method"`
	UInt16Method      uint16      `json:"uint16_method"`
	UInt32Method      uint32      `json:"uint32_method"`
	UInt64Method      uint64      `json:"uint64_method"`
	MethodWithContext string      `json:"method_with_context" deepcopier:"context"`
	SuperMethod       string      `json:"super_method" deepcopier:"field:MethodWithDifferentName"`
}

func NewEntityCopy(now time.Time) *EntityCopy {
	return &EntityCopy{
		Title:             "Chuck Norris",
		Date:              now,
		Float32:           float32(10.0),
		Float64:           float64(10.0),
		Int:               int(10),
		Int8:              int8(10),
		Int16:             int16(10),
		Int32:             int32(10),
		Int64:             int64(10),
		UInt:              uint(10),
		UInt8:             uint8(10),
		UInt16:            uint16(10),
		UInt32:            uint32(10),
		UInt64:            uint64(10),
		StringSlice:       []string{"Chuck", "Norris"},
		IntSlice:          []int{0, 8, 15},
		NullString:        null.StringFrom("I'm null"),
		IntMethod:         int(10),
		Int8Method:        int8(10),
		Int16Method:       int16(10),
		Int32Method:       int32(10),
		Int64Method:       int64(10),
		UIntMethod:        uint(10),
		UInt8Method:       uint8(10),
		UInt16Method:      uint16(10),
		UInt32Method:      uint32(10),
		UInt64Method:      uint64(10),
		MethodWithContext: "1",
		SuperMethod:       "hello",
	}
}

type EntityCopyExtended struct {
	EntityCopy
}

func NewEntityCopyExtended(now time.Time) *EntityCopyExtended {
	return &EntityCopyExtended{EntityCopy: *NewEntityCopy(now)}
}

// ----------------------------------------------------------------------------
// New refacto
// ----------------------------------------------------------------------------

// type R1 struct{ String string }
// type R2 struct{ String string }

// type T1 struct {
// 	Int64            int64
// 	String           string
// 	StringPtr        *string
// 	StringPtrToValue *string
// 	Time             time.Time
// 	TimePtr          *time.Time
// 	TimePtrToValue   *time.Time
// 	SliceString      []string
// 	SliceInt         []int
// 	Map              map[string]interface{}
// 	NullString       null.String
// 	R1               R1
// 	R1Ptr            *R1
// 	R1PtrToValue     *R1
// 	R2               R2
// 	R2Ptr            *R2
// 	R2PtrToValue     *R2
// }

// func (T1) MethodString() string {
// 	return "method string"
// }

// func (T1) AnotherMethodString() string {
// 	return "another method string"
// }

// func (T1) MethodWithContext(c map[string]interface{}) map[string]interface{} {
// 	return c
// }

// type T2 struct {
// 	Int64             int64
// 	Int64Renamed      int64 `deepcopier:"field:Int64"`
// 	String            string
// 	StringPtr         *string
// 	StringPtrToValue  string
// 	Time              time.Time
// 	TimePtr           *time.Time
// 	TimePtrToValue    time.Time
// 	SliceString       []string
// 	SliceInt          []int
// 	Map               map[string]interface{}
// 	NullString        null.String
// 	MethodString      string
// 	MString           string                 `deepcopier:"field:AnotherMethodString"`
// 	MethodWithContext map[string]interface{} `deepcopier:"context:true"`
// 	R1                R1
// 	R1Ptr             *R1
// 	R1PtrToValue      R1
// 	R2                R2
// 	R2Ptr             *R2
// 	R2PtrToValue      R2
// }

// func TestCopier(t *testing.T) {
// 	var (
// 		strPtr        = "ptr"
// 		strPtrToValue = "ptrToValue"
// 		now           = time.Now()
// 		sliceStr      = []string{"Chuck", "Norris"}
// 		sliceInt      = []int{0, 8, 15}
// 		nullStr       = null.StringFrom("I'm null")
// 		mapInterfaces = map[string]interface{}{"message": "ok", "valid": true}
// 		methodContext = map[string]interface{}{"url": "https://ulule.com"}
// 	)

// 	r1 := &R1{String: "r1 string"}

// 	t1 := &T1{
// 		Int64:            1,
// 		String:           "hello",
// 		StringPtr:        &strPtr,
// 		StringPtrToValue: &strPtrToValue,
// 		Time:             now,
// 		TimePtr:          &now,
// 		TimePtrToValue:   &now,
// 		SliceString:      sliceStr,
// 		SliceInt:         sliceInt,
// 		NullString:       nullStr,
// 		Map:              mapInterfaces,
// 		R1:               *r1,
// 		R1Ptr:            r1,
// 		R1PtrToValue:     r1,
// 	}

// 	t2 := &T2{}

// 	options := Options{
// 		Context: methodContext,
// 	}

// 	assert.Nil(t, cp(t2, t1, options))

// 	table := []struct {
// 		in  interface{}
// 		out interface{}
// 	}{
// 		{t1.Int64, t2.Int64},
// 		{t1.Int64, t2.Int64Renamed},
// 		{t1.String, t2.String},
// 		{t1.StringPtr, t2.StringPtr},
// 		{strPtrToValue, t2.StringPtrToValue},
// 		{t1.Time, t2.Time},
// 		{t1.TimePtr, t2.TimePtr},
// 		{*t1.TimePtrToValue, t2.TimePtrToValue},
// 		{t1.SliceString, t2.SliceString},
// 		{t1.SliceInt, t2.SliceInt},
// 		{t1.Map, t2.Map},
// 		{t1.NullString, t2.NullString},
// 		{t1.MethodString(), t2.MethodString},
// 		{t1.AnotherMethodString(), t2.MString},
// 		{t1.MethodWithContext(methodContext), t2.MethodWithContext},
// 		{t1.R1, t2.R1},
// 		{t1.R1Ptr, t2.R1Ptr},
// 		{*r1, t2.R1PtrToValue},
// 	}

// 	for i, tt := range table {
// 		assertEqual(t, i, tt.in, tt.out)
// 	}
// }

// assertEqual is a verbose version of assert.Equal()
func assertEqual(t *testing.T, idx int, in interface{}, out interface{}) {
	assert.Equal(t, in, out, fmt.Sprintf("%d -- %v not equal to %v", idx, in, out))
}
