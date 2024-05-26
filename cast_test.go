package cast

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTo(t *testing.T) {
	assert := assert.New(t)

	// Testing integers
	valInt, err := To[int]("56")
	assert.NoError(err)
	assert.Equal(56, valInt)

	// Testing pointers to integers
	valIntPtr, err := To[*int]("42")
	assert.NoError(err)
	assert.Equal(42, *valIntPtr)

	// Testing booleans
	valBool, err := To[bool]("true")
	assert.NoError(err)
	assert.Equal(true, valBool)

	// Testing pointers to booleans
	valBoolPtr, err := To[*bool]("true")
	assert.NoError(err)
	assert.Equal(true, *valBoolPtr)

	// Testing strings
	valString, err := To[string]("hello")
	assert.NoError(err)
	assert.Equal("hello", valString)

	// Testing pointers to strings
	valStringPtr, err := To[*string]("world")
	assert.NoError(err)
	assert.Equal("world", *valStringPtr)

	// Testing integers of different bits
	valInt8, err := To[int8]("12")
	assert.NoError(err)
	assert.Equal(int8(12), valInt8)

	valInt16, err := To[int16]("1234")
	assert.NoError(err)
	assert.Equal(int16(1234), valInt16)

	valInt32, err := To[int32]("12345")
	assert.NoError(err)
	assert.Equal(int32(12345), valInt32)

	valInt64, err := To[int64]("123456789")
	assert.NoError(err)
	assert.Equal(int64(123456789), valInt64)

	// Testing pointers to integers of different bits
	valIntPtr8, err := To[*int8]("42")
	assert.NoError(err)
	assert.Equal(int8(42), *valIntPtr8)

	valIntPtr16, err := To[*int16]("1234")
	assert.NoError(err)
	assert.Equal(int16(1234), *valIntPtr16)

	valIntPtr32, err := To[*int32]("12345")
	assert.NoError(err)
	assert.Equal(int32(12345), *valIntPtr32)

	valIntPtr64, err := To[*int64]("123456789")
	assert.NoError(err)
	assert.Equal(int64(123456789), *valIntPtr64)

	// Testing unsigned integers of different bits
	valUint8, err := To[uint8]("12")
	assert.NoError(err)
	assert.Equal(uint8(12), valUint8)

	valUint16, err := To[uint16]("1234")
	assert.NoError(err)
	assert.Equal(uint16(1234), valUint16)

	valUint32, err := To[uint32]("12345")
	assert.NoError(err)
	assert.Equal(uint32(12345), valUint32)

	valUint64, err := To[uint64]("123456789")
	assert.NoError(err)
	assert.Equal(uint64(123456789), valUint64)

	// Testing pointers to unsigned integers of different bits
	valUintPtr8, err := To[*uint8]("42")
	assert.NoError(err)
	assert.Equal(uint8(42), *valUintPtr8)

	valUintPtr16, err := To[*uint16]("1234")
	assert.NoError(err)
	assert.Equal(uint16(1234), *valUintPtr16)

	valUintPtr32, err := To[*uint32]("12345")
	assert.NoError(err)
	assert.Equal(uint32(12345), *valUintPtr32)

	valUintPtr64, err := To[*uint64]("123456789")
	assert.NoError(err)
	assert.Equal(uint64(123456789), *valUintPtr64)

	// Testing floats of different bits
	valFloat32, err := To[float32]("123.45")
	assert.NoError(err)
	assert.Equal(float32(123.45), valFloat32)

	valFloat64, err := To[float64]("123.456789")
	assert.NoError(err)
	assert.Equal(float64(123.456789), valFloat64)

	// Testing pointers to integers of floats
	valFloatPtr32, err := To[*float32]("42.42")
	assert.NoError(err)
	assert.Equal(float32(42.42), *valFloatPtr32)

	valFloatPtr64, err := To[*float64]("123.456789")
	assert.NoError(err)
	assert.Equal(float64(123.456789), *valFloatPtr64)

	date := time.Date(2024, time.May, 22, 11, 36, 57, 0, time.UTC)
	valTime, err := To[time.Time](date.Format(time.RFC3339))
	assert.NoError(err)
	assert.Equal(date, valTime)

	valPtrTime, err := To[*time.Time](date.Format(time.RFC3339))
	assert.NoError(err)
	assert.Equal(date, *valPtrTime)

	id := uuid.New()
	valId, err := To[uuid.UUID](id.String())
	assert.NoError(err)
	assert.Equal(id, valId)

	valPtrId, err := To[*uuid.UUID](id.String())
	assert.NoError(err)
	assert.Equal(id, *valPtrId)
}

