package deepcopier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	type (
		Rel struct {
			Int int
		}

		Src struct {
			Int       int
			IntPtr    *int
			Slice     []string
			SlicePtr  *[]string
			Map       map[string]interface{}
			MapPtr    *map[string]interface{}
			Struct    Rel
			StructPtr *Rel
			Skipped   string `deepcopier:"skip"`
		}

		Dst struct {
			Int       int
			IntPtr    *int
			Slice     []string
			SlicePtr  *[]string
			Map       map[string]interface{}
			MapPtr    *map[string]interface{}
			Struct    Rel
			StructPtr *Rel
			Skipped   string `deepcopier:"skip"`
		}

		Renamed struct {
			MyInt       int                     `deepcopier:"field:Int"`
			MyIntPtr    *int                    `deepcopier:"field:IntPtr"`
			MySlice     []string                `deepcopier:"field:Slice"`
			MySlicePtr  *[]string               `deepcopier:"field:SlicePtr"`
			MyMap       map[string]interface{}  `deepcopier:"field:Map"`
			MyMapPtr    *map[string]interface{} `deepcopier:"field:MapPtr"`
			MyStruct    Rel                     `deepcopier:"field:Struct"`
			MyStructPtr *Rel                    `deepcopier:"field:StructPtr"`
			Skipped     string                  `deepcopier:"skip"`
		}
	)

	var (
		integer = 1
		rel     = Rel{Int: 1}
		slc     = []string{"one", "two"}
		mp      = map[string]interface{}{"one": 1}
	)

	src := &Src{
		Int:       integer,
		IntPtr:    &integer,
		Slice:     slc,
		SlicePtr:  &slc,
		Map:       mp,
		MapPtr:    &mp,
		Struct:    rel,
		StructPtr: &rel,
		Skipped:   "I should be skipped",
	}

	srcRenamed := &Renamed{
		MyInt:       integer,
		MyIntPtr:    &integer,
		MySlice:     slc,
		MySlicePtr:  &slc,
		MyMap:       mp,
		MyMapPtr:    &mp,
		MyStruct:    rel,
		MyStructPtr: &rel,
		Skipped:     "I should be skipped",
	}

	//
	// To()
	//

	dst := &Dst{}
	assert.Nil(t, Copy(src).To(dst))
	assert.Equal(t, src.Int, dst.Int)
	assert.Equal(t, src.IntPtr, dst.IntPtr)
	assert.Equal(t, src.Slice, dst.Slice)
	assert.Equal(t, src.SlicePtr, dst.SlicePtr)
	assert.Equal(t, src.Map, dst.Map)
	assert.Equal(t, src.MapPtr, dst.MapPtr)
	assert.Equal(t, src.Struct, dst.Struct)
	assert.Equal(t, src.StructPtr, dst.StructPtr)
	assert.Zero(t, dst.Skipped)

	dstRenamed := &Renamed{}
	assert.Nil(t, Copy(src).To(dstRenamed))
	assert.Equal(t, src.Int, dstRenamed.MyInt)
	assert.Equal(t, src.IntPtr, dstRenamed.MyIntPtr)
	assert.Equal(t, src.Slice, dstRenamed.MySlice)
	assert.Equal(t, src.SlicePtr, dstRenamed.MySlicePtr)
	assert.Equal(t, src.Map, dstRenamed.MyMap)
	assert.Equal(t, src.MapPtr, dstRenamed.MyMapPtr)
	assert.Equal(t, src.Struct, dstRenamed.MyStruct)
	assert.Equal(t, src.StructPtr, dstRenamed.MyStructPtr)
	assert.Zero(t, dstRenamed.Skipped)

	//
	// From()
	//

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(src))
	assert.Equal(t, src.Int, dst.Int)
	assert.Equal(t, src.IntPtr, dst.IntPtr)
	assert.Equal(t, src.Slice, dst.Slice)
	assert.Equal(t, src.SlicePtr, dst.SlicePtr)
	assert.Equal(t, src.Map, dst.Map)
	assert.Equal(t, src.MapPtr, dst.MapPtr)
	assert.Equal(t, src.Struct, dst.Struct)
	assert.Equal(t, src.StructPtr, dst.StructPtr)
	assert.Zero(t, dst.Skipped)

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(srcRenamed))
	assert.Equal(t, srcRenamed.MyInt, dst.Int)
	assert.Equal(t, srcRenamed.MyIntPtr, dst.IntPtr)
	assert.Equal(t, srcRenamed.MySlice, dst.Slice)
	assert.Equal(t, srcRenamed.MySlicePtr, dst.SlicePtr)
	assert.Equal(t, srcRenamed.MyMap, dst.Map)
	assert.Equal(t, srcRenamed.MyMapPtr, dst.MapPtr)
	assert.Equal(t, srcRenamed.MyStruct, dst.Struct)
	assert.Equal(t, srcRenamed.MyStructPtr, dst.StructPtr)
	assert.Zero(t, dst.Skipped)
}

