package deepcopier

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func TestCopyTo(t *testing.T) {
	var (
		is                       = assert.New(t)
		now                      = time.Now()
		user                     = NewUser(now)
		userCopy                 = &UserCopy{}
		userCopyExtended         = &UserCopyExtended{}
		expectedUserCopy         = NewUserCopy(now)
		expectedUserCopyExtended = NewUserCopyExtended(now)
	)

	is.Nil(Copy(user).WithContext(map[string]interface{}{"version": "1"}).To(userCopy))

	userCopyTests := []struct {
		in  interface{}
		out interface{}
	}{
		{expectedUserCopy.Title, userCopy.Title},
		{expectedUserCopy.Date, userCopy.Date},
		{expectedUserCopy.Float32, userCopy.Float32},
		{expectedUserCopy.Float64, userCopy.Float64},
		{expectedUserCopy.Int, userCopy.Int},
		{expectedUserCopy.Int8, userCopy.Int8},
		{expectedUserCopy.Int16, userCopy.Int16},
		{expectedUserCopy.Int32, userCopy.Int32},
		{expectedUserCopy.Int64, userCopy.Int64},
		{expectedUserCopy.UInt, userCopy.UInt},
		{expectedUserCopy.UInt8, userCopy.UInt8},
		{expectedUserCopy.UInt16, userCopy.UInt16},
		{expectedUserCopy.UInt32, userCopy.UInt32},
		{expectedUserCopy.UInt64, userCopy.UInt64},
		{expectedUserCopy.StringSlice, userCopy.StringSlice},
		{expectedUserCopy.IntSlice, userCopy.IntSlice},
		{expectedUserCopy.IntMethod, userCopy.IntMethod},
		{expectedUserCopy.Int8Method, userCopy.Int8Method},
		{expectedUserCopy.Int16Method, userCopy.Int16Method},
		{expectedUserCopy.Int32Method, userCopy.Int32Method},
		{expectedUserCopy.Int64Method, userCopy.Int64Method},
		{expectedUserCopy.UIntMethod, userCopy.UIntMethod},
		{expectedUserCopy.UInt8Method, userCopy.UInt8Method},
		{expectedUserCopy.UInt16Method, userCopy.UInt16Method},
		{expectedUserCopy.UInt32Method, userCopy.UInt32Method},
		{expectedUserCopy.UInt64Method, userCopy.UInt64Method},
		{expectedUserCopy.MethodWithContext, userCopy.MethodWithContext},
		{expectedUserCopy.SuperMethod, userCopy.SuperMethod},
		{expectedUserCopy.StringSlice, userCopy.StringSlice},
		{expectedUserCopy.IntSlice, userCopy.IntSlice},
	}

	for _, tt := range userCopyTests {
		is.Equal(tt.in, tt.out)
	}

	is.Nil(Copy(user).WithContext(map[string]interface{}{"version": "1"}).To(userCopyExtended))

	userCopyExtendedTests := []struct {
		in  interface{}
		out interface{}
	}{
		{expectedUserCopyExtended.Title, userCopyExtended.Title},
		{expectedUserCopyExtended.Date, userCopyExtended.Date},
		{expectedUserCopyExtended.Float32, userCopyExtended.Float32},
		{expectedUserCopyExtended.Float64, userCopyExtended.Float64},
		{expectedUserCopyExtended.Int, userCopyExtended.Int},
		{expectedUserCopyExtended.Int8, userCopyExtended.Int8},
		{expectedUserCopyExtended.Int16, userCopyExtended.Int16},
		{expectedUserCopyExtended.Int32, userCopyExtended.Int32},
		{expectedUserCopyExtended.Int64, userCopyExtended.Int64},
		{expectedUserCopyExtended.UInt, userCopyExtended.UInt},
		{expectedUserCopyExtended.UInt8, userCopyExtended.UInt8},
		{expectedUserCopyExtended.UInt16, userCopyExtended.UInt16},
		{expectedUserCopyExtended.UInt32, userCopyExtended.UInt32},
		{expectedUserCopyExtended.UInt64, userCopyExtended.UInt64},
		{expectedUserCopyExtended.StringSlice, userCopyExtended.StringSlice},
		{expectedUserCopyExtended.IntSlice, userCopyExtended.IntSlice},
		{expectedUserCopyExtended.IntMethod, userCopyExtended.IntMethod},
		{expectedUserCopyExtended.Int8Method, userCopyExtended.Int8Method},
		{expectedUserCopyExtended.Int16Method, userCopyExtended.Int16Method},
		{expectedUserCopyExtended.Int32Method, userCopyExtended.Int32Method},
		{expectedUserCopyExtended.Int64Method, userCopyExtended.Int64Method},
		{expectedUserCopyExtended.UIntMethod, userCopyExtended.UIntMethod},
		{expectedUserCopyExtended.UInt8Method, userCopyExtended.UInt8Method},
		{expectedUserCopyExtended.UInt16Method, userCopyExtended.UInt16Method},
		{expectedUserCopyExtended.UInt32Method, userCopyExtended.UInt32Method},
		{expectedUserCopyExtended.UInt64Method, userCopyExtended.UInt64Method},
		{expectedUserCopyExtended.MethodWithContext, userCopyExtended.MethodWithContext},
		{expectedUserCopyExtended.SuperMethod, userCopyExtended.SuperMethod},
		{expectedUserCopyExtended.StringSlice, userCopyExtended.StringSlice},
		{expectedUserCopyExtended.IntSlice, userCopyExtended.IntSlice},
	}

	for _, tt := range userCopyExtendedTests {
		is.Equal(tt.in, tt.out)
	}
}

