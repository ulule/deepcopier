package deepcopier

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCopyTo_Struct(t *testing.T) {
	var (
		data       = NewEntityData()
		entity     = NewEntity(data)
		entityCopy = &EntityCopy{}
		expected   = NewEntityCopy(data)
	)

	assert.Nil(t, Copy(entity).WithContext(data.MethodContext).To(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.String, entityCopy.String},
		2:  {expected.Time, entityCopy.Time},
		4:  {expected.Float64, entityCopy.Float64},
		5:  {expected.Int, entityCopy.Int},
		9:  {expected.Int64, entityCopy.Int64},
		10: {expected.UInt, entityCopy.UInt},
		14: {expected.UInt64, entityCopy.UInt64},
		15: {expected.StringSlice, entityCopy.StringSlice},
		16: {expected.IntSlice, entityCopy.IntSlice},
		17: {expected.IntMethod, entityCopy.IntMethod},
		21: {expected.Int64Method, entityCopy.Int64Method},
		22: {expected.UIntMethod, entityCopy.UIntMethod},
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
		data       = NewEntityData()
		entity     = NewEntity(data)
		entityCopy = &EntityCopyExtended{}
		expected   = NewEntityCopyExtended(data)
	)

	assert.Nil(t, Copy(entity).WithContext(data.MethodContext).To(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.String, entityCopy.String},
		2:  {expected.Time, entityCopy.Time},
		4:  {expected.Float64, entityCopy.Float64},
		5:  {expected.Int, entityCopy.Int},
		9:  {expected.Int64, entityCopy.Int64},
		10: {expected.UInt, entityCopy.UInt},
		14: {expected.UInt64, entityCopy.UInt64},
		15: {expected.StringSlice, entityCopy.StringSlice},
		16: {expected.IntSlice, entityCopy.IntSlice},
		17: {expected.IntMethod, entityCopy.IntMethod},
		21: {expected.Int64Method, entityCopy.Int64Method},
		22: {expected.UIntMethod, entityCopy.UIntMethod},
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
		data       = NewEntityData()
		entity     = &Entity{}
		entityCopy = NewEntityCopy(data)
		expected   = NewEntity(data)
	)

	assert.Nil(t, Copy(entity).From(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.String, entity.String},
		2:  {expected.Time, entity.Time},
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
		data       = NewEntityData()
		entity     = &Entity{}
		entityCopy = NewEntityCopyExtended(data)
		expected   = NewEntity(data)
	)

	assert.Nil(t, Copy(entity).From(entityCopy))

	for i, tt := range []struct {
		in  interface{}
		out interface{}
	}{
		1:  {expected.String, entity.String},
		2:  {expected.Time, entity.Time},
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

type EntityData struct {
	Time           time.Time
	TimePtr        *time.Time
	String         string
	StringPtr      *string
	Int            int
	IntPtr         *int
	Int64          int64
	Int64Ptr       *int64
	Uint           uint
	UintPtr        *uint
	Uint64         uint64
	Uint64Ptr      *uint64
	Float64        float64
	Float64Ptr     *float64
	StringSlice    []string
	StringSlicePtr *[]string
	StringPtrSlice []*string
	IntSlice       []int
	IntSlicePtr    *[]int
	IntPtrSlice    []*int
	Struct         RelatedEntity
	StructPtr      *RelatedEntity
	Map            map[string]interface{}
	MapPtr         *map[string]interface{}
	NullString     null.String
	PQNullTime     pq.NullTime
	SQLNullString  sql.NullString
	SQLNullInt64   sql.NullInt64
	MethodContext  map[string]interface{}
}

func NewEntityData() *EntityData {
	var (
		now              = time.Now()
		str              = "hello"
		integer          = 10
		integerPtr       = &integer
		integer64        = int64(64)
		integer64Ptr     = &integer64
		uinteger         = uint(10)
		uintegerPtr      = &uinteger
		uinteger64       = uint64(64)
		uinteger64Ptr    = &uinteger64
		f64              = float64(64)
		f64Ptr           = &f64
		stringSlice      = []string{"Chuck", "Norris"}
		stringSlicePtr   = &stringSlice
		stringPtrSlice   = []*string{&str}
		integerSlice     = []int{0, 8, 15}
		integerSlicePtr  = &integerSlice
		integerPtrSlice  = []*int{integerPtr}
		relatedEntity    = RelatedEntity{String: "I am the related entity"}
		relatedEntityPtr = &relatedEntity
		mp               = map[string]interface{}{"message": "ok", "valid": true}
		mpPtr            = &mp
		methodContext    = map[string]interface{}{"version": "1"}
	)

	return &EntityData{
		Time:           now,
		TimePtr:        &now,
		String:         str,
		StringPtr:      &str,
		Int:            integer,
		IntPtr:         integerPtr,
		Int64:          integer64,
		Int64Ptr:       integer64Ptr,
		Uint:           uinteger,
		UintPtr:        uintegerPtr,
		Uint64:         uinteger64,
		Uint64Ptr:      uinteger64Ptr,
		Float64:        f64,
		Float64Ptr:     f64Ptr,
		StringSlice:    stringSlice,
		StringSlicePtr: stringSlicePtr,
		StringPtrSlice: stringPtrSlice,
		IntSlice:       integerSlice,
		IntSlicePtr:    integerSlicePtr,
		IntPtrSlice:    integerPtrSlice,
		Struct:         relatedEntity,
		StructPtr:      relatedEntityPtr,
		Map:            mp,
		MapPtr:         mpPtr,
		MethodContext:  methodContext,
	}
}

type RelatedEntity struct {
	String string
}

type Entity struct {
	String       string
	Time         time.Time
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
}

func NewEntity(data *EntityData) *Entity {
	return &Entity{
		String:       data.String,
		Time:         data.Time,
		AnInt:        data.Int,
		AnInt64:      data.Int64,
		AnUInt:       data.Uint,
		AnUInt64:     data.Uint64,
		AFloat64:     data.Float64,
		AStringSlice: data.StringSlice,
		AnIntSlice:   data.IntSlice,
		ANullString:  data.NullString,
	}
}

func (e *Entity) IntMethod() int                                    { return e.AnInt }
func (e *Entity) Int64Method() int64                                { return e.AnInt64 }
func (e *Entity) UIntMethod() uint                                  { return e.AnUInt }
func (e *Entity) UInt64Method() uint64                              { return e.AnUInt64 }
func (e *Entity) Float64Method() float64                            { return e.AFloat64 }
func (e *Entity) MethodWithDifferentName() string                   { return e.String }
func (e *Entity) MethodWithContext(c map[string]interface{}) string { return c["version"].(string) }

type EntityCopy struct {
	String            string      `json:"string"`
	Time              time.Time   `json:"time"`
	Int               int         `json:"an_int" deepcopier:"field:AnInt"`
	Int64             int64       `json:"an_int64" deepcopier:"field:AnInt64"`
	UInt              uint        `json:"an_uint" deepcopier:"field:AnUInt"`
	UInt64            uint64      `json:"an_uint64" deepcopier:"field:AnUInt64"`
	Float64           float64     `json:"a_float64" deepcopier:"field:AFloat64"`
	NullString        null.String `json:"a_null_string" deepcopier:"field:ANullString"`
	StringSlice       []string    `json:"a_string_slice" deepcopier:"field:AStringSlice"`
	IntSlice          []int       `json:"an_int_slice" deepcopier:"field:AnIntSlice"`
	IntMethod         int         `json:"int_method"`
	Int64Method       int64       `json:"int64_method"`
	UIntMethod        uint        `json:"uint_method"`
	UInt64Method      uint64      `json:"uint64_method"`
	MethodWithContext string      `json:"method_with_context" deepcopier:"context"`
	SuperMethod       string      `json:"super_method" deepcopier:"field:MethodWithDifferentName"`
}

func NewEntityCopy(data *EntityData) *EntityCopy {
	return &EntityCopy{
		String:            data.String,
		Time:              data.Time,
		Float64:           data.Float64,
		Int:               data.Int,
		Int64:             data.Int64,
		UInt:              data.Uint,
		UInt64:            data.Uint64,
		StringSlice:       data.StringSlice,
		IntSlice:          data.IntSlice,
		NullString:        data.NullString,
		IntMethod:         data.Int,
		Int64Method:       data.Int64,
		UIntMethod:        data.Uint,
		UInt64Method:      data.Uint64,
		MethodWithContext: "1",
		SuperMethod:       "hello",
	}
}

type EntityCopyExtended struct {
	EntityCopy
}

func NewEntityCopyExtended(data *EntityData) *EntityCopyExtended {
	return &EntityCopyExtended{EntityCopy: *NewEntityCopy(data)}
}

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

// assertEqual is a verbose version of assert.Equal()
func assertEqual(t *testing.T, idx int, in interface{}, out interface{}) {
	assert.Equal(t, in, out, fmt.Sprintf("%d -- %v not equal to %v", idx, in, out))
}