func TestField_PointerToValue(t *testing.T) {
	type (
		Rel struct {
			Int int
		}

		Src struct {
			Int    *int
			Slice  *[]string
			Map    *map[string]interface{}
			Struct *Rel
		}

		Dst struct {
			Int    int
			Slice  []string
			Map    map[string]interface{}
			Struct Rel
		}

		SrcRenamed struct {
			MyInt    *int                    `deepcopier:"field:Int"`
			MySlice  *[]string               `deepcopier:"field:Slice"`
			MyMap    *map[string]interface{} `deepcopier:"field:Map"`
			MyStruct *Rel                    `deepcopier:"field:Struct"`
		}

		DstRenamed struct {
			MyInt    int                    `deepcopier:"field:Int"`
			MySlice  []string               `deepcopier:"field:Slice"`
			MyMap    map[string]interface{} `deepcopier:"field:Map"`
			MyStruct Rel                    `deepcopier:"field:Struct"`
		}
	)

	var (
		rel     = Rel{Int: 1}
		integer = 1
		slc     = []string{"one", "two"}
		mp      = map[string]interface{}{"one": 1}
	)

	src := &Src{
		Int:    &integer,
		Slice:  &slc,
		Map:    &mp,
		Struct: &rel,
	}

	srcRenamed := &SrcRenamed{
		MyInt:    &integer,
		MySlice:  &slc,
		MyMap:    &mp,
		MyStruct: &rel,
	}

	//
	// To()
	//

	dst := &Dst{}
	assert.Nil(t, Copy(src).To(dst))
	assert.Equal(t, *src.Int, dst.Int)
	assert.Equal(t, *src.Slice, dst.Slice)
	assert.Equal(t, *src.Map, dst.Map)
	assert.Equal(t, *src.Struct, dst.Struct)

	dstRenamed := &DstRenamed{}
	assert.Nil(t, Copy(src).To(dstRenamed))
	assert.Equal(t, *src.Int, dstRenamed.MyInt)
	assert.Equal(t, *src.Slice, dstRenamed.MySlice)
	assert.Equal(t, *src.Map, dstRenamed.MyMap)
	assert.Equal(t, *src.Struct, dstRenamed.MyStruct)

	//
	// From()
	//

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(src))
	assert.Equal(t, *src.Int, dst.Int)
	assert.Equal(t, *src.Slice, dst.Slice)
	assert.Equal(t, *src.Map, dst.Map)
	assert.Equal(t, *src.Struct, dst.Struct)

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(srcRenamed))
	assert.Equal(t, *srcRenamed.MyInt, dst.Int)
	assert.Equal(t, *srcRenamed.MySlice, dst.Slice)
	assert.Equal(t, *srcRenamed.MyMap, dst.Map)
	assert.Equal(t, *srcRenamed.MyStruct, dst.Struct)
}

func TestField_Unexported(t *testing.T) {
	type (
		Src struct {
			Exported   int
			unexported string
		}

		Dst struct {
			Exported   int
			unexported string
		}
	)

	src := &Src{Exported: 1, unexported: "unexported"}

	//
	// To()
	//

	dst := &Dst{}
	assert.Nil(t, Copy(src).To(dst))
	assert.Equal(t, "", dst.unexported)

	//
	// From()
	//

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(src))
	assert.Equal(t, "", dst.unexported)
}

func TestField_Unknown(t *testing.T) {
	type (
		Original struct {
			Int int
		}

		Renamed struct {
			MyInt int `deepcopier:"field:Integer"`
		}
	)

	//
	// To()
	//

	src := &Original{Int: 1}
	dstRenamed := &Renamed{}
	assert.Nil(t, Copy(src).To(dstRenamed))
	assert.Equal(t, 0, dstRenamed.MyInt)

	//
	// From()
	//

	srcRenamed := &Renamed{MyInt: 1}
	dst := &Original{}
	assert.Nil(t, Copy(dst).From(srcRenamed))
	assert.Equal(t, 0, dst.Int)
}