func TestCopyFrom(t *testing.T) {
	var (
		is               = assert.New(t)
		now              = time.Now()
		user             = &User{}
		userExpected     = NewUser(now)
		userCopy         = NewUserCopy(now)
		userCopyExtended = NewUserCopyExtended(now)
	)

	is.Nil(Copy(user).From(userCopy))

	userCopyTests := []struct {
		in  interface{}
		out interface{}
	}{
		{userExpected.Name, user.Name},
		{userExpected.Date, user.Date},
		{userExpected.AFloat32, user.AFloat32},
		{userExpected.AFloat64, user.AFloat64},
		{userExpected.AnInt, user.AnInt},
		{userExpected.AnInt8, user.AnInt8},
		{userExpected.AnInt16, user.AnInt16},
		{userExpected.AnInt32, user.AnInt32},
		{userExpected.AnInt64, user.AnInt64},
		{userExpected.AnUInt, user.AnUInt},
		{userExpected.AnUInt8, user.AnUInt8},
		{userExpected.AnUInt16, user.AnUInt16},
		{userExpected.AnUInt32, user.AnUInt32},
		{userExpected.AnUInt64, user.AnUInt64},
		{userExpected.AStringSlice, user.AStringSlice},
		{userExpected.AnIntSlice, user.AnIntSlice},
	}

	for _, tt := range userCopyTests {
		is.Equal(tt.in, tt.out)
	}

	is.Nil(Copy(user).From(userCopyExtended))

	userCopyExtendedTests := []struct {
		in  interface{}
		out interface{}
	}{
		{userExpected.Name, user.Name},
		{userExpected.Date, user.Date},
		{userExpected.AFloat32, user.AFloat32},
		{userExpected.AFloat64, user.AFloat64},
		{userExpected.AnInt, user.AnInt},
		{userExpected.AnInt8, user.AnInt8},
		{userExpected.AnInt16, user.AnInt16},
		{userExpected.AnInt32, user.AnInt32},
		{userExpected.AnInt64, user.AnInt64},
		{userExpected.AnUInt, user.AnUInt},
		{userExpected.AnUInt8, user.AnUInt8},
		{userExpected.AnUInt16, user.AnUInt16},
		{userExpected.AnUInt32, user.AnUInt32},
		{userExpected.AnUInt64, user.AnUInt64},
		{userExpected.AStringSlice, user.AStringSlice},
		{userExpected.AnIntSlice, user.AnIntSlice},
	}

	for _, tt := range userCopyExtendedTests {
		is.Equal(tt.in, tt.out)
	}
}

// -----------------------------------------------------------------------------
// Fixtures
// -----------------------------------------------------------------------------