func TestFromTo(t *testing.T) {
	assert := assert.New(t)

	// Define the variable for ToFrom function
	var (
		valInt     int
		valBool    bool
		valString  string
		valInt8    int8
		valInt16   int16
		valInt32   int32
		valInt64   int64
		valUint8   uint8
		valUint16  uint16
		valUint32  uint32
		valUint64  uint64
		valFloat32 float32
		valFloat64 float64
		valTime    time.Time
		valUUID    uuid.UUID
	)

	// Assume that ToFrom method accepts two arguments: a string and a pointer to set the converted value.
	// Testing integers
	assert.NoError(ToFrom("56", &valInt))
	assert.Equal(56, valInt)

	// Testing booleans
	assert.NoError(ToFrom("true", &valBool))
	assert.Equal(true, valBool)

	// Testing strings
	assert.NoError(ToFrom("hello", &valString))
	assert.Equal("hello", valString)

	// Testing integers of different bits
	assert.NoError(ToFrom("12", &valInt8))
	assert.Equal(int8(12), valInt8)

	assert.NoError(ToFrom("1234", &valInt16))
	assert.Equal(int16(1234), valInt16)

	assert.NoError(ToFrom("12345", &valInt32))
	assert.Equal(int32(12345), valInt32)

	assert.NoError(ToFrom("123456789", &valInt64))
	assert.Equal(int64(123456789), valInt64)

	// Testing unsigned integers of different bits
	assert.NoError(ToFrom("12", &valUint8))
	assert.Equal(uint8(12), valUint8)

	assert.NoError(ToFrom("1234", &valUint16))
	assert.Equal(uint16(1234), valUint16)

	assert.NoError(ToFrom("12345", &valUint32))
	assert.Equal(uint32(12345), valUint32)

	assert.NoError(ToFrom("123456789", &valUint64))
	assert.Equal(uint64(123456789), valUint64)

	// Testing floats of different bits
	assert.NoError(ToFrom("123.45", &valFloat32))
	assert.Equal(float32(123.45), valFloat32)

	assert.NoError(ToFrom("123.456789", &valFloat64))
	assert.Equal(float64(123.456789), valFloat64)

	date := time.Date(2024, time.May, 22, 11, 36, 57, 0, time.UTC)
	assert.NoError(ToFrom(date.Format(time.RFC3339), &valTime))
	assert.Equal(date, valTime)

	id := uuid.New()
	assert.NoError(ToFrom(id.String(), &valUUID))
	assert.Equal(id, valUUID)
}

