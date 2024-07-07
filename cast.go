package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	typeOfPointerInt     = reflect.PointerTo(reflect.TypeOf(0))
	typeOfPointerInt8    = reflect.PointerTo(reflect.TypeOf(int8(0)))
	typeOfPointerInt16   = reflect.PointerTo(reflect.TypeOf(int16(0)))
	typeOfPointerInt32   = reflect.PointerTo(reflect.TypeOf(int32(0)))
	typeOfPointerInt64   = reflect.PointerTo(reflect.TypeOf(int64(0)))
	typeOfPointerUint    = reflect.PointerTo(reflect.TypeOf(uint(0)))
	typeOfPointerUint8   = reflect.PointerTo(reflect.TypeOf(uint8(0)))
	typeOfPointerUint16  = reflect.PointerTo(reflect.TypeOf(uint16(0)))
	typeOfPointerUint32  = reflect.PointerTo(reflect.TypeOf(uint32(0)))
	typeOfPointerUint64  = reflect.PointerTo(reflect.TypeOf(uint64(0)))
	typeOfPointerFloat32 = reflect.PointerTo(reflect.TypeOf(float32(0)))
	typeOfPointerFloat64 = reflect.PointerTo(reflect.TypeOf(float64(0)))
	typeOfPointerBool    = reflect.PointerTo(reflect.TypeOf(false))
	typeOfPointerString  = reflect.PointerTo(reflect.TypeOf(""))
	typeOfPointerTime    = reflect.PointerTo(reflect.TypeOf(time.Time{}))
	typeOfPointerUUID    = reflect.PointerTo(reflect.TypeOf(uuid.UUID{}))
)

var mapFunc = map[reflect.Type]func(string, any) error{
	typeOfPointerInt:     fromToInt,
	typeOfPointerInt8:    fromToInt,
	typeOfPointerInt16:   fromToInt,
	typeOfPointerInt32:   fromToInt,
	typeOfPointerInt64:   fromToInt,
	typeOfPointerUint:    fromToUint,
	typeOfPointerUint8:   fromToUint,
	typeOfPointerUint16:  fromToUint,
	typeOfPointerUint32:  fromToUint,
	typeOfPointerUint64:  fromToUint,
	typeOfPointerFloat32: fromToFloat,
	typeOfPointerFloat64: fromToFloat,
	typeOfPointerBool:    fromToBool,
	typeOfPointerString:  fromToString,
	typeOfPointerTime:    fromToTime,
	typeOfPointerUUID:    fromToUUID,
}

func fromToInt(s string, to any) error {
	switch ptr := to.(type) {
	case *int:
		val, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return err
		}
		*ptr = int(val)
	case *int8:
		val, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return err
		}
		*ptr = int8(val)
	case *int16:
		val, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return err
		}
		*ptr = int16(val)
	case *int32:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		*ptr = int32(val)
	case *int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*ptr = val
	}
	return nil
}

func fromToUint(s string, to any) error {
	switch ptr := to.(type) {
	case *uint:
		val, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return err
		}
		*ptr = uint(val)
	case *uint8:
		val, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return err
		}
		*ptr = uint8(val)
	case *uint16:
		val, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return err
		}
		*ptr = uint16(val)
	case *uint32:
		val, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return err
		}
		*ptr = uint32(val)
	case *uint64:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*ptr = val
	}
	return nil
}

func fromToFloat(s string, to any) error {
	switch ptr := to.(type) {
	case *float32:
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		*ptr = float32(val)
	case *float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
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

func ToFrom(s string, to any) error {
	typeOfPtr := reflect.TypeOf(to)
	typeOfVal := typeOfPtr.Elem()
	val, err := ToReflect(s, typeOfVal)
	if err != nil {
		return err
	}
	reflect.ValueOf(to).Elem().Set(reflect.ValueOf(val))
	return nil
}

func To[T supported](s string) (T, error) {
	to := *new(T)
	toTypeOf := reflect.TypeOf(to)
	val, err := ToReflect(s, toTypeOf)
	if err != nil {
		return to, err
	}
	return val.(T), nil
}

func ToReflect(s string, toType reflect.Type) (any, error) {
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
	return nil, fmt.Errorf("failed to convert %s to %s, unsupported type", s, toType.String())
}