func TestField_EmptyInterface(t *testing.T) {
	type (
		Rel struct {
			Int int
		}

		Src struct {
			Rel *Rel
		}

		SrcForce struct {
			Rel *Rel `deepcopier:"force"`
		}

		Dst struct {
			Rel interface{}
		}

		DstForce struct {
			Rel interface{} `deepcopier:"force"`
		}
	)

	var (
		rel      = &Rel{Int: 1}
		src      = &Src{Rel: rel}
		srcForce = &SrcForce{Rel: rel}
	)

	//
	// Without force
	//

	dst := &Dst{}
	assert.Nil(t, Copy(src).To(dst))
	assert.Nil(t, dst.Rel)

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(src))
	assert.Nil(t, dst.Rel)

	//
	// With force
	//

	dstForce := &DstForce{}
	assert.Nil(t, Copy(src).To(dstForce))
	assert.Equal(t, src.Rel, dstForce.Rel)

	dstForce = &DstForce{}
	assert.Nil(t, Copy(dstForce).From(srcForce))
	assert.Equal(t, srcForce.Rel, dstForce.Rel)
}

func TestField_SameNameWithDifferentType(t *testing.T) {
	type (
		FooInt struct {
			Foo int
		}

		FooStr struct {
			Foo string
		}
	)

	//
	// To()
	//

	srcInt := &FooInt{Foo: 1}
	dstStr := &FooStr{}

	assert.Nil(t, Copy(dstStr).From(srcInt))
	assert.Empty(t, dstStr.Foo)

	//
	// From()
	//

	dstStr = &FooStr{}
	assert.Nil(t, Copy(dstStr).From(srcInt))
	assert.Empty(t, dstStr.Foo)
}

func TestMethod(t *testing.T) {
	var (
		c   = map[string]interface{}{"message": "hello"}
		src = &MethodTesterFoo{}
		dst = &MethodTesterBar{}
	)

	//
	// To()
	//

	assert.Nil(t, Copy(src).WithContext(c).To(dst))
	assert.Equal(t, c, dst.FooContext)
	assert.Equal(t, MethodTesterFoo{}.FooInteger(), dst.FooInteger)
	assert.Empty(t, dst.FooSkipped)

	//
	// From()
	//

	dst = &MethodTesterBar{}
	assert.Nil(t, Copy(dst).WithContext(c).From(src))
	assert.Equal(t, c, dst.FooContext)
	assert.Equal(t, MethodTesterFoo{}.FooInteger(), dst.FooInteger)
	assert.Empty(t, dst.FooSkipped)
}

func TestAnonymousStruct(t *testing.T) {
	type (
		Embedded             struct{ Int int }
		EmbeddedRenamedField struct {
			MyInt int `deepcopier:"field:Int"`
		}

		Src             struct{ Embedded }
		SrcRenamedField struct{ EmbeddedRenamedField }

		Dst             struct{ Int int }
		DstRenamedField struct {
			MyInt int `deepcopier:"field:Int"`
		}
	)

	var (
		embedded             = Embedded{Int: 1}
		embeddedRenamedField = EmbeddedRenamedField{MyInt: 1}
		src                  = &Src{Embedded: embedded}
		srcRenamedField      = &SrcRenamedField{EmbeddedRenamedField: embeddedRenamedField}
	)

	//
	// To()
	//

	dst := &Dst{}
	assert.Nil(t, Copy(src).To(dst))
	assert.Equal(t, src.Int, dst.Int)

	dstRenamedField := &DstRenamedField{}
	assert.Nil(t, Copy(src).To(dstRenamedField))
	assert.Equal(t, src.Int, dstRenamedField.MyInt)

	//
	// From()
	//

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(src))
	assert.Equal(t, src.Int, dst.Int)

	dst = &Dst{}
	assert.Nil(t, Copy(dst).From(srcRenamedField))
	assert.Equal(t, srcRenamedField.MyInt, dst.Int)
}

// ----------------------------------------------------------------------------
// Method testers
// ----------------------------------------------------------------------------

type MethodTesterFoo struct {
	BarInteger int
	BarContext map[string]interface{} `deepcopier:"context"`
	BarSkipped string                 `deepcopier:"skip"`
}

func (MethodTesterFoo) FooInteger() int {
	return 1
}

func (MethodTesterFoo) FooContext(c map[string]interface{}) map[string]interface{} {
	return c
}

func (MethodTesterFoo) FooSkipped() string {
	return "skipped"
}

type MethodTesterBar struct {
	FooInteger int
	FooContext map[string]interface{} `deepcopier:"context"`
	FooSkipped string                 `deepcopier:"skip"`
}

func (MethodTesterBar) BarInteger() int {
	return 1
}

func (MethodTesterBar) BarContext(c map[string]interface{}) map[string]interface{} {
	return c
}

func (MethodTesterBar) BarSkipped() string {
	return "skipped"
}
