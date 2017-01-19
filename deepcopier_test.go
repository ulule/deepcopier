package deepcopier

import (
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func ATestCopyTo(t *testing.T) {
	is := assert.New(t)
	now := time.Now()
	pointer := "pointer"

	user := NewUser(now, &pointer)
	userCopy := &UserCopy{}
	userCopyExtended := &UserCopyExtended{}
	expectedUserCopy := NewUserCopy(now, &pointer)
	expectedUserCopyExtended := NewUserCopyExtended(now, &pointer)

	is.Nil(Copy(user).WithContext(map[string]interface{}{"version": "1"}).To(userCopy))

	is.Equal(expectedUserCopy.Title, userCopy.Title)
	is.Equal(expectedUserCopy.Date, userCopy.Date)
	is.Equal(expectedUserCopy.Float32, userCopy.Float32)
	is.Equal(expectedUserCopy.Float64, userCopy.Float64)
	is.Equal(expectedUserCopy.Int, userCopy.Int)
	is.Equal(expectedUserCopy.Int8, userCopy.Int8)
	is.Equal(expectedUserCopy.Int16, userCopy.Int16)
	is.Equal(expectedUserCopy.Int32, userCopy.Int32)
	is.Equal(expectedUserCopy.Int64, userCopy.Int64)
	is.Equal(expectedUserCopy.UInt, userCopy.UInt)
	is.Equal(expectedUserCopy.UInt8, userCopy.UInt8)
	is.Equal(expectedUserCopy.UInt16, userCopy.UInt16)
	is.Equal(expectedUserCopy.UInt32, userCopy.UInt32)
	is.Equal(expectedUserCopy.UInt64, userCopy.UInt64)
	is.Equal(expectedUserCopy.StringSlice, userCopy.StringSlice)
	is.Equal(expectedUserCopy.IntSlice, userCopy.IntSlice)
	is.Equal(expectedUserCopy.IntMethod, userCopy.IntMethod)
	is.Equal(expectedUserCopy.Int8Method, userCopy.Int8Method)
	is.Equal(expectedUserCopy.Int16Method, userCopy.Int16Method)
	is.Equal(expectedUserCopy.Int32Method, userCopy.Int32Method)
	is.Equal(expectedUserCopy.Int64Method, userCopy.Int64Method)
	is.Equal(expectedUserCopy.UIntMethod, userCopy.UIntMethod)
	is.Equal(expectedUserCopy.UInt8Method, userCopy.UInt8Method)
	is.Equal(expectedUserCopy.UInt16Method, userCopy.UInt16Method)
	is.Equal(expectedUserCopy.UInt32Method, userCopy.UInt32Method)
	is.Equal(expectedUserCopy.UInt64Method, userCopy.UInt64Method)
	is.Equal(expectedUserCopy.MethodWithContext, userCopy.MethodWithContext)
	is.Equal(expectedUserCopy.SuperMethod, userCopy.SuperMethod)
	is.Equal(expectedUserCopy.StringSlice, userCopy.StringSlice)
	is.Equal(expectedUserCopy.IntSlice, userCopy.IntSlice)

	is.Nil(Copy(user).WithContext(map[string]interface{}{"version": "1"}).To(userCopyExtended))

	is.Equal(expectedUserCopyExtended.Title, userCopyExtended.Title)
	is.Equal(expectedUserCopyExtended.Date, userCopyExtended.Date)
	is.Equal(expectedUserCopyExtended.Float32, userCopyExtended.Float32)
	is.Equal(expectedUserCopyExtended.Float64, userCopyExtended.Float64)
	is.Equal(expectedUserCopyExtended.Int, userCopyExtended.Int)
	is.Equal(expectedUserCopyExtended.Int8, userCopyExtended.Int8)
	is.Equal(expectedUserCopyExtended.Int16, userCopyExtended.Int16)
	is.Equal(expectedUserCopyExtended.Int32, userCopyExtended.Int32)
	is.Equal(expectedUserCopyExtended.Int64, userCopyExtended.Int64)
	is.Equal(expectedUserCopyExtended.UInt, userCopyExtended.UInt)
	is.Equal(expectedUserCopyExtended.UInt8, userCopyExtended.UInt8)
	is.Equal(expectedUserCopyExtended.UInt16, userCopyExtended.UInt16)
	is.Equal(expectedUserCopyExtended.UInt32, userCopyExtended.UInt32)
	is.Equal(expectedUserCopyExtended.UInt64, userCopyExtended.UInt64)
	is.Equal(expectedUserCopyExtended.StringSlice, userCopyExtended.StringSlice)
	is.Equal(expectedUserCopyExtended.IntSlice, userCopyExtended.IntSlice)
	is.Equal(expectedUserCopyExtended.IntMethod, userCopyExtended.IntMethod)
	is.Equal(expectedUserCopyExtended.Int8Method, userCopyExtended.Int8Method)
	is.Equal(expectedUserCopyExtended.Int16Method, userCopyExtended.Int16Method)
	is.Equal(expectedUserCopyExtended.Int32Method, userCopyExtended.Int32Method)
	is.Equal(expectedUserCopyExtended.Int64Method, userCopyExtended.Int64Method)
	is.Equal(expectedUserCopyExtended.UIntMethod, userCopyExtended.UIntMethod)
	is.Equal(expectedUserCopyExtended.UInt8Method, userCopyExtended.UInt8Method)
	is.Equal(expectedUserCopyExtended.UInt16Method, userCopyExtended.UInt16Method)
	is.Equal(expectedUserCopyExtended.UInt32Method, userCopyExtended.UInt32Method)
	is.Equal(expectedUserCopyExtended.UInt64Method, userCopyExtended.UInt64Method)
	is.Equal(expectedUserCopyExtended.MethodWithContext, userCopyExtended.MethodWithContext)
	is.Equal(expectedUserCopyExtended.SuperMethod, userCopyExtended.SuperMethod)
	is.Equal(expectedUserCopyExtended.StringSlice, userCopyExtended.StringSlice)
	is.Equal(expectedUserCopyExtended.IntSlice, userCopyExtended.IntSlice)
}

func TestCopyFrom(t *testing.T) {
	is := assert.New(t)
	now := time.Now()
	pointer := "pointer"

	user := &User{}
	userExpected := NewUser(now, &pointer)
	userCopy := NewUserCopy(now, &pointer)
	userCopyExtended := NewUserCopyExtended(now, &pointer)

	is.Nil(Copy(user).From(userCopy))

	is.Equal(userExpected.Name, user.Name)
	is.Equal(userExpected.Date, user.Date)
	is.Equal(userExpected.AFloat32, user.AFloat32)
	is.Equal(userExpected.AFloat64, user.AFloat64)
	is.Equal(userExpected.AnInt, user.AnInt)
	is.Equal(userExpected.AnInt8, user.AnInt8)
	is.Equal(userExpected.AnInt16, user.AnInt16)
	is.Equal(userExpected.AnInt32, user.AnInt32)
	is.Equal(userExpected.AnInt64, user.AnInt64)
	is.Equal(userExpected.AnUInt, user.AnUInt)
	is.Equal(userExpected.AnUInt8, user.AnUInt8)
	is.Equal(userExpected.AnUInt16, user.AnUInt16)
	is.Equal(userExpected.AnUInt32, user.AnUInt32)
	is.Equal(userExpected.AnUInt64, user.AnUInt64)
	is.Equal(userExpected.AStringSlice, user.AStringSlice)
	is.Equal(userExpected.AnIntSlice, user.AnIntSlice)

	is.Nil(Copy(user).From(userCopyExtended))

	is.Equal(userExpected.Name, user.Name)
	is.Equal(userExpected.Date, user.Date)
	is.Equal(userExpected.AFloat32, user.AFloat32)
	is.Equal(userExpected.AFloat64, user.AFloat64)
	is.Equal(userExpected.AnInt, user.AnInt)
	is.Equal(userExpected.AnInt8, user.AnInt8)
	is.Equal(userExpected.AnInt16, user.AnInt16)
	is.Equal(userExpected.AnInt32, user.AnInt32)
	is.Equal(userExpected.AnInt64, user.AnInt64)
	is.Equal(userExpected.AnUInt, user.AnUInt)
	is.Equal(userExpected.AnUInt8, user.AnUInt8)
	is.Equal(userExpected.AnUInt16, user.AnUInt16)
	is.Equal(userExpected.AnUInt32, user.AnUInt32)
	is.Equal(userExpected.AnUInt64, user.AnUInt64)
	is.Equal(userExpected.AStringSlice, user.AStringSlice)
	is.Equal(userExpected.AnIntSlice, user.AnIntSlice)
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

func NewUser(now time.Time, pointer *string) *User {
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
		APointer:     *pointer,
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
	PointerString     *string     `json:"a_pointer_string" deepcopier:"field:APointer"`
}

type UserCopyExtended struct {
	UserCopy
}

func NewUserCopy(now time.Time, pointer *string) *UserCopy {
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
		PointerString:     pointer,
	}
}

func NewUserCopyExtended(now time.Time, pointer *string) *UserCopyExtended {
	return &UserCopyExtended{
		UserCopy: *NewUserCopy(now, pointer),
	}
}
