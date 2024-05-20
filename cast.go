package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

func getBit(val any) int {
	bit := 0

	switch val.(type) {
	case int8, uint8, *int8, *uint8:
		bit = 8
	case int16, uint16, *int16, *uint16:
		bit = 16
	case int, uint, *int, *uint:
		bit = 0
	case int32, uint32, float32, *int32, *uint32, *float32:
		bit = 32
	case int64, uint64, float64, *int64, *uint64, *float64:
		bit = 64
	}
	return bit
}

func getNumberType(val any) numberType {
	t := numberTypeInteger
	switch val.(type) {
	case uint, uint8, uint16, uint32, uint64, *uint, *uint8, *uint16, *uint32, *uint64:
		t = numberTypeUnsignedInteger
	case float32, float64, *float32, *float64:
		t = numberTypeFloat
	}
	return t
}

func numberFromString[T any](s string) (T, error) {
	val := any(*new(T))
	isPtr := true
	switch val.(type) {
	case *int:
		val = any(int(0))
	case *int8:
		val = any(int8(0))
	case *int16:
		val = any(int16(0))
	case *int32:
		val = any(int32(0))
	case *int64:
		val = any(int64(0))
	case *uint:
		val = any(uint(0))
	case *uint8:
		val = any(uint8(0))
	case *uint16:
		val = any(uint16(0))
	case *uint32:
		val = any(uint32(0))
	case *uint64:
		val = any(uint64(0))
	case *float32:
		val = any(float32(0))
	case *float64:
		val = any(float64(0))
	default:
		isPtr = false
	}
	bit := getBit(val)
	switch getNumberType(val) {
	case numberTypeFloat:
		intVal, err := strconv.ParseFloat(s, bit)
		if isPtr {
			ptr := &intVal
			return *(*T)(unsafe.Pointer(&ptr)), err
		}
		return *(*T)(unsafe.Pointer(&intVal)), err
	case numberTypeInteger:
		intVal, err := strconv.ParseInt(s, 10, bit)
		if isPtr {
			ptr := &intVal
			return *(*T)(unsafe.Pointer(&ptr)), err
		}
		return *(*T)(unsafe.Pointer(&intVal)), err
	case numberTypeUnsignedInteger:
		intVal, err := strconv.ParseUint(s, 10, bit)
		if isPtr {
			ptr := &intVal
			return *(*T)(unsafe.Pointer(&ptr)), err
		}
		return *(*T)(unsafe.Pointer(&intVal)), err
	}
	return val.(T), nil
}

func ptrNumberFromString[T any](s string) (T, error) {
	val, err := numberFromString[T](s)
	if err != nil {
		return *new(T), err
	}
	return val, nil
}

func To[T supported](s string) (T, error) {
	val := any(*new(T))
	var err error
	switch val.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		val, err = numberFromString[T](s)
	case *int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64,
		*float32, *float64:
		val, err = ptrNumberFromString[T](s)
	case bool:
		val, err = strconv.ParseBool(s)
	case *bool:
		val, err = strconv.ParseBool(s)
		if err != nil {
			val = nil
		}
	case string:
		val = s
	case *string:
		val = &s
	default:
		return val.(T), fmt.Errorf("cast: unsupported type: %s", val)
	}
	return val.(T), err
}