func TestReflectTo(t *testing.T) {
	assert := assert.New(t)

	// Define the variables for ToReflect function
	var (
		valInt        int
		valIntPtr     *int
		valBool       bool
		valBoolPtr    *bool
		valString     string
		valStringPtr  *string
		valInt8       int8
		valInt16      int16
		valInt32      int32
		valInt64      int64
		valIntPtr8    *int8
		valIntPtr16   *int16
		valIntPtr32   *int32
		valIntPtr64   *int64
		valUint8      uint8
		valUint16     uint16
		valUint32     uint32
		valUint64     uint64
		valUintPtr8   *uint8
		valUintPtr16  *uint16
		valUintPtr32  *uint32
		valUintPtr64  *uint64
		valFloat32    float32
		valFloat64    float64
		valFloatPtr32 *float32
		valFloatPtr64 *float64
		valTime       time.Time
		valUUID       uuid.UUID
		valTimePtr    *time.Time
		valUUIDPtr    *uuid.UUID
	)

	// Testing integers
	res, err := ToReflect("56", reflect.TypeOf(valInt))
	assert.NoError(err)
	assert.Equal(56, res.(int))

	// Continue testing other types
	res, err = ToReflect("true", reflect.TypeOf(valBool))
	assert.NoError(err)
	assert.Equal(true, res.(bool))

	res, err = ToReflect("hello", reflect.TypeOf(valString))
	assert.NoError(err)
	assert.Equal("hello", res.(string))

	// And some more tests
	res, err = ToReflect("1", reflect.TypeOf(valInt8))
	assert.NoError(err)
	assert.Equal(int8(1), res.(int8))

	res, err = ToReflect("2", reflect.TypeOf(valInt16))
	assert.NoError(err)
	assert.Equal(int16(2), res.(int16))

	res, err = ToReflect("3", reflect.TypeOf(valInt32))
	assert.NoError(err)
	assert.Equal(int32(3), res.(int32))

	res, err = ToReflect("4", reflect.TypeOf(valInt64))
	assert.NoError(err)
	assert.Equal(int64(4), res.(int64))

	res, err = ToReflect("5", reflect.TypeOf(valUint8))
	assert.NoError(err)
	assert.Equal(uint8(5), res.(uint8))

	res, err = ToReflect("6", reflect.TypeOf(valUint16))
	assert.NoError(err)
	assert.Equal(uint16(6), res.(uint16))

	res, err = ToReflect("7", reflect.TypeOf(valUint32))
	assert.NoError(err)
	assert.Equal(uint32(7), res.(uint32))

	res, err = ToReflect("8", reflect.TypeOf(valUint64))
	assert.NoError(err)
	assert.Equal(uint64(8), res.(uint64))

	res, err = ToReflect("0.9", reflect.TypeOf(valFloat32))
	assert.NoError(err)
	assert.Equal(float32(0.9), res.(float32))

	res, err = ToReflect("1.0", reflect.TypeOf(valFloat64))
	assert.NoError(err)
	assert.Equal(float64(1.0), res.(float64))

	res, err = ToReflect("56", reflect.TypeOf(valIntPtr))
	assert.NoError(err)
	assert.Equal(56, *res.(*int))

	res, err = ToReflect("true", reflect.TypeOf(valBoolPtr))
	assert.NoError(err)
	assert.Equal(true, *res.(*bool))

	res, err = ToReflect("hello", reflect.TypeOf(valStringPtr))
	assert.NoError(err)
	assert.Equal("hello", *res.(*string))

	res, err = ToReflect("1", reflect.TypeOf(valIntPtr8))
	assert.NoError(err)
	assert.Equal(int8(1), *res.(*int8))

	res, err = ToReflect("2", reflect.TypeOf(valIntPtr16))
	assert.NoError(err)
	assert.Equal(int16(2), *res.(*int16))

	res, err = ToReflect("3", reflect.TypeOf(valIntPtr32))
	assert.NoError(err)
	assert.Equal(int32(3), *res.(*int32))

	res, err = ToReflect("4", reflect.TypeOf(valIntPtr64))
	assert.NoError(err)
	assert.Equal(int64(4), *res.(*int64))

	res, err = ToReflect("5", reflect.TypeOf(valUintPtr8))
	assert.NoError(err)
	assert.Equal(uint8(5), *res.(*uint8))

	res, err = ToReflect("6", reflect.TypeOf(valUintPtr16))
	assert.NoError(err)
	assert.Equal(uint16(6), *res.(*uint16))

	res, err = ToReflect("7", reflect.TypeOf(valUintPtr32))
	assert.NoError(err)
	assert.Equal(uint32(7), *res.(*uint32))

	res, err = ToReflect("8", reflect.TypeOf(valUintPtr64))
	assert.NoError(err)
	assert.Equal(uint64(8), *res.(*uint64))

	res, err = ToReflect("0.9", reflect.TypeOf(valFloatPtr32))
	assert.NoError(err)
	assert.Equal(float32(0.9), *res.(*float32))

	res, err = ToReflect("1.0", reflect.TypeOf(valFloatPtr64))
	assert.NoError(err)
	assert.Equal(float64(1.0), *res.(*float64))

	date := time.Date(2024, time.May, 22, 11, 36, 57, 0, time.UTC)
	res, err = ToReflect(date.Format(time.RFC3339), reflect.TypeOf(valTime))
	assert.NoError(err)
	assert.Equal(date, res.(time.Time))

	res, err = ToReflect(date.Format(time.RFC3339), reflect.TypeOf(valTimePtr))
	assert.NoError(err)
	assert.Equal(date, *res.(*time.Time))

	id := uuid.New()
	res, err = ToReflect(id.String(), reflect.TypeOf(valUUID))
	assert.NoError(err)
	assert.Equal(id, res.(uuid.UUID))

	res, err = ToReflect(id.String(), reflect.TypeOf(valUUIDPtr))
	assert.NoError(err)
	assert.Equal(id, *res.(*uuid.UUID))
}