type User struct {
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

func NewUser(now time.Time) *User {
	return &User{
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

func (u *User) Float32Method() float32 {
	return float32(10.0)
}

func (u *User) Float64Method() float64 {
	return float64(10.0)
}

func (u *User) IntMethod() int {
	return int(10)
}

func (u *User) Int8Method() int8 {
	return int8(10)
}

func (u *User) Int16Method() int16 {
	return int16(10)
}

func (u *User) Int32Method() int32 {
	return int32(10)
}

func (u *User) Int64Method() int64 {
	return int64(10)
}

func (u *User) UIntMethod() uint {
	return uint(10)
}

func (u *User) UInt8Method() uint8 {
	return uint8(10)
}

func (u *User) UInt16Method() uint16 {
	return uint16(10)
}

func (u *User) UInt32Method() uint32 {
	return uint32(10)
}

func (u *User) UInt64Method() uint64 {
	return uint64(10)
}

func (u *User) MethodWithDifferentName() string {
	return "hello"
}

func (u *User) MethodWithContext(context map[string]interface{}) string {
	return context["version"].(string)
}

type UserCopy struct {
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

type UserCopyExtended struct {
	UserCopy
}

func NewUserCopy(now time.Time) *UserCopy {
	return &UserCopy{
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

func NewUserCopyExtended(now time.Time) *UserCopyExtended {
	return &UserCopyExtended{
		UserCopy: *NewUserCopy(now),
	}
}

// ----------------------------------------------------------------------------
// New refacto
// ----------------------------------------------------------------------------

type R1 struct{ String string }
type R2 struct{ String string }

type T1 struct {
	Int64            int64
	String           string
	StringPtr        *string
	StringPtrToValue *string
	Time             time.Time
	TimePtr          *time.Time
	TimePtrToValue   *time.Time
	SliceString      []string
	SliceInt         []int
	Map              map[string]interface{}
	NullString       null.String
	R1               R1
	R1Ptr            *R1
	R1PtrToValue     *R1
	R2               R2
	R2Ptr            *R2
	R2PtrToValue     *R2
}

func (T1) MethodString() string {
	return "method string"
}

func (T1) AnotherMethodString() string {
	return "another method string"
}

func (T1) MethodWithContext(c map[string]interface{}) map[string]interface{} {
	return c
}

type T2 struct {
	Int64             int64
	Int64Renamed      int64 `deepcopier:"field:Int64"`
	String            string
	StringPtr         *string
	StringPtrToValue  string
	Time              time.Time
	TimePtr           *time.Time
	TimePtrToValue    time.Time
	SliceString       []string
	SliceInt          []int
	Map               map[string]interface{}
	NullString        null.String
	MethodString      string
	MString           string                 `deepcopier:"field:AnotherMethodString"`
	MethodWithContext map[string]interface{} `deepcopier:"context:true"`
	R1                R1
	R1Ptr             *R1
	R1PtrToValue      R1
	R2                R2
	R2Ptr             *R2
	R2PtrToValue      R2
}

func TestCopier(t *testing.T) {
	var (
		strPtr        = "ptr"
		strPtrToValue = "ptrToValue"
		now           = time.Now()
		sliceStr      = []string{"Chuck", "Norris"}
		sliceInt      = []int{0, 8, 15}
		nullStr       = null.StringFrom("I'm null")
		mapInterfaces = map[string]interface{}{"message": "ok", "valid": true}
		methodContext = map[string]interface{}{"url": "https://ulule.com"}
	)

	r1 := &R1{String: "r1 string"}

	t1 := &T1{
		Int64:            1,
		String:           "hello",
		StringPtr:        &strPtr,
		StringPtrToValue: &strPtrToValue,
		Time:             now,
		TimePtr:          &now,
		TimePtrToValue:   &now,
		SliceString:      sliceStr,
		SliceInt:         sliceInt,
		NullString:       nullStr,
		Map:              mapInterfaces,
		R1:               *r1,
		R1Ptr:            r1,
		R1PtrToValue:     r1,
	}

	t2 := &T2{}

	options := Options{
		Context: methodContext,
	}

	assert.Nil(t, Copier(t2, t1, options))

	table := []struct {
		in  interface{}
		out interface{}
	}{
		{t1.Int64, t2.Int64},
		{t1.Int64, t2.Int64Renamed},
		{t1.String, t2.String},
		{t1.StringPtr, t2.StringPtr},
		{strPtrToValue, t2.StringPtrToValue},
		{t1.Time, t2.Time},
		{t1.TimePtr, t2.TimePtr},
		{*t1.TimePtrToValue, t2.TimePtrToValue},
		{t1.SliceString, t2.SliceString},
		{t1.SliceInt, t2.SliceInt},
		{t1.Map, t2.Map},
		{t1.NullString, t2.NullString},
		{t1.MethodString(), t2.MethodString},
		{t1.AnotherMethodString(), t2.MString},
		{t1.MethodWithContext(methodContext), t2.MethodWithContext},
		{t1.R1, t2.R1},
		{t1.R1Ptr, t2.R1Ptr},
		{*r1, t2.R1PtrToValue},
	}

	for _, tt := range table {
		assert.Equal(t, tt.in, tt.out,
			fmt.Sprintf("%v (%v) not equal to %v (%v)",
				tt.in,
				reflect.TypeOf(tt.in),
				tt.out,
				reflect.TypeOf(tt.out)))
	}
}