var mapFunc = map[reflect.Type]func(string, any) error{
	reflect.PointerTo(reflect.TypeOf(0)):           fromToInt,
	reflect.PointerTo(reflect.TypeOf(int8(0))):     fromToInt,
	reflect.PointerTo(reflect.TypeOf(int16(0))):    fromToInt,
	reflect.PointerTo(reflect.TypeOf(int32(0))):    fromToInt,
	reflect.PointerTo(reflect.TypeOf(int64(0))):    fromToInt,
	reflect.PointerTo(reflect.TypeOf(uint(0))):     fromToUint,
	reflect.PointerTo(reflect.TypeOf(uint8(0))):    fromToUint,
	reflect.PointerTo(reflect.TypeOf(uint16(0))):   fromToUint,
	reflect.PointerTo(reflect.TypeOf(uint32(0))):   fromToUint,
	reflect.PointerTo(reflect.TypeOf(uint64(0))):   fromToUint,
	reflect.PointerTo(reflect.TypeOf(float32(0))):  fromToFloat,
	reflect.PointerTo(reflect.TypeOf(float64(0))):  fromToFloat,
	reflect.PointerTo(reflect.TypeOf(true)):        fromToFloat,
	reflect.PointerTo(reflect.TypeOf("")):          fromToString,
	reflect.PointerTo(reflect.TypeOf(time.Now())):  fromToTime,
	reflect.PointerTo(reflect.TypeOf(uuid.UUID{})): fromToUUID,
}

func fromToInt(s string, to any) error {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	switch ptr := to.(type) {
	case *int:
		*ptr = int(val)
	case *int8:
		*ptr = int8(val)
	case *int16:
		*ptr = int16(val)
	case *int32:
		*ptr = int32(val)
	case *int64:
		*ptr = val
	}
	return nil
}

func fromToUint(s string, to any) error {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	switch ptr := to.(type) {
	case *uint:
		*ptr = uint(val)
	case *uint8:
		*ptr = uint8(val)
	case *uint16:
		*ptr = uint16(val)
	case *uint32:
		*ptr = uint32(val)
	case *uint64:
		*ptr = val
	}
	return nil
}

func fromToFloat(s string, to any) error {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	switch ptr := to.(type) {
	case *float32:
		*ptr = float32(val)
	case *float64:
		*ptr = val
	}
	return nil
}

func fromToBool(s string, to any) error {
	val, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	tt := to.(*bool)
	*tt = val
	return nil
}

func fromToString(s string, to any) error {
	tt := to.(*string)
	*tt = s
	return nil
}

func fromToTime(s string, to any) error {
	val, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	tt := to.(*time.Time)
	*tt = val
	return nil
}

func fromToUUID(s string, to any) error {
	val, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	tt := to.(*uuid.UUID)
	*tt = val
	return nil
}

func FromTo(s string, to any) error {
	switch ptr := to.(type) {
	case *time.Time:
		return fromToTime(s, ptr)
	case *int, *int8, *int16, *int32, *int64:
		return fromToInt(s, ptr)
	case *uint, *uint8, *uint16, *uint32, *uint64:
		return fromToUint(s, ptr)
	case *float64, *float32:
		return fromToFloat(s, ptr)
	case *bool:
		return fromToBool(s, ptr)
	case *string:
		return fromToString(s, ptr)
	case *uuid.UUID:
		return fromToUUID(s, ptr)
	}
	return nil
}

func GenericTo[T supported](s string) (T, error) {
	var to any = new(T)
	toTypeOf := reflect.TypeOf(to)
	if toTypeOf.Kind() == reflect.Ptr && toTypeOf.Elem().Kind() == reflect.Ptr {
		to = *new(T)
	}
	err := FromTo(s, to)
	if toTypeOf.Kind() == reflect.Ptr && toTypeOf.Elem().Kind() == reflect.Ptr {
		return to.(T), err
	}
	return to.(T), err
}

func ReflectTo(s string, toType reflect.Type) (any, error) {
	pointerWrap := 0
	for toType.Kind() == reflect.Ptr {
		pointerWrap++
		toType = toType.Elem()
	}
	if f, ok := mapFunc[reflect.PointerTo(toType)]; ok {
		to := reflect.New(toType).Interface()
		err := f(s, to)
		if err != nil {
			return nil, err
		}
		val := reflect.ValueOf(to).Elem()
		for i := 0; i < pointerWrap; i++ {
			val = val.Addr()
		}
		return val.Interface(), err
	}
	return nil, fmt.Errorf("failed to convert %s to %s", s, toType.String())
}
